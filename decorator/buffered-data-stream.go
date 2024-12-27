package main

import (
	"bytes"
	"io"
)

type BufferedDataStream struct {
	data bytes.Buffer
}

func (s *BufferedDataStream) read() (io.ReadCloser, error) {
	return io.NopCloser(bytes.NewReader(s.data.Bytes())), nil
}

func (s *BufferedDataStream) write(data []byte) (int, error) {
	return s.data.Write(data)
}
