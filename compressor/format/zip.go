package format

import (
	"archive/zip"
	"io"
	"io/ioutil"
)

type ZipCompressor struct{}

func NewZipCompressor() *ZipCompressor {
	return &ZipCompressor{}
}

func (c *ZipCompressor) Compress(compressedFile io.Writer, targetDir string, files []string) error {
	w := zip.NewWriter(compressedFile)

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

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}
