package zip

import (
	"archive/zip"
	"fmt"
	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"
	"io/ioutil"
	"os"
	"time"
)

const (
	TIME_FORMAT = "20060102150405"
)

func Compress(cfg *config.Config) error {
	var zipfile *os.File
	var err error

	if _, err := os.Stat(cfg.OutputDir); err != nil {
		if err := os.Mkdir(cfg.OutputDir, 0777); err != nil {
			return err
		}
	}

	// Create a buffer to write our archive to.
	timestamp := time.Now().Format(TIME_FORMAT)
	output := fmt.Sprintf("%s/%s_%s.zip", cfg.OutputDir, cfg.ZipName, timestamp)
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
