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

func Compress(cfg *config.Config) (string, error) {
	var compressedFile *os.File
	var err error

	var compressor format.Compressor

	switch cfg.CompressionConfig.Format {
	case "zip":
		compressor = format.NewZipCompressor()
	case "tar":
		compressor = format.NewTarCompressor()
	default:
		return "", fmt.Errorf("Not exists compression format")
	}

	output := makeCompressedFileName(cfg)

	paths, _ := find(cfg.Path)

	//ZIPファイル作成
	if compressedFile, err = os.Create("/tmp/" + output); err != nil {
		return "", err
	}
	defer compressedFile.Close()

	if err := compressor.Compress(compressedFile, cfg.Path, paths); err != nil {
		return "", err
	}

	return output, nil
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
	output := fmt.Sprintf("%s_%s.%s", cfg.CompressionConfig.Prefix, timestamp, extention)
	return output
}
