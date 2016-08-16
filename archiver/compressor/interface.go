package compressor

import (
	"github.com/shinofara/stand/config"
	"context"
	"io"
)

type Compressor interface {
	Compress(io.Writer) error
}

func New(ctx context.Context, cfg *config.Config) Compressor {
	switch cfg.CompressionConfig.Format {
	case "tar":
		return NewTarCompressor(ctx, cfg)
	default:
		return NewZipCompressor(ctx, cfg)
	}
}
