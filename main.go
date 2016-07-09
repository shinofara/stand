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
		return config.GenerateSimpleConfigs(flag.Arg(0), *flOutputPath), nil
	}

	return config.Load(*flCfgPath)
}
