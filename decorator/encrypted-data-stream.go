package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io"
)

type EncryptionDecorator struct {
	DataStream
	encryptionKey []byte
}

func (e *EncryptionDecorator) write(data []byte) (int, error) {
	encryptedData, err := encrypt(string(data), e.encryptionKey)
	if err != nil {
		return 0, fmt.Errorf("encryption failed: %d", err)
	}
	return e.DataStream.write([]byte(encryptedData))
}

func (e *EncryptionDecorator) read() (io.ReadCloser, error) {
	rc, err := e.DataStream.read()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	var buff bytes.Buffer
	if _, err := io.Copy(&buff, rc); err != nil {
		return nil, err
	}
	decryptedData, err := decrypt(buff.String(), e.encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}
	return io.NopCloser(bytes.NewReader([]byte(decryptedData))), nil
}

func decrypt(data string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	plaintext := make([]byte, len(ciphertext))
	iv := make([]byte, aes.BlockSize)
	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(plaintext, ciphertext)

	padding := int(plaintext[len(plaintext)-1])
	if padding > aes.BlockSize || padding <= 0 {
		return "", fmt.Errorf("invalid padding")
	}
	plaintext = plaintext[:len(plaintext)-padding]
	return string(plaintext), err
}

func encrypt(s string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	plaintext := []byte(s)
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	for i := 0; i < padding; i++ {
		plaintext = append(plaintext, byte(padding))
	}
	ciphertext := make([]byte, len(plaintext))
	iv := make([]byte, aes.BlockSize)
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
