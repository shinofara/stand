package main

import (
	"flag"
	"github.com/shinofara/stand/backup"
	"github.com/shinofara/stand/compressor"
	"github.com/shinofara/stand/config"
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
			uploadFileName, err = compressor.Compress(cfg)
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
