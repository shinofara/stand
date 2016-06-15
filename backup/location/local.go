package location

import (
	"fmt"
	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"
	"os"
	"sort"
)

type Local struct {
	storageCfg *config.StorageConfig
}

func NewLocal(storageCfg *config.StorageConfig) *Local {
	return &Local{
		storageCfg: storageCfg,
	}
}

func (l *Local) Save(localDir string, filename string) error {
	if err := mkdir(l.storageCfg.Path); err != nil {
		return err
	}

	tmpPath := localDir + "/" + filename
	storagePath := l.storageCfg.Path + "/" + filename
	if err := os.Rename(tmpPath, storagePath); err != nil {
		return err
	}

	return nil
}

func (l *Local) Clean() error {
	files, _ := find.All(l.storageCfg.Path)
	sort.Sort(sort.Reverse(sort.StringSlice(files)))

	var num int64 = 0
	for _, file := range files {
		if num > l.storageCfg.LifeCyrcle {
			path := fmt.Sprintf("%s/%s", l.storageCfg.Path, file)
			if err := os.RemoveAll(path); err != nil {
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
