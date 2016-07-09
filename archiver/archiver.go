package archiver

import (
	"bytes"

	"github.com/shinofara/stand/archiver/compressor"
	"github.com/shinofara/stand/config"

	"github.com/uber-go/zap"
	"golang.org/x/net/context"
)

type Archiver struct {
	cfg *config.Config
	ctx context.Context
}

func New(ctx context.Context, cfg *config.Config) *Archiver {

	return &Archiver{
		cfg: cfg,
		ctx: ctx,
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
	c := compressor.New(a.ctx, a.cfg)
	if err := c.Compress(buf, paths); err != nil {
		return nil, err
	}

	logger := a.ctx.Value("logger").(zap.Logger)
	logger.Info(
		"Compression has been completed",
		zap.Int("size", len(buf.Bytes())),
	)

	return buf, nil
}
