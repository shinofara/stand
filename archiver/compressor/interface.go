package compressor

import (
	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"
	"golang.org/x/net/context"
	"io"
)

type Compressor interface {
	Compress(io.Writer, []find.File) error
}

func New(ctx context.Context, cfg *config.Config) Compressor {
	switch cfg.CompressionConfig.Format {
	case "tar":
		return NewTarCompressor(ctx, cfg)
	default:
		return NewZipCompressor(ctx, cfg)
	}
}
