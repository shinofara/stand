package cleaner

import (
	"fmt"
	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"
	"os"
	"sort"
)

func Exec(cfg *config.Config) error {

	files, _ := find.All(cfg.OutputDir)
	sort.Sort(sort.Reverse(sort.StringSlice(files)))

	var num int64 = 0
	for _, file := range files {
		if num >= cfg.LifeCyrcle {
			path := fmt.Sprintf("%s/%s", cfg.OutputDir, file)
			if err := os.RemoveAll(path); err != nil {
				return err
			}
		}

		num++
	}
	return nil
}
