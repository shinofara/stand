package backup

import (
	"context"
	"os"

	"github.com/shinofara/stand/backup/location"
	"github.com/shinofara/stand/config"
)

//Backup manages all of the settings for backup
type Backup struct {
	Config *config.Config
	ctx    context.Context
}

func New(ctx context.Context, cfg *config.Config) *Backup {
	return &Backup{
		Config: cfg,
		ctx:    ctx,
	}
}

func (b *Backup) Exec(filepath string) error {
	var loc location.Location

	for _, storageCfg := range b.Config.StorageConfigs {
		loc = location.New(&storageCfg)
		if err := loc.Save(filepath); err != nil {
			return err
		}

		if storageCfg.LifeCyrcle == 0 {
			break
		}

		if err := loc.Clean(); err != nil {
			return err
		}
	}

	if err := os.Remove(filepath); err != nil {
		return err
	}

	return nil
}
