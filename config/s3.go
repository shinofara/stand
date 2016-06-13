package config

import (
	"os"
)

type S3Config struct {
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	Region          string `yaml:"region"`
	BucketName      string `yaml:"bucket_name"`
}

func mergeDefaultS3Config(s3Config *S3Config) *S3Config {
	defaultS3Config := &S3Config{
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		Region:          os.Getenv("AWS_DEFAULT_REGION"),
	}

	if s3Config.AccessKeyID != "" {
		defaultS3Config.AccessKeyID = s3Config.AccessKeyID
	}

	if s3Config.SecretAccessKey != "" {
		defaultS3Config.SecretAccessKey = s3Config.SecretAccessKey
	}

	if s3Config.Region != "" {
		defaultS3Config.Region = s3Config.Region
	}

	if s3Config.BucketName != "" {
		defaultS3Config.BucketName = s3Config.BucketName
	}

	return defaultS3Config
}
