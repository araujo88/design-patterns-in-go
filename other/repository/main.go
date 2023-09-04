package main

import "fmt"

// Book represents a book in a library
type Book struct {
	ID     string
	Title  string
	Author string
}

// BookRepository interface for persisting books
type BookRepository interface {
	Add(Book) error
	Remove(string) error
	FindById(string) (*Book, error)
	FindAll() ([]*Book, error)
}

// InMemoryBookRepository is an in-memory implementation of BookRepository
type InMemoryBookRepository struct {
	books map[string]*Book
}

// NewInMemoryBookRepository initializes a new InMemoryBookRepository
func NewInMemoryBookRepository() *InMemoryBookRepository {
	return &InMemoryBookRepository{
		books: make(map[string]*Book),
	}
}

func (r *InMemoryBookRepository) Add(b Book) error {
	r.books[b.ID] = &b
	return nil
}

func (r *InMemoryBookRepository) Remove(id string) error {
	delete(r.books, id)
	return nil
}

func (r *InMemoryBookRepository) FindById(id string) (*Book, error) {
	if book, ok := r.books[id]; ok {
		return book, nil
	}
	return nil, fmt.Errorf("Book not found")
}

func (r *InMemoryBookRepository) FindAll() ([]*Book, error) {
	var books []*Book
	for _, book := range r.books {
		books = append(books, book)
	}
	return books, nil
}

func main() {
	repo := NewInMemoryBookRepository()

	// Add books
	repo.Add(Book{ID: "1", Title: "1984", Author: "George Orwell"})
	repo.Add(Book{ID: "2", Title: "Brave New World", Author: "Aldous Huxley"})

	// Find a book by ID
	book, err := repo.FindById("1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found book:", *book)
	}

	// Find all books
	allBooks, err := repo.FindAll()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, book := range allBooks {
			fmt.Println(*book)
		}
	}

	// Remove a book
	err = repo.Remove("1")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Try to find the removed book
	book, err = repo.FindById("1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found book:", *book)
	}
}
