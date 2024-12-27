package main

import (
	"io"
	"os"
)

type FileDataStream struct {
	filePath string
}

func (f *FileDataStream) read() (io.ReadCloser, error) {
	file, err := os.Open(f.filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (f *FileDataStream) write(data []byte) (int, error) {
	file, err := os.Create(f.filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(data)
}
