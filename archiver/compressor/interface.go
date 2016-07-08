package compressor

import (
	"github.com/shinofara/stand/config"
	"golang.org/x/net/context"
	"io"
)

type Compressor interface {
	Compress(io.Writer, string, []string) error
}

func New(ctx context.Context, cfg *config.CompressionConfig) Compressor {
	switch cfg.Format {
	case "tar":
		return NewTarCompressor(ctx)
	default:
		return NewZipCompressor(ctx)
	}
}
