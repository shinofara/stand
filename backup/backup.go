package backup

import (
	"github.com/shinofara/stand/backup/location"
	"github.com/shinofara/stand/config"
)

type Backup struct {
	Config *config.Config
}

func (b *Backup) Exec(file string) error {
	var loc location.Location
	loc = &location.Local{b.Config}
	if err := loc.Save(file); err != nil {
		return err
	}

	if err := loc.Clean(); err != nil {
		return err
	}

	return nil
}
