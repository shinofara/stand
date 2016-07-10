package location

import (
	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"
	"os"
	"sort"

	"bytes"
)

type Local struct {
	storageCfg *config.StorageConfig
}

func NewLocal(storageCfg *config.StorageConfig) *Local {
	return &Local{
		storageCfg: storageCfg,
	}
}

func (l *Local) Save(filename string, buf *bytes.Buffer) error {
	if err := mkdir(l.storageCfg.Path); err != nil {
		return err
	}

	storagePath := l.storageCfg.Path + "/" + filename
	f, err := os.Create(storagePath)
	if err != nil {
		return err
	}
	f.Write(buf.Bytes())
	f.Close()

	return nil
}

func (l *Local) Clean() error {
	files, _ := find.Find(l.storageCfg.Path, false, true)
	sort.Sort(files)

	var num int64 = 0
	for _, file := range files {
		if num > l.storageCfg.LifeCyrcle {
			if err := os.RemoveAll(file.FullPath); err != nil {
				return err
			}
		}

		num++
	}
	return nil
}

func mkdir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if err := os.Mkdir(path, 0777); err != nil {
			return err
		}
	}
	return nil
}
