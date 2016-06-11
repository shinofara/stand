package zip

import (
	"archive/zip"
	"github.com/shinofara/stand/find"
	"io/ioutil"
	"os"
)

func Compress(targetDir string, zipName string) error {
	var zipfile *os.File
	var err error

	// Create a buffer to write our archive to.
	if zipfile, err = os.Create(zipName); err != nil {
		return err
	}
	defer zipfile.Close()

	w := zip.NewWriter(zipfile)

	files, _ := find.All(targetDir)
	for _, file := range files {
		f, err := w.Create(file)
		if err != nil {
			return err
		}

		contents, _ := ioutil.ReadFile(targetDir + "/" + file)
		_, err = f.Write(contents)
		if err != nil {
			return err
		}
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}
