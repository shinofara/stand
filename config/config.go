package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	TargetDir  string `yaml:"target"`     //path to backup target dir
	OutputDir  string `yaml:"output"`     //path to output dir
	ZipName    string `yaml:"zip_name"`   //zip name
	LifeCyrcle int64  `yaml:"life_cycle"` //generation management
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
