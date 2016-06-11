package format

import (
	"archive/tar"
	"io"
	"io/ioutil"
)

type TarCompressor struct{}

func NewTarCompressor() *TarCompressor {
	return &TarCompressor{}
}

func (c *TarCompressor) Compress(compressedFile io.Writer, targetDir string, files []string) error {

	tw := tar.NewWriter(compressedFile)

	for _, file := range files {
		body, _ := ioutil.ReadFile(targetDir + "/" + file)
		hdr := &tar.Header{
			Name: file,
			Mode: 0600,
			Size: int64(len(body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		if _, err := tw.Write(body); err != nil {
			return err
		}
	}

	if err := tw.Close(); err != nil {
		return err
	}

	return nil
}
