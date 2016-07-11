package find

import (
	"os"
	"fmt"
	"path/filepath"
)

//mode flgs
const (
	DeepSearchMode = true //deep search mode flg
	NotDeepSearchMode = false //not deep search mode flg
	FileOnlyMode = true //file only mode flg
	NotFileOnlyMode = false //not file only mode flg
)

//findCallBack is a middleware function definition to the search results.
type findCallBack func(path string, file *os.File) error

//Find contains search options
type Find struct {
	callback findCallBack
	deepMode bool
	fileOnlyMode bool
	targetDir string
}

//New creates a new Find
func New(callback findCallBack, targetDir string, deepMode bool, fileOnlyMode bool) *Find {
	return &Find{
		callback: callback,
		deepMode: deepMode,
		fileOnlyMode: fileOnlyMode,
		targetDir: targetDir,
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
				
				filePath = fmt.Sprintf("%s/", rel)
				} else {
					filePath = rel
				}
			
			fullPath := fmt.Sprintf("%s/%s", f.targetDir, filePath)
			addFile, _ := os.Open(fullPath)
			defer addFile.Close()
			
			if err := f.callback(filePath, addFile); err != nil {
				return err
			}

			return nil
		})
}
