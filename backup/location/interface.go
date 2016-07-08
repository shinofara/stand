package location

import (
	"bytes"
	"github.com/shinofara/stand/config"
)

type Location interface {
	Save(filename string, buf *bytes.Buffer) error
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
