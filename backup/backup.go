package backup

import (
	"github.com/shinofara/stand/backup/location"
	"github.com/shinofara/stand/config"
	"os"
	"path"
)

type Backup struct {
	Config *config.Config
}

func (b *Backup) Exec(file string) error {
	var loc location.Location
	dir, filename := path.Split(file)

	for _, storageCfg := range b.Config.StorageConfigs {
		switch storageCfg.Type {
		case "s3":
			loc = location.NewS3(&storageCfg)
		default:
			loc = location.NewLocal(&storageCfg)
		}

		if err := loc.Save(dir, filename); err != nil {
			return err
		}

		if storageCfg.LifeCyrcle == 0 {
			break
		}

		if err := loc.Clean(); err != nil {
			return err
		}
	}

	if err := os.RemoveAll(file); err != nil {
		return err
	}

	return nil
}
