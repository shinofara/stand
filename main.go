package main

import (
	"flag"
	"fmt"
	"github.com/shinofara/stand/archiver"
	"github.com/shinofara/stand/backup"
	"github.com/shinofara/stand/config"
	"time"
)

const (
	TIME_FORMAT = "20060102150405"
)

type Args struct {
	ConfigPath string
}

func main() {
	var args Args

	flag.StringVar(&args.ConfigPath, "conf", "", "-c path to config yaml")
	// コマンドライン引数を解析
	flag.Parse()

	if args.ConfigPath == "" {
		panic("-c is empty")
	}

	cfgs, _ := config.Load(args.ConfigPath)

	for _, cfg := range *cfgs {
		var uploadFileName string
		var err error

		switch cfg.Type {
		case "dir":
			output := makeCompressedFileName(cfg)
			a := archiver.New(cfg)
			uploadFileName, err = a.Archive(output)
			if err != nil {
				panic(err)
			}
		case "file":
			uploadFileName = cfg.Path
		default:
			panic("upload target type is not found")
		}

		b := &backup.Backup{Config: cfg}
		if err := b.Exec(uploadFileName); err != nil {
			panic(err)
		}
	}
}

//input, output
func makeCompressedFileName(cfg *config.Config) string {
	timestamp := time.Now().Format(TIME_FORMAT)

	extention := "zip"
	switch cfg.CompressionConfig.Format {
	case "tar":
		extention = "tar.gz"
	}

	var output string
	if cfg.CompressionConfig.Prefix != "" {
		output = fmt.Sprintf("%s%s.%s", cfg.CompressionConfig.Prefix, timestamp, extention)
	} else {
		output = fmt.Sprintf("%s.%s", timestamp, extention)
	}
	return "/tmp/" + output
}
