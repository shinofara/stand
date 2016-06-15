package config

import (
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	expected := &Configs{
		&Config{
			Type: "dir",
			Path: "/path/to/target/1",
			StorageConfigs: []StorageConfig{
				StorageConfig{
					Type:       "local",
					Path:       "/path/to/output/1",
					LifeCyrcle: 1,
				},
			},
			CompressionConfig: &CompressionConfig{
				Prefix: "prefix1",
				Format: "zip",
			},
		},
	}

	actual, _ := Load("./testdata/test.yml")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, actual)
	}
}

func TestLoadMulti(t *testing.T) {
	expected := &Configs{
		&Config{
			Type: "dir",
			Path: "/path/to/target/1",
			StorageConfigs: []StorageConfig{
				StorageConfig{
					Type:       "local",
					Path:       "/path/to/output/1",
					LifeCyrcle: 1,
				},
			},
			CompressionConfig: &CompressionConfig{
				Prefix: "prefix1",
				Format: "zip",
			},
		},
		&Config{
			Type: "dir",
			Path: "/path/to/target/2",
			StorageConfigs: []StorageConfig{
				StorageConfig{
					Type:       "local",
					Path:       "/path/to/output/2",
					LifeCyrcle: 2,
				},
			},
			CompressionConfig: &CompressionConfig{
				Prefix: "prefix2",
				Format: "tar",
			},
		},
	}

	actual, _ := Load("./testdata/test_multi.yml")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, actual)
	}
}

func TestLoadS3(t *testing.T) {
	expected := &Configs{
		&Config{
			Type: "dir",
			Path: "/path/to/target/1",
			StorageConfigs: []StorageConfig{
				StorageConfig{
					Type:       "s3",
					Path:       "/path/to/output/1",
					LifeCyrcle: 1,
					S3Config: &S3Config{
						AccessKeyID:     "ACCESSKEY",
						SecretAccessKey: "SECRETKEY",
						Region:          "ap-northeast-1",
						BucketName:      "sample",
					},
				},
			},
			CompressionConfig: &CompressionConfig{
				Prefix: "prefix1",
				Format: "zip",
			},
		},
	}

	actual, _ := Load("./testdata/test_s3.yml")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, actual)
	}
}
