package book

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sb-67-go-quiz-salman-input-data-buku/helper"
	helperData "github.com/sb-67-go-quiz-salman-input-data-buku/helper/data"
	errorHelper "github.com/sb-67-go-quiz-salman-input-data-buku/helper/error"
	"github.com/sb-67-go-quiz-salman-input-data-buku/model/repository/book"
	"github.com/sb-67-go-quiz-salman-input-data-buku/structs"
)

func NewBookService(repo book.BookRepositoryInterface) BookServiceInterface {
	return &BookService{
		repo,
	}
}

func (c *BookService) CreateBook(ctx *gin.Context) error {
	var newBook structs.Book
	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		result := errorHelper.ValidationCheck(err)
		if result != nil {
			encode, _ := json.Marshal(result)
			return errors.New(string(encode))
		}

		return errorHelper.EncodeError(err.Error())
	}

	dataJwt, err := helper.GetJwtData(ctx)
	if err != nil {
		return errorHelper.EncodeError(err.Error())
	}

	exists := helperData.IsBookExists(newBook.Title)
	if exists {
		return errorHelper.EncodeError("data book sudah ada")
	}

	err = c.repo.CreateBook(newBook, dataJwt.Username)
	if err != nil {
		return errorHelper.EncodeError(err.Error())
	}

	return nil
}

func (c *BookService) GetAllBook(ctx *gin.Context) (listBook []structs.Book, errorResult error) {
	listBook, err := c.repo.GetAllBook()
	if err != nil {
		return listBook, errorHelper.EncodeError(err.Error())
	}

	return listBook, nil
}

func (c *BookService) GetBookById(ctx *gin.Context) (dataBook structs.Book, errorResult error) {
	id := ctx.Param("id")

	dataBook, err := c.repo.GetBookById(id)
	if err != nil {
		return dataBook, errorHelper.EncodeError(err.Error())
	}

	return dataBook, nil
}

func (c *BookService) DeleteBook(ctx *gin.Context) error {
	id := ctx.Param("id")

	err := c.repo.DeleteBook(id)
	if err != nil {
		return errorHelper.EncodeError(err.Error())
	}

	return nil
}
