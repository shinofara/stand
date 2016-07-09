package compressor

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/shinofara/stand/config"

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

func (c *TarCompressor) Compress(compressedFile io.Writer, files []string) error {

	gw := gzip.NewWriter(compressedFile)
	defer gw.Close()

	tw := tar.NewWriter(gw)

	for _, filename := range files {
		filepath := fmt.Sprintf("%s/%s", c.cfg.Path, filename)

		info, err := os.Stat(filepath)
		if err != nil {
			return err
		}

		if info.IsDir() {
			continue
		}

		file, err := os.Open(filepath)
		if err != nil {
			return err
		}
		defer file.Close()

		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		hdr.Name = filename
		if err := tw.WriteHeader(hdr); err != nil {
			log.Print(err)
			return err
		}

		if _, err = io.Copy(tw, file); err != nil {
			return err
		}
	}

	if err := tw.Close(); err != nil {
		return err
	}

	return nil
}
