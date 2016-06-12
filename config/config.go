package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configs []*Config

type Config struct {
	TargetDir         string             `yaml:"target"`     //path to backup target dir
	Location          string             `yaml:"location"`   //backup location, local or s3
	OutputDir         string             `yaml:"output"`     //path to output dir
	LifeCyrcle        int64              `yaml:"life_cycle"` //generation management
	CompressionConfig *CompressionConfig `yaml:"compression"`
	S3Config          *S3Config          `yaml:"s3"`
}

type CompressionConfig struct {
	Prefix string `yaml:"prefix"` // prefix of the compression file name.
	Format string `yaml:"format"` // format of the compression file.
}

type S3Config struct {
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	Region          string `yaml:"region"`
	BucketName      string `yaml:"bucket_name"`
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
