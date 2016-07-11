package location

import (
	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"
	"os"
	"sort"
	"fmt"
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

type FindFiles []File
type File struct {
	Info     os.FileInfo
	Path     string
	FullPath string
}

func (fi FindFiles) Len() int {
	return len(fi)
}
func (fi FindFiles) Swap(i, j int) {
	fi[i], fi[j] = fi[j], fi[i]
}
func (fi FindFiles) Less(i, j int) bool {
	return fi[j].Info.ModTime().Unix() < fi[i].Info.ModTime().Unix()
}

//Clean contains the information about the cleaning.
type Clean struct {
	targets FindFiles
	storageCfg *config.StorageConfig 
}

//New creates a new Clean
func NewCLean(storageCfg *config.StorageConfig) *Clean {
	return &Clean{
		storageCfg: storageCfg, 
	}
}

//Run removes the old generation of the file.
func (c *Clean) Run() error {
	f := find.New(c.findMiddleware, c.storageCfg.Path, find.NotDeepSearchMode, find.FileOnlyMode)
	if err := f.Run(); err != nil {
		return err
	}

	for _, file := range c.targets {
		fmt.Println(file)
	}
	return nil
}

//findMiddleware is middeware of find.findCallBack interface's
func (c *Clean) findMiddleware(path string, file *os.File) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}

	fullPath := fmt.Sprintf("%s/%s", c.storageCfg.Path, path)
	
	fInfo := File{
		Info:     info,
		Path:     path,
		FullPath: fullPath}
	
	c.targets = append(c.targets, fInfo)
	return nil
}

func (l *Local) Clean() error {
	c := NewCLean(l.storageCfg)
	if err := c.Run(); err != nil {
		panic(err)
	}
	files := c.targets
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
