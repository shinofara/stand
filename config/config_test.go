package config

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	expected := &Configs{
		&Config{
			TargetDir:  "/path/to/target/1",
			OutputDir:  "/path/to/output/1",
			LifeCyrcle: 1,
			CompressionConfig: &CompressionConfig{
				Prefix: "prefix1",
				Format: "zip",
			},
		},
	}

	actual, _ := New("./testdata/test.yml")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, actual)
	}
}

func TestNewMulti(t *testing.T) {
	expected := &Configs{
		&Config{
			TargetDir:  "/path/to/target/1",
			OutputDir:  "/path/to/output/1",
			LifeCyrcle: 1,
			CompressionConfig: &CompressionConfig{
				Prefix: "prefix1",
				Format: "zip",
			},
		},
		&Config{
			TargetDir:  "/path/to/target/2",
			OutputDir:  "/path/to/output/2",
			LifeCyrcle: 2,
			CompressionConfig: &CompressionConfig{
				Prefix: "prefix2",
				Format: "tar",
			},
		},
	}

	actual, _ := New("./testdata/test_multi.yml")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, actual)
	}
}
