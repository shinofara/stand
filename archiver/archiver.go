package archiver

import (
	"github.com/shinofara/stand/archiver/compressor"
	"github.com/shinofara/stand/config"
	"os"

	"github.com/uber-go/zap"
	"golang.org/x/net/context"
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

func (a *Archiver) Archive(output string) (string, error) {
	var compressedFile *os.File
	var err error
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
