package location

import (
	"testing"
	"github.com/shinofara/stand/config"
	"path/filepath"
)

func TestCleanRun(t *testing.T) {
	storageCfg := &config.StorageConfig{
		Type: "local",
		Path: "./testdata",
		LifeCyrcle: 1,
	}
	clean := NewCLean(storageCfg)
	if err := clean.Run(); err != nil {
		t.Fatal(err.Error())
	}

	if len(clean.targets) != 2 {
		t.Errorf("Must equal 2, but it is %d", len(clean.targets))
	}

	for _, file := range clean.targets {
		if filepath.Ext(file.Path) != ".zip" && filepath.Ext(file.Path) != ".gz" {
			t.Errorf("Do not disappear other than tar and gz. %s", file.Path)
		}
	}
}
