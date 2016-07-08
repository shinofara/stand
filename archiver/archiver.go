package archiver

import (
	"bytes"

	"github.com/shinofara/stand/archiver/compressor"
	"github.com/shinofara/stand/config"

	"github.com/uber-go/zap"
	"golang.org/x/net/context"
)

type Archiver struct {
	cfg        *config.Config
	ctx        context.Context
	compressor compressor.Compressor
}

func New(ctx context.Context, cfg *config.Config) *Archiver {
	c := compressor.New(ctx, cfg.CompressionConfig)

	return &Archiver{
		cfg:        cfg,
		ctx:        ctx,
		compressor: c,
	}
}

//Archive generates a buffer of compressed files.
func (a *Archiver) Archive() (*bytes.Buffer, error) {
	paths, err := find(a.cfg.Path)
	if err != nil {
		return nil, err
	}

	//ZIPファイル作成
	buf := new(bytes.Buffer)
	if err := a.compressor.Compress(buf, a.cfg.Path, paths); err != nil {
		return nil, err
	}

	logger := a.ctx.Value("logger").(zap.Logger)
	logger.Info(
		"Compression has been completed",
		zap.Int("size", len(buf.Bytes())),
	)

	return buf, nil
}
