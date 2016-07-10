package compressor

import (
	"os"
	"testing"

	"github.com/shinofara/stand/find"	
)

func TestCreateFileHeader(t *testing.T) {
	info, _ := os.Stat("testdata/sample.txt")

	fInfo := find.File{
		Info: info,
		Path: "test/sample.txt",
	}
	
	zipHeader, _ := createZipFileHeader(fInfo)

	if info.ModTime() == zipHeader.ModTime() {
		t.Error("Must not be equall")
	}

	currentTime := info.ModTime().Format("2006-01-02 15:04:06")
	zipTime := zipHeader.ModTime().Format("2006-01-02 15:04:06")

	if currentTime != zipTime {
		t.Error("Must be equal")
	}
}
