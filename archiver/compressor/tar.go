package compressor

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"fmt"

	"github.com/shinofara/stand/config"

	"golang.org/x/net/context"
)

type (
	TarCompressor struct {
	ctx context.Context
	cfg *config.Config
	}

	tarCompression struct {
		tw *tar.Writer
	}	
)

func (c *tarCompression) perform(d string, f string) {

	fw, err := os.Open(fmt.Sprintf("%s/%s", d, f))
	if err != nil {
		panic(err)
	}

	info, err := fw.Stat()
	if err != nil {
		fw.Close()
		panic(err)
	}
	defer fw.Close()	

	if info.IsDir() {
		return 
	}

	hdr, err := tar.FileInfoHeader(info, "")
	if err != nil {
		panic(err)
	}

	hdr.Name = f
	if err := c.tw.WriteHeader(hdr); err != nil {
		panic(err)
	}

	if _, err = io.Copy(c.tw, fw); err != nil {
		panic(err)
	}
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

	tc := &tarCompression{
		tw: tw,
	}
	
	for _, f := range files {
		tc.perform(c.cfg.Path, f)		
	}
	tw.Close()
	
	return nil

}
