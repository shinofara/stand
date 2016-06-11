package main

import (
	"flag"
	"fmt"
	"github.com/shinofara/stand/find"
)

type Args struct {
	Directory string
}

func main() {
	var args Args

	// -fオプション flag.Arg(0)だとファイル名が展開されてしまうようなので
	flag.StringVar(&args.Directory, "d", "", "-d target directory path")
	// コマンドライン引数を解析
	flag.Parse()

	if args.Directory == "" {
		panic("-d is empty")
	}

	files, _ := find.All(args.Directory)

	for _, name := range files {
		fmt.Println(name)
	}
}
