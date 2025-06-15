package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sb-67-go-quiz-salman-input-data-buku/helper"
	repository "github.com/sb-67-go-quiz-salman-input-data-buku/model/repository/book"
	service "github.com/sb-67-go-quiz-salman-input-data-buku/model/service/book"
)

var (
	bookRepo    = repository.NewBookRepository()
	bookService = service.NewBookService(bookRepo)
)

// @Summary Create Book
// @Description Endpoint untuk membuat book
// @Tags book
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer <jwt>"
// @param Data body structs.Book true "Data book"
// @Success 200 {object} swagger.ResponseDeleteBook
// @Failure 400 {object} swagger.ResponseDeleteBook
// @Router /api/books/ [post]
func CreateBook(ctx *gin.Context) {
	err := bookService.CreateBook(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessage(ctx, "berhasil menambahkan data book")
}

// @Summary Get All Book
// @Description Endpoint untuk menampilkan semua data book
// @Tags book
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer <jwt>"
// @Success 200 {object} swagger.ResponseGetAllBook
// @Failure 400 {object} swagger.ResponseGetAllBook
// @Router /api/books/ [get]
func GetAllBook(ctx *gin.Context) {
	data, err := bookService.GetAllBook(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessageWithData(ctx, "berhasil mendapatkan data book", data)
}

// @Summary Get Book By Id
// @Description Endpoint untuk menampilkan data book berdasarkan book id
// @Tags book
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer <jwt>"
// @Success 200 {object} swagger.ResponseGetBookById
// @Failure 400 {object} swagger.ResponseGetBookById
// @Router /api/books/:id [get]
func GetBookById(ctx *gin.Context) {
	data, err := bookService.GetBookById(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessageWithData(ctx, "berhasil mendapatkan data book", data)
}

// @Summary Delete Book
// @Description Endpoint untuk menghapus data book menggunakan book id
// @Tags book
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer <jwt>"
// @Success 200 {object} swagger.ResponseDeleteBook
// @Failure 400 {object} swagger.ResponseDeleteBook
// @Router /api/books/:id [delete]
func DeleteBook(ctx *gin.Context) {
	err := bookService.DeleteBook(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessage(ctx, "berhasil menghapus data book")
}
