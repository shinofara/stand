package coordinator

import (
	"golang.org/x/net/context"

	"github.com/shinofara/stand/archiver"
	"github.com/shinofara/stand/backup"
	"github.com/shinofara/stand/config"
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

func (c *Coordinator) Perform() error {
	a := archiver.New(c.ctx, c.cfg)
	filepath, err := a.Archive()

	if err != nil {
		return err
	}

	b := backup.New(c.ctx, c.cfg)
	if err := b.Exec(filepath); err != nil {
		return err
	}

	return nil

}
