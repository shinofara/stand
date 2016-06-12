package backup

//バックアップ

import (
	"os"
)

type Backup struct {
	BackupDir string
}

func (b *Backup) Exec(file string) error {
	if err := mkdir(b.BackupDir); err != nil {
		return err
	}

	if err := os.Rename("/tmp/"+file, b.BackupDir+"/"+file); err != nil {
		return err
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
