package model

// Health Check Struct
type HealthCheck struct {
	Data string
}

// Book Struct
type Book struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Interface to work with Book
type BookController interface {
	Add(Book) error
	Update(string, Book) error
	Delete(string) error
	GetBook(string) (Book, error)
	GetBooks() ([]Book, error)
}
