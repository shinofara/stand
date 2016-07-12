package location

import (
	"github.com/shinofara/stand/config"
)

type Location interface {
	Save(filename string) error
	Clean() error
}

func New(cfg *config.StorageConfig) Location {
	switch cfg.Type {
	case "s3":
		return NewS3(cfg)
	default:
		return NewLocal(cfg)
	}
}
