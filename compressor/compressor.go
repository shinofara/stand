package compressor

import (
	"fmt"
	"github.com/shinofara/stand/compressor/format"
	"github.com/shinofara/stand/config"
	"os"
	"time"
)

const (
	TIME_FORMAT = "20060102150405"
)

func Compress(cfg *config.Config) error {
	var compressedFile *os.File
	var err error

	if err := mkdir(cfg.OutputDir); err != nil {
		return err
	}

	var compressor format.Compressor

	switch cfg.CompressionConfig.Format {
	case "zip":
		compressor = format.NewZipCompressor()
	case "tar":
		compressor = format.NewTarCompressor()
	default:
		return fmt.Errorf("Not exists compression format")
	}

	output := makeCompressedFileName(cfg)

	paths, _ := find(cfg.TargetDir)

	//ZIPファイル作成
	if compressedFile, err = os.Create(output); err != nil {
		return err
	}
	defer compressedFile.Close()

	if err := compressor.Compress(compressedFile, cfg.TargetDir, paths); err != nil {
		return err
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

func makeCompressedFileName(cfg *config.Config) string {
	timestamp := time.Now().Format(TIME_FORMAT)

	extention := ""
	switch cfg.CompressionConfig.Format {
	case "zip":
		extention = "zip"
	case "tar":
		extention = "tar.gz"
	default:
		panic("")
	}
	output := fmt.Sprintf("%s/%s_%s.%s", cfg.OutputDir, cfg.CompressionConfig.Prefix, timestamp, extention)
	return output
}
