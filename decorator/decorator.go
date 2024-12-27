/*
*
Decorator pattern is a design pattern that allows behavior to be added to an individual object dynamically.
It provides a way to wrap an existing object and extend its functionality without modifying its structure.

In this example, we will create a simple decorator pattern to add logging functionality to a simple calculator.

The Calculator interface defines a single method for performing calculations.
*/
package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	bs := &BufferedDataStream{
		data: bytes.Buffer{},
	}

	es := &EncryptionDecorator{
		DataStream:    bs,
		encryptionKey: []byte("1234567890123456"),
	}

	_, err := es.write([]byte("Hello Secure world"))
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}

	rc, err := es.read()
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}
	defer func(rc io.ReadCloser) {
		err := rc.Close()
		if err != nil {
			fmt.Println("Error closing ReadCloser:", err)
		}
	}(rc)

	data, err := io.ReadAll(rc)
	if err != nil {
		fmt.Println("Error reading data from ReadCloser:", err)
		return
	}

	fmt.Println("Original data:", string(data))
	fmt.Println("Encrypted data:", string(bs.data.Bytes()))
	fmt.Println("Decrypted data:", string(data))
}
