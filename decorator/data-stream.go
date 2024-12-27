package main

import "io"

// DataStream is a Component
type DataStream interface {
	read() (io.ReadCloser, error)
	write(data []byte) (int, error)
}
