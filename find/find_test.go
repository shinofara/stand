package find

import (
	"testing"
	"os"
)

func TestRun(t *testing.T) {
	var files []string
	middleware := func(path string, file *os.File) error {
		files = append(files, path)
		return nil
	}

	f := New(middleware, "./testdata", DeepSearchMode, NotFileOnlyMode)
	f.Run()

	if len(files) != 6 {
		t.Errorf("Must be equal 6, but it is %d", len(files))
	}

	//file only mode
	files = []string{}
	f = New(middleware, "./testdata", DeepSearchMode, FileOnlyMode)
	f.Run()

	if len(files) != 4 {
		t.Errorf("Must be equal 4, but it is %d", len(files))
	}

	//file only mode & no deep search mode
	files = []string{}
	f = New(middleware, "./testdata", NotDeepSearchMode, FileOnlyMode)
	f.Run()

	if len(files) != 2 {
		t.Errorf("Must be equal 2, but it is %d", len(files))
	}
}
