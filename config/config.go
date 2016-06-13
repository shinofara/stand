package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	TYPE_FILE = "file"
	TYPE_DIR  = "dir"
)

type Configs []*Config

type Config struct {
	Type              string             `yaml:"type"` //type of target object
	Path              string             `yaml:"path"` //path to backup target dir
	CompressionConfig *CompressionConfig `yaml:"compression"`
	StorageConfig     StorageConfig      `yaml:"storage"`
}

func Load(path string) (*Configs, error) {
	cfgs, err := loadYAML(path)
	if err != nil {
		return nil, err
	}

	for _, cfg := range *cfgs {
		if cfg.StorageConfig.Type == "s3" {
			cfg.StorageConfig.S3Config = mergeDefaultS3Config(cfg.StorageConfig.S3Config)
		}
		if cfg.Type == TYPE_DIR {
			cfg.CompressionConfig = mergeDefaultCompressionConfig(cfg.CompressionConfig)
		}
	}

	return cfgs, nil
}

func loadYAML(path string) (*Configs, error) {
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
