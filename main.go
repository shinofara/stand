package main

import (
	"github.com/shinofara/stand/archiver"
	"github.com/shinofara/stand/backup"
	"github.com/shinofara/stand/config"

	flag "github.com/docker/docker/pkg/mflag"
	"github.com/uber-go/zap"
	"golang.org/x/net/context"
)

var (
	flCfgPath    = flag.String([]string{"c", "-conf"}, "", "path to config yaml")
	flOutputPath = flag.String([]string{"o", "-out"}, "", "path to output dir")
	logger       = zap.NewJSON()
)

func main() {
	// For repeatable tests, pretend that it's always 1970.
	ctx := context.Background()
	ctx = context.WithValue(ctx, "logger", logger)

	cfgs, err := initCfg()
	if err != nil {
		logger.Fatal(err.Error())
	}

	for _, cfg := range *cfgs {
		var uploadFileName string
		var err error

		switch cfg.Type {
		case "dir":

			a := archiver.New(ctx, cfg)
			uploadFileName, err = a.Archive()
			if err != nil {
				logger.Fatal(err.Error())
			}
		case "file":
			uploadFileName = cfg.Path
		default:
			logger.Fatal("upload target type is not found")
		}

		b := backup.New(ctx, cfg)
		if err := b.Exec(uploadFileName); err != nil {
			logger.Fatal(err.Error())
		}
	}
}

//initCfg initialize configs
func initCfg() (*config.Configs, error) {
	flag.Parse()

	if *flCfgPath == "" {
		return loadOption()
	}

	return loadCfg(*flCfgPath)

}

func loadCfg(path string) (*config.Configs, error) {
	cfgs, err := config.Load(path)
	if err != nil {
		return nil, err
	}

	return cfgs, nil
}

func loadOption() (*config.Configs, error) {
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

	return &config.Configs{cfg}, nil
}
