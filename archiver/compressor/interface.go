package compressor

import (
	"io"
)

type Compressor interface {
	Compress(io.Writer, string, []string) error
}
