package models

import "bytes"

// OutputFile represents a file to be generated.
type OutputFile struct {
	Path string
	Link string
	Buf  *bytes.Buffer
}
