package config

type StorageConfig struct {
	Type       string    `yaml:"type"`       //type of storage
	Path       string    `yaml:"path"`       //path of storage
	LifeCyrcle int64     `yaml:"life_cycle"` //generation management
	S3Config   *S3Config `yaml:"s3"`
}
