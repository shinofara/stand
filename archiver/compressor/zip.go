package compressor

import (
	"archive/zip"
	"context"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"
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

func (c *ZipCompressor) Compress(compressedFile io.Writer) error {
	w := zip.NewWriter(compressedFile)
	defer w.Close()

	middeware := func(path string, file *os.File) error {
		hdr, err := createZipFileHeader(file, path)
		if err != nil {
			return err
		}

		wf, err := w.CreateHeader(hdr)
		if err != nil {
			return err
		}

		contents, err := ioutil.ReadAll(file)
		_, err = wf.Write(contents)
		if err != nil {
			return err
		}

		return nil
	}

	f := find.New(middeware, c.cfg.Path, find.DeepSearchMode, find.NotFileOnlyMode)
	if err := f.Run(); err != nil {
		return err
	}

	return nil
}

func createZipFileHeader(f *os.File, path string) (*zip.FileHeader, error) {
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	hdr, err := zip.FileInfoHeader(info)
	if err != nil {
		return nil, err
	}

	hdr.Name = path

	local := time.Now().Local()

	//現時刻のオフセットを取得
	_, offset := local.Zone()

	//差分を追加
	hdr.SetModTime(hdr.ModTime().Add(time.Duration(offset) * time.Second))

	return hdr, nil
}
