package coordinator

import (
	"context"
	"sort"

	"github.com/uber-go/zap"

	"github.com/shinofara/stand/archiver"
	"github.com/shinofara/stand/backup"
	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"
)

type Coordinator struct {
	ctx context.Context
	cfg *config.Config
}

func New(ctx context.Context, cfg *config.Config) *Coordinator {
	return &Coordinator{
		ctx: ctx,
		cfg: cfg,
	}
}

func (c *Coordinator) Run() error {
	logger := c.ctx.Value("logger").(zap.Logger)
	var filepath string
	var err error

	if c.cfg.Type == config.TYPE_DIR {
		a := archiver.New(c.ctx, c.cfg)
		filepath, err = a.Archive()
	} else {
		var files find.FindFiles
		files, err = find.GetFiles(c.cfg.Path)

		if len(files) < 1 {
			logger.Info(
				"file not found",
				zap.String("path", c.cfg.Path),
			)
			return nil
		}

		sort.Sort(files)
		filepath = files[0].FullPath
	}

	if err != nil {
		return err
	}

	b := backup.New(c.ctx, c.cfg)
	if err := b.Exec(filepath); err != nil {
		return err
	}

	return nil

}
