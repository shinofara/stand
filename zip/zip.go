package zip

import (
	"archive/zip"
	"fmt"
	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"
	"io/ioutil"
	"os"
)

func Compress(cfg *config.Config) error {
	var zipfile *os.File
	var err error

	if err := os.Mkdir(cfg.OutputDir, 0777); err != nil {
		return err
	}

	// Create a buffer to write our archive to.
	output := fmt.Sprintf("%s/%s", cfg.OutputDir, cfg.ZipName)
	if zipfile, err = os.Create(output); err != nil {
		return err
	}
	defer zipfile.Close()

	w := zip.NewWriter(zipfile)

	files, _ := find.All(cfg.TargetDir)
	for _, file := range files {
		f, err := w.Create(file)
		if err != nil {
			return err
		}

		contents, _ := ioutil.ReadFile(cfg.TargetDir + "/" + file)
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
