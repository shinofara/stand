package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configs []*Config

type Config struct {
	Type              string             `yaml:"type"` //type of target object
	Path              string             `yaml:"path"` //path to backup target dir
	CompressionConfig *CompressionConfig `yaml:"compression"`
	StorageConfig     StorageConfig      `yaml:"storage"`
}

type CompressionConfig struct {
	Prefix string `yaml:"prefix"` // prefix of the compression file name.
	Format string `yaml:"format"` // format of the compression file.
}

func New(path string) (*Configs, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfgs, err := loadSingleConfig(buf)
	if err != nil {
		cfgs, err = loadMultiConfig(buf)
		if err != nil {
			return nil, err
		}
	}

	return cfgs, nil
}

func loadSingleConfig(buf []byte) (*Configs, error) {
	var cfg Config
	if err := yaml.Unmarshal(buf, &cfg); err != nil {
		return nil, err
	}

	return &Configs{&cfg}, nil
}

func loadMultiConfig(buf []byte) (*Configs, error) {
	var cfgs Configs
	if err := yaml.Unmarshal(buf, &cfgs); err != nil {
		return nil, err
	}

	return &cfgs, nil
}
