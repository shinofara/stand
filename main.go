package main

import (
	"flag"
	"github.com/shinofara/stand/cleaner"
	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/zip"
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

	cfg, _ := config.New(args.ConfigPath)

	err := zip.Compress(cfg)

	if err != nil {
		panic(err)
	}

	if err := cleaner.Exec(cfg); err != nil {
		panic(err)
	}
}
