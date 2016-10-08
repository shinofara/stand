package archiver

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/shinofara/stand/archiver/compressor"
	"github.com/shinofara/stand/config"

	"context"
	"github.com/uber-go/zap"
)

//file const.
const (
	TimeFormat = "20060102150405" //TimeFormat is prefix of archive file.
)

//Archiver contains archive configs.
type Archiver struct {
	cfg *config.Config
	ctx context.Context
}

//New creates a Archiver.
func New(ctx context.Context, cfg *config.Config) *Archiver {

	return &Archiver{
		cfg: cfg,
		ctx: ctx,
	}
}

//Archive generates a buffer of compressed files.
func (a *Archiver) Archive() (string, error) {
	logger := a.ctx.Value("logger").(zap.Logger)

	filepath := a.makeCompressedFileName()
	buf, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer buf.Close()

	c := compressor.New(a.ctx, a.cfg)

	if err := c.Compress(buf); err != nil {
		return "", err
	}

	info, err := buf.Stat()
	if err != nil {
		return "", err
	}

	logger.Info(
		"Compression has been completed",
		zap.Int64("size", info.Size()),
	)

	return filepath, nil
}

//makeCompressedFileName makes a compress file name.
func (a *Archiver) makeCompressedFileName() string {
	timestamp := time.Now().Format(TimeFormat)

	extention := "zip"
	switch a.cfg.CompressionConfig.Format {
	case "tar":
		extention = "tar.gz"
	}

	var filename string
	if a.cfg.CompressionConfig.Prefix != "" {
		filename = fmt.Sprintf("%s%s.%s", a.cfg.CompressionConfig.Prefix, timestamp, extention)
	} else {
		filename = fmt.Sprintf("%s.%s", timestamp, extention)
	}
	return filepath.Join(os.TempDir(), filename)
}
