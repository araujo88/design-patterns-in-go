package main

import (
	"fmt"
	"io/ioutil"
)

// Writer interface, the 'Service' in our example
type Writer interface {
	Write([]byte) error
}

// ConsoleWriter writes to the console, one of the 'ConcreteService'
type ConsoleWriter struct{}

func (cw *ConsoleWriter) Write(data []byte) error {
	_, err := fmt.Println(string(data))
	return err
}

// FileWriter writes to a file, another 'ConcreteService'
type FileWriter struct {
	filename string
}

func (fw *FileWriter) Write(data []byte) error {
	return ioutil.WriteFile(fw.filename, data, 0644)
}

// Author is the 'Client'
type Author struct {
	writer Writer // Author depends on a Writer to do the writing
}

// NewAuthor creates a new author with a writer
func NewAuthor(writer Writer) *Author {
	return &Author{writer}
}

// WriteArticle uses the injected writer to write an article
func (a *Author) WriteArticle(article string) error {
	return a.writer.Write([]byte(article))
}

func main() {
	// Inject ConsoleWriter dependency into Author
	author1 := NewAuthor(&ConsoleWriter{})
	_ = author1.WriteArticle("Article to console")

	// Inject FileWriter dependency into Author
	author2 := NewAuthor(&FileWriter{filename: "article.txt"})
	_ = author2.WriteArticle("Article to file")
}
