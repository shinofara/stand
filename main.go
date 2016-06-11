package main

import (
	"flag"
	"fmt"
	"github.com/shinofara/stand/zip"
)

type Args struct {
	Directory string
	ZipName   string
}

func main() {
	var args Args

	flag.StringVar(&args.Directory, "d", "", "-d target directory path")
	flag.StringVar(&args.ZipName, "zip", "", "-zip zip name")
	// コマンドライン引数を解析
	flag.Parse()

	if args.Directory == "" {
		panic("-d is empty")
	}

	if args.ZipName == "" {
		panic("-zip is empty")
	}

	err := zip.Compress(args.Directory, args.ZipName)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Created %s", args.ZipName)
}
