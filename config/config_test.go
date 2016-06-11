package config

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	expected := &Config{
		TargetDir:  "target",
		OutputDir:  "output",
		LifeCyrcle: 12,
		CompressionConfig: &CompressionConfig{
			Prefix: "sample",
			Format: "zip",
		},
	}

	cfg, _ := New("./testdata/test.yml")

	if !reflect.DeepEqual(expected, cfg) {
		t.Error("Must be equal")
	}
}
