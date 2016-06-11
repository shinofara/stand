package config

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	expected := &Config{
		TargetDir:  "target",
		OutputDir:  "output",
		ZipName:    "name",
		LifeCyrcle: 12,
	}

	cfg, _ := New("./testdata/test.yml")

	if !reflect.DeepEqual(expected, cfg) {
		t.Error("Must be equal")
	}
}
