package archiver

import (
	"time"
	"os"
	"fmt"

	"github.com/shinofara/stand/archiver/compressor"
	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/find"	

	"github.com/uber-go/zap"
	"golang.org/x/net/context"
)

const (
	TimeFormat = "20060102150405"
)

type Archiver struct {
	cfg *config.Config
	ctx context.Context
}

func New(ctx context.Context, cfg *config.Config) *Archiver {

	return &Archiver{
		cfg: cfg,
		ctx: ctx,
	}
}

func findFiles(findPath string) ([]string, error) {
	var files []string

	middeware := func(path string, file *os.File) error {
		files = append(files, path)
		return nil
	}
	
	f := find.New(middeware, findPath, find.DeepSearchMode, find.NotFileOnlyMode)
	if err := f.Run(); err != nil {
		return nil, err
	}

	return files, nil
}

//Archive generates a buffer of compressed files.
func (a *Archiver) Archive() (string, error) {
	logger := a.ctx.Value("logger").(zap.Logger)

	filepath :=a.makeCompressedFileName()
	buf, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer buf.Close()
	
	c := compressor.New(a.ctx, a.cfg)

	files, err := findFiles(a.cfg.Path)
	if err != nil {
		return "", err
	}
	
	if err := c.Compress(buf, files); err != nil {
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
	return fmt.Sprintf("/tmp/%s", filename)
}

