package location

import (
	"os"
	"testing"

	"github.com/shinofara/stand/config"
)

func TestReplacePattern(t *testing.T) {
	e := "test/*"
	a := replacePattern("/test")

	if e != a {
		t.Error("Must be equal")
	}

	e = "test/*"
	a = replacePattern("/test/")

	if e != a {
		t.Error("Must be equal")
	}

	e = "test/test2/*"
	a = replacePattern("test/test2")

	if e != a {
		t.Error("Must be equal")
	}

	e = "test/test2/*"
	a = replacePattern("test/test2/")

	if e != a {
		t.Error("Must be equal")
	}
}

func TestMakeUploadPath(t *testing.T) {
	s := &S3{
		storageCfg: &config.StorageConfig{
			Path: "/path/to/samples",
		},
	}

	file, _ := os.Open("./testdata/sample1.txt")
	defer file.Close()

	path := s.makeUploadPath(file)
	if path != "/path/to/samples/sample1.txt" {
		t.Error("Must be equal")
	}
}
