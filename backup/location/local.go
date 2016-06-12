package location

import (
	"fmt"
	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"
	"os"
	"sort"
)

type Local struct {
	Config *config.Config
}

func (l *Local) Save(localDir string, filename string) error {
	if err := mkdir(l.Config.StorageConfig.Path); err != nil {
		return err
	}

	if err := os.Rename(localDir+"/"+filename, l.Config.StorageConfig.Path+"/"+filename); err != nil {
		return err
	}

	return nil
}

func (l *Local) Clean() error {
	files, _ := find.All(l.Config.StorageConfig.Path)
	sort.Sort(sort.Reverse(sort.StringSlice(files)))

	var num int64 = 0
	for _, file := range files {
		if num >= l.Config.StorageConfig.LifeCyrcle {
			path := fmt.Sprintf("%s/%s", l.Config.StorageConfig.Path, file)
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
