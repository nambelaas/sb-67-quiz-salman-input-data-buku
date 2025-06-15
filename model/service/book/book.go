package book

import (
	"github.com/gin-gonic/gin"
	"github.com/sb-67-go-quiz-salman-input-data-buku/model/repository/book"
	"github.com/sb-67-go-quiz-salman-input-data-buku/structs"
)

type BookServiceInterface interface {
	CreateBook(ctx *gin.Context) error
	GetAllBook(ctx *gin.Context) (listBook []structs.Book, errorResult error)
	GetBookById(ctx *gin.Context) (dataBook structs.Book, errorResult error)
	DeleteBook(ctx *gin.Context) error
}

type BookService struct {
	repo book.BookRepositoryInterface
}
