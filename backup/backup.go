package backup

import (
	"github.com/shinofara/stand/backup/location"
	"github.com/shinofara/stand/config"
	"path"
)

type Backup struct {
	Config *config.Config
}

func (b *Backup) Exec(file string) error {
	var loc location.Location
	dir, filename := path.Split(file)

	switch b.Config.StorageConfig.Type {
	case "s3":
		loc = location.NewS3(b.Config)
	default:
		loc = &location.Local{Config: b.Config}
	}

	if err := loc.Save(dir, filename); err != nil {
		return err
	}

	if err := loc.Clean(); err != nil {
		return err
	}

	return nil
}
