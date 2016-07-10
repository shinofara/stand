package compressor

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"

	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"

	"golang.org/x/net/context"
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

func (c *TarCompressor) Compress(compressedFile io.Writer, files []find.File) error {

	gw := gzip.NewWriter(compressedFile)
	defer gw.Close()

	tw := tar.NewWriter(gw)

	for _, f := range files {
		err := func(f find.File, tw *tar.Writer) error {
			if f.Info.IsDir() {
				//dirは不要
				return nil
			}

			file, err := os.Open(f.FullPath)
			if err != nil {
				return err
			}
			defer file.Close()

			hdr, err := tar.FileInfoHeader(f.Info, "")
			if err != nil {
				return err
			}

			hdr.Name = f.Path
			if err := tw.WriteHeader(hdr); err != nil {
				return err
			}

			if _, err = io.Copy(tw, file); err != nil {
				return err
			}

			return nil
		}(f, tw)

		if err != nil {
			return err
		}
	}

	if err := tw.Close(); err != nil {
		return err
	}

	return nil
}
