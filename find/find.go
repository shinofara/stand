package find

import (
	"fmt"
	"os"
	"path/filepath"
)

type File struct {
	Info     os.FileInfo
	Path     string
	FullPath string
}

func Find(targetDir string) ([]File, error) {
	var paths []File
	err := filepath.Walk(targetDir,
		func(path string, info os.FileInfo, err error) error {
			rel, err := filepath.Rel(targetDir, path)
			if err != nil {
				return err
			}

			if info == nil {
				return fmt.Errorf("file info is not found")
			}

			var filePath string

			if info.IsDir() {
				filePath = fmt.Sprintf("%s/", rel)
			} else {
				filePath = rel
			}

			fullPath := fmt.Sprintf("%s/%s", targetDir, filePath)

			info, err = os.Stat(fullPath)
			if err != nil {
				return err
			}

			fInfo := File{
				Info:     info,
				Path:     filePath,
				FullPath: fullPath}

			paths = append(paths, fInfo)
			return nil
		})

	if err != nil {
		return nil, err
	}

	return paths, nil
}
