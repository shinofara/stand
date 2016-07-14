package find

import (
	"os"
	"sort"
	"testing"
	"time"
)

func prepare() {
	dirs := []string{
		"testdata/dir1",
		"testdata/dir2/dir3",
	}

	files := []string{
		"testdata/a.txt",
		"testdata/c.txt",
		"testdata/b.txt",
		"testdata/dir1/a.txt",
		"testdata/dir1/b.txt",
		"testdata/dir2/dir3/a.txt",
	}

	for _, dir := range dirs {
		os.MkdirAll(dir, os.ModePerm)
	}

	for _, file := range files {
		f, _ := os.Create(file)
		f.Close()
	}
	os.Chtimes(
		"testdata/a.txt",
		time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
		time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
	)

	os.Chtimes(
		"testdata/b.txt",
		time.Date(2014, time.December, 29, 12, 13, 24, 0, time.UTC),
		time.Date(2014, time.December, 29, 12, 13, 24, 0, time.UTC),
	)

	os.Chtimes(
		"testdata/c.txt",
		time.Date(2014, time.December, 30, 12, 13, 24, 0, time.UTC),
		time.Date(2014, time.December, 30, 12, 13, 24, 0, time.UTC),
	)
}

func clean(t *testing.T) {
	dirs := []string{
		"testdata/dir1",
		"testdata/dir2",
		"testdata/a.txt",
		"testdata/c.txt",
		"testdata/b.txt",
	}
	for _, dir := range dirs {
		if err := os.RemoveAll(dir); err != nil {
			t.Error(dir)
		}
	}
}

func TestRun(t *testing.T) {
	prepare()
	defer clean(t)

	var files []string
	middleware := func(path string, file *os.File) error {
		files = append(files, path)
		return nil
	}

	f := New(middleware, "./testdata", DeepSearchMode, NotFileOnlyMode)
	f.Run()

	if len(files) != 9 {
		t.Errorf("Must be equal 9, but it is %d", len(files))
	}

	//file only mode
	files = []string{}
	f = New(middleware, "./testdata", DeepSearchMode, FileOnlyMode)
	f.Run()

	if len(files) != 6 {
		t.Errorf("Must be equal 6, but it is %d", len(files))
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
	prepare()
	defer clean(t)

	files, _ := GetFiles("./testdata")

	if len(files) != 3 {
		t.Error("Must be equal 3")
	}

	names := []string{"a.txt", "b.txt", "c.txt"}
	for key, file := range files {
		if names[key] != file.Path {
			t.Errorf("Must equal %s and %s", names[key], file.Path)
		}
	}

	//降順
	sort.Sort(files)
	names = []string{"a.txt", "c.txt", "b.txt"}
	for key, file := range files {
		if names[key] != file.Path {
			t.Errorf("Must equal %s and %s", names[key], file.Path)
		}
	}
}
