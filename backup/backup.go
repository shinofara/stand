package backup

import (
	"bytes"
	"fmt"
	"time"

	"github.com/shinofara/stand/backup/location"
	"github.com/shinofara/stand/config"

	"golang.org/x/net/context"
)

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

func (b *Backup) Exec(buf *bytes.Buffer) error {
	var loc location.Location
	filename := b.makeCompressedFileName()

	for _, storageCfg := range b.Config.StorageConfigs {
		loc = location.New(&storageCfg)
		if err := loc.Save(filename, buf); err != nil {
			return err
		}

		if storageCfg.LifeCyrcle == 0 {
			break
		}

		if err := loc.Clean(); err != nil {
			return err
		}
	}

	return nil
}

const (
	TimeFormat = "20060102150405"
)

func (b *Backup) makeCompressedFileName() string {
	timestamp := time.Now().Format(TimeFormat)

	extention := "zip"
	switch b.Config.CompressionConfig.Format {
	case "tar":
		extention = "tar.gz"
	}

	var filename string
	if b.Config.CompressionConfig.Prefix != "" {
		filename = fmt.Sprintf("%s%s.%s", b.Config.CompressionConfig.Prefix, timestamp, extention)
	} else {
		filename = fmt.Sprintf("%s.%s", timestamp, extention)
	}
	return filename
}
