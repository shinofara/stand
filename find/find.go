package find

import (
	"fmt"
	"os"
	"path/filepath"
)

//mode flgs
const (
	DeepSearchMode    = true  //deep search mode flg
	NotDeepSearchMode = false //not deep search mode flg
	FileOnlyMode      = true  //file only mode flg
	NotFileOnlyMode   = false //not file only mode flg
)

//findCallBack is a middleware function definition to the search results.
type findCallBack func(path string, file *os.File) error

//Find contains search options
type Find struct {
	callback     findCallBack
	deepMode     bool
	fileOnlyMode bool
	targetDir    string
}

//New creates a new Find
func New(callback findCallBack, targetDir string, deepMode bool, fileOnlyMode bool) *Find {
	return &Find{
		callback:     callback,
		deepMode:     deepMode,
		fileOnlyMode: fileOnlyMode,
		targetDir:    targetDir,
	}
}

//Run starts file search
func (f *Find) Run() error {
	return filepath.Walk(f.targetDir,
		func(path string, info os.FileInfo, err error) error {
			rel, err := filepath.Rel(f.targetDir, path)
			if err != nil {
				return err
			}

			if info == nil {
				return fmt.Errorf("file info is not found")
			}

			var filePath string

			if info.IsDir() {
				if rel == "." {
					return nil
				}

				if f.deepMode == NotDeepSearchMode {
					return filepath.SkipDir
				}

				if f.fileOnlyMode == FileOnlyMode {
					return nil
				}

				filePath = rel + string(os.PathSeparator)
			} else {
				filePath = rel
			}

			fullPath := filepath.Join(f.targetDir, filePath)
			addFile, _ := os.Open(fullPath)
			defer addFile.Close()

			if err := f.callback(filePath, addFile); err != nil {
				return err
			}

			return nil
		})
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

//findMiddleware is middeware of find.findCallBack interface's
func GetFiles(dir string) (FindFiles, error) {
	var files []File
	middleware := func(path string, file *os.File) error {
		info, err := file.Stat()
		if err != nil {
			return err
		}

		fullPath := filepath.Join(dir, path)

		fInfo := File{
			Info:     info,
			Path:     path,
			FullPath: fullPath}

		files = append(files, fInfo)
		return nil
	}

	f := New(middleware, dir, NotDeepSearchMode, FileOnlyMode)
	if err := f.Run(); err != nil {
		return nil, err
	}

	return files, nil
}
