package entity

import (
	"fmt"
)

type LocalFile struct {
	Dir  string
	Name string
}

func (f *LocalFile) CreateFullPath() string {
	return fmt.Sprintf("%s/%s", f.Dir, f.Name)
}
