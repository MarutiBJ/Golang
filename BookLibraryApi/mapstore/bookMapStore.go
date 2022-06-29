package mapstore

import (
	"errors"

	"basicwebapp/model"
)

// Library to store Book
type BookLibrary struct {
	Library map[string]model.Book
}

// Function to Get Instance of BookController interface
func NewBookLibrary() *BookLibrary {
	return &BookLibrary{Library: make(map[string]model.Book, 0)}
}

// BookLibrary implements BookController
func (library *BookLibrary) Add(book model.Book) error {

	if _, ok := library.Library[book.Id]; ok {
		return errors.New("Book Already Exist")
	}

	library.Library[book.Id] = book
	return nil
}

func (library BookLibrary) Update(id string, book model.Book) error {
	if _, ok := library.Library[id]; !ok {
		return errors.New("Book Not Found")
	}
	library.Library[id] = book
	return nil
}

func (library BookLibrary) Delete(id string) error {
	if _, ok := library.Library[id]; !ok {
		return errors.New("Book Not Found")
	}
	delete(library.Library, id)
	return nil

}

func (library BookLibrary) GetBook(id string) (model.Book, error) {
	if book, ok := library.Library[id]; !ok {
		return book, errors.New("Book Not Found")
	} else {
		return book, nil
	}
}

func (library BookLibrary) GetBooks() ([]model.Book, error) {
	books := make([]model.Book, 0, len(library.Library))

	if len(library.Library) == 0 {
		return books, errors.New("No Book Found")
	}

	for _, book := range library.Library {
		books = append(books, book)
	}

	return books, nil
}
