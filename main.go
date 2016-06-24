package main

import (
	"fmt"
	"time"

	"github.com/shinofara/stand/archiver"
	"github.com/shinofara/stand/backup"
	"github.com/shinofara/stand/config"

	flag "github.com/docker/docker/pkg/mflag"
	"log"
)

const (
	TIME_FORMAT = "20060102150405"
)

var (
	flCfgPath    = flag.String([]string{"c", "-conf"}, "", "path to config yaml")
	flOutputPath = flag.String([]string{"o", "-out"}, "", "path to output dir")
)

func main() {
	cfgs := initCfg()

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

//initCfg initialize configs
func initCfg() *config.Configs {
	flag.Parse()

	if *flCfgPath == "" {
		return loadOption()
	}

	return loadCfg(*flCfgPath)
}

func loadCfg(path string) *config.Configs {
	cfgs, err := config.Load(path)
	if err != nil {
		panic(err)
	}

	return cfgs
}

func loadOption() *config.Configs {
	cfg := &config.Config{
		Type: "dir",
		Path: flag.Arg(0),
		CompressionConfig: &config.CompressionConfig{
			Format: "zip",
		},
		StorageConfigs: []config.StorageConfig{
			config.StorageConfig{
				Type: "local",
				Path: *flOutputPath,
			},
		},
	}

	return &config.Configs{cfg}
}
