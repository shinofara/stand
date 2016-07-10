package compressor

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"time"

	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"

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

func (c *ZipCompressor) Compress(compressedFile io.Writer, files []find.File) error {
	w := zip.NewWriter(compressedFile)

	for _, f := range files {
		if f.Info.IsDir() {
			continue
		}

		hdr, err := createZipFileHeader(f)
		if err != nil {
			return err
		}

		wf, err := w.CreateHeader(hdr)
		if err != nil {
			return err
		}

		contents, _ := ioutil.ReadFile(f.FullPath)
		_, err = wf.Write(contents)
		if err != nil {
			return err
		}
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

func createZipFileHeader(f find.File) (*zip.FileHeader, error) {
	hdr, err := zip.FileInfoHeader(f.Info)
	if err != nil {
		return nil, err
	}

	hdr.Name = f.Path

	local := time.Now().Local()

	//現時刻のオフセットを取得
	_, offset := local.Zone()

	//差分を追加
	hdr.SetModTime(hdr.ModTime().Add(time.Duration(offset) * time.Second))

	return hdr, nil
}
