package find

import (
	"testing"
	"os"
	"sort"
)

func TestRun(t *testing.T) {
	var files []string
	middleware := func(path string, file *os.File) error {
		files = append(files, path)
		return nil
	}

	f := New(middleware, "./testdata", DeepSearchMode, NotFileOnlyMode)
	f.Run()

	if len(files) != 7 {
		t.Errorf("Must be equal 7, but it is %d", len(files))
	}

	//file only mode
	files = []string{}
	f = New(middleware, "./testdata", DeepSearchMode, FileOnlyMode)
	f.Run()

	if len(files) != 5 {
		t.Errorf("Must be equal 5, but it is %d", len(files))
	}

	//file only mode & no deep search mode
	files = []string{}
	f = New(middleware, "./testdata", NotDeepSearchMode, FileOnlyMode)
	f.Run()

	if len(files) != 3 {
		t.Errorf("Must be equal 3, but it is %d", len(files))
	}
}

func TestGetFiles(t *testing.T) {
	files, _ := GetFiles("./testdata")

	if len(files) != 3 {
		t.Error("Must be equal 3")
	}

	names := []string{"a.txt", "d.txt", "e.txt"}
	for key, file := range files {
		if names[key] != file.Path {
			t.Errorf("Must equal %s and %s" , names[key], file.Path)
		}
	}

	sort.Sort(files)
	names = []string{"e.txt", "a.txt", "d.txt"}
	for key, file := range files {
		if names[key] != file.Path {
			t.Errorf("Must equal %s and %s" , names[key], file.Path)
		}
	}	
}
