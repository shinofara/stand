package archiver

import (
	"github.com/shinofara/stand/archiver/compressor"
	"github.com/shinofara/stand/config"
	"os"
)

type Archiver struct {
	cfg        *config.Config
	compressor compressor.Compressor
}

func New(cfg *config.Config) *Archiver {

	var c compressor.Compressor

	switch cfg.CompressionConfig.Format {
	case "tar":
		c = compressor.NewTarCompressor()
	default:
		c = compressor.NewZipCompressor()
	}

	return &Archiver{
		cfg:        cfg,
		compressor: c,
	}
}

func (a *Archiver) Archive(output string) (string, error) {
	var compressedFile *os.File
	var err error
	paths, _ := find(a.cfg.Path)

	//ZIPファイル作成
	if compressedFile, err = os.Create(output); err != nil {
		return "", err
	}
	defer compressedFile.Close()

	if err := a.compressor.Compress(compressedFile, a.cfg.Path, paths); err != nil {
		return "", err
	}

	return output, nil
}
