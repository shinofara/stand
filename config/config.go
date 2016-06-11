package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	TargetDir         string             `yaml:"target"`     //path to backup target dir
	OutputDir         string             `yaml:"output"`     //path to output dir
	LifeCyrcle        int64              `yaml:"life_cycle"` //generation management
	CompressionConfig *CompressionConfig `yaml:"compression"`
}

type CompressionConfig struct {
	Prefix string `yaml:"prefix"` // prefix of the compression file name.
	Format string `yaml:"format"` // format of the compression file.
}

func New(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
