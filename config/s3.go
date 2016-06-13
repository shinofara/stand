package config

import (
	"os"
)

var (
	defaultS3Config = &S3Config{}
)

func init() {
	defaultS3Config = &S3Config{
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		Region:          os.Getenv("AWS_DEFAULT_REGION"),
	}
}

type S3Config struct {
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	Region          string `yaml:"region"`
	BucketName      string `yaml:"bucket_name"`
}

func mergeDefaultS3Config(s3Config *S3Config) {
	if s3Config.AccessKeyID == "" {
		s3Config.AccessKeyID = defaultS3Config.AccessKeyID
	}
	if s3Config.SecretAccessKey == "" {
		s3Config.SecretAccessKey = defaultS3Config.SecretAccessKey
	}
	if s3Config.Region == "" {
		s3Config.Region = defaultS3Config.Region
	}
}
