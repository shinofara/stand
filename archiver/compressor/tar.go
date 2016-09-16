package compressor

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"io"
	"os"

	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"
)

type TarCompressor struct {
	ctx context.Context
	cfg *config.Config
}

func NewTarCompressor(ctx context.Context, cfg *config.Config) *TarCompressor {
	return &TarCompressor{
		ctx: ctx,
		cfg: cfg,
	}
}

func (c *TarCompressor) Compress(compressedFile io.Writer) error {

	gw := gzip.NewWriter(compressedFile)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	middeware := func(path string, file *os.File) error {
		info, err := file.Stat()
		if err != nil {
			return err
		}

		if info.IsDir() {
			//dirは不要
			return nil
		}

		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		hdr.Name = path
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		if _, err = io.Copy(tw, file); err != nil {
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
