package compressor

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/shinofara/stand/config"

	"golang.org/x/net/context"
)

type ZipCompressor struct {
	ctx context.Context
	cfg *config.Config
}

func NewZipCompressor(ctx context.Context, cfg *config.Config) *ZipCompressor {
	return &ZipCompressor{
		ctx: ctx,
		cfg: cfg,
	}
}

func (c *ZipCompressor) Compress(compressedFile io.Writer, files []string) error {
	w := zip.NewWriter(compressedFile)

	for _, filename := range files {
		filepath := fmt.Sprintf("%s/%s", c.cfg.Path, filename)
		info, err := os.Stat(filepath)
		if err != nil {
			return err
		}

		if info.IsDir() {
			continue
		}

		hdr, err := createZipFileHeader(filename, info)

		f, err := w.CreateHeader(hdr)
		if err != nil {
			return err
		}

		contents, _ := ioutil.ReadFile(filepath)
		_, err = f.Write(contents)
		if err != nil {
			return err
		}
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

func createZipFileHeader(filename string, info os.FileInfo) (*zip.FileHeader, error) {
	hdr, err := zip.FileInfoHeader(info)
	if err != nil {
		return nil, err
	}

	hdr.Name = filename

	local := time.Now().Local()

	//現時刻のオフセットを取得
	_, offset := local.Zone()

	//差分を追加
	hdr.SetModTime(hdr.ModTime().Add(time.Duration(offset) * time.Second))

	return hdr, nil
}
