package find

import (
	"testing"
)

func TestFind(t *testing.T) {
	files, _ := Find("./testdata", true, false)

	if len(files) != 7 {
		t.Errorf("length is 7, but it is %d", len(files))
	}

	files, _ = Find("./testdata", true, true)

	if len(files) != 4 {
		t.Errorf("length is 4, but it is %d", len(files))
	}

	files, _ = Find("./testdata", false, false)

	if len(files) != 3 {
		t.Errorf("length is 3, but it is %d", len(files))
	}

	files, _ = Find("./testdata", false, true)

	if len(files) != 2 {
		t.Errorf("length is 2, but it is %d", len(files))
	}

	files, _ = Find("./testdata/sub1", false, false)

	if len(files) != 2 {
		t.Errorf("length is 2, but it is %d", len(files))
	}
}
