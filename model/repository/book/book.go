package book

import "github.com/sb-67-go-quiz-salman-input-data-buku/structs"

type BookRepositoryInterface interface {
	CreateBook(book structs.Book, responsibleUser string) error
	GetAllBook() (result []structs.Book, err error)
	GetBookById(id string) (result structs.Book, err error)
	DeleteBook(id string) error
}

type BookRepository struct{}
