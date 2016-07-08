package main

import (
	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/coordinator"

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
	ctx := context.WithValue(
		context.Background(),
		"logger",
		logger)

	cfgs, err := initCfg()
	if err != nil {
		logger.Fatal(err.Error())
	}

	for _, cfg := range *cfgs {
		c := coordinator.New(ctx, cfg)
		if err := c.Perform(); err != nil {
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
