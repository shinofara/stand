package location

import (
	"testing"
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
