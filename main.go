package main

import (
	"flag"
	"github.com/shinofara/stand/backup"
	"github.com/shinofara/stand/compressor"
	"github.com/shinofara/stand/config"
	"log"
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

	cfgs, _ := config.New(args.ConfigPath)

	for _, cfg := range *cfgs {
		output, err := compressor.Compress(cfg)

		if err != nil {
			panic(err)
		}

		b := &backup.Backup{Config: cfg}

		log.Print(output)
		b.Exec(output)
	}
}
