package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type JsonFileStream struct {
	DataStream
}

func (j *JsonFileStream) read() (io.ReadCloser, error) {
	rc, err := j.DataStream.read()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	var buff bytes.Buffer
	if _, err := io.Copy(&buff, rc); err != nil {
		return nil, err
	}
	var parsed map[string]interface{}
	if err := json.Unmarshal(buff.Bytes(), &parsed); err != nil {
		return nil, fmt.Errorf("json.Unmarshal failed: %v", err)
	}

	return io.NopCloser(bytes.NewReader(buff.Bytes())), nil
}

func (j *JsonFileStream) write(data []byte) (int, error) {
	var temp map[string]interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return 0, fmt.Errorf("invalid Json: %w", err)
	}
	return j.DataStream.write(data)
}
