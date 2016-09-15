package main

import (
	"context"

	"github.com/shinofara/stand/config"
	"github.com/shinofara/stand/coordinator"

	"flag"

	"github.com/uber-go/zap"
)

var (
	flCfgPath    string
	flOutputPath string
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
		if err := c.Run(); err != nil {
			logger.Fatal(err.Error())
		}
	}
}

//initCfg initialize configs
func initCfg() (*config.Configs, error) {
	flag.StringVar(&flCfgPath, "c", "", "path to config yaml")
	flag.StringVar(&flCfgPath, "conf", "", "path to config yaml")
	flag.StringVar(&flOutputPath, "o", "", "path to output dir")
	flag.StringVar(&flOutputPath, "out", "", "path to output dir")

	flag.Parse()

	if flCfgPath == "" {
		return config.GenerateSimpleConfigs(flag.Arg(0), flOutputPath), nil
	}

	return config.Load(flCfgPath)
}
