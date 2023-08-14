package main

import (
	"fmt"
)

// Reader is our base component interface
type Reader interface {
	Read() string
}

// FileReader reads from a file
type FileReader struct{}

func (fr *FileReader) Read() string {
	return "data from file"
}

// Buffered reader as a decorator
type BufferedReader struct {
	reader Reader
}

func (br *BufferedReader) Read() string {
	data := br.reader.Read()
	return "buffered " + data
}

// Decrypt reader as another decorator
type DecryptReader struct {
	reader Reader
}

func (dr *DecryptReader) Read() string {
	data := dr.reader.Read()
	return "decrypted " + data
}

func main() {
	fileReader := &FileReader{}

	// Buffered read
	bufferedReader := &BufferedReader{reader: fileReader}
	fmt.Println(bufferedReader.Read())

	// Buffered and decrypted read
	decryptBufferedReader := &DecryptReader{reader: bufferedReader}
	fmt.Println(decryptBufferedReader.Read())
}
