package compressor

import (
	"context"
	"github.com/shinofara/stand/config"
	"io"
)

//Compressor defines interface.
type Compressor interface {
	Compress(io.Writer) error
}

//New creates a structure that is based on the interface.
func New(ctx context.Context, cfg *config.Config) Compressor {
	switch cfg.CompressionConfig.Format {
	case "tar":
		return NewTarCompressor(ctx, cfg)
	default:
		return NewZipCompressor(ctx, cfg)
	}
}
