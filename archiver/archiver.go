package archiver

import (
	"fmt"
	"github.com/shinofara/stand/archiver/compressor"
	"github.com/shinofara/stand/config"
	"os"
	"time"

	"github.com/uber-go/zap"
	"golang.org/x/net/context"
)

const (
	TimeFormat = "20060102150405"
)

type Archiver struct {
	cfg        *config.Config
	ctx        context.Context
	compressor compressor.Compressor
}

func New(ctx context.Context, cfg *config.Config) *Archiver {
	var c compressor.Compressor

	switch cfg.CompressionConfig.Format {
	case "tar":
		c = compressor.NewTarCompressor(ctx)
	default:
		c = compressor.NewZipCompressor(ctx)
	}

	return &Archiver{
		cfg:        cfg,
		ctx:        ctx,
		compressor: c,
	}
}

func (a *Archiver) Archive() (string, error) {
	output := a.makeCompressedFileName()
	var compressedFile *os.File

	paths, err := find(a.cfg.Path)
	if err != nil {
		return "", err
	}

	//ZIPファイル作成
	if compressedFile, err = os.Create(output); err != nil {
		return "", err
	}
	defer compressedFile.Close()

	if err := a.compressor.Compress(compressedFile, a.cfg.Path, paths); err != nil {
		return "", err
	}

	info, err := compressedFile.Stat()
	if err != nil {
		return "", err
	}

	logger := a.ctx.Value("logger").(zap.Logger)
	logger.Info(
		"Compression has been completed",
		zap.String("name", output),
		zap.Int64("size", info.Size()),
	)

	return output, nil
}

func (a *Archiver) makeCompressedFileName() string {
	timestamp := time.Now().Format(TimeFormat)

	extention := "zip"
	switch a.cfg.CompressionConfig.Format {
	case "tar":
		extention = "tar.gz"
	}

	var output string
	if a.cfg.CompressionConfig.Prefix != "" {
		output = fmt.Sprintf("%s%s.%s", a.cfg.CompressionConfig.Prefix, timestamp, extention)
	} else {
		output = fmt.Sprintf("%s.%s", timestamp, extention)
	}
	return "/tmp/" + output
}
