package io

// Reader defines Read method for reading data.
type Reader interface {
	Read(b []byte) (int, error)
}

// Writer defines Write method for writing data.
type Writer interface {
	Write(b []byte) (int, error)
}

type ReadWriter interface {
	Reader
	Writer
}
