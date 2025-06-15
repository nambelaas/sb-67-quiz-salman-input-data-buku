package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sb-67-go-quiz-salman-input-data-buku/helper"
	repository "github.com/sb-67-go-quiz-salman-input-data-buku/model/repository/category"
	service "github.com/sb-67-go-quiz-salman-input-data-buku/model/service/category"
)

var (
	categoryRepo    = repository.NewCategoryRepository()
	categoryService = service.NewCategoryService(categoryRepo)
)

// @Summary Create Category
// @Description Endpoint untuk membuat category
// @Tags category
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer <jwt>"
// @Param Data  body  structs.Category  true  "Data category"
// @Success 200 {object} swagger.ResponseCreateCategory
// @Failure 400 {object} swagger.ResponseCreateCategory
// @Router /api/categories/ [post]
func CreateCategory(ctx *gin.Context) {
	err := categoryService.CreateCategory(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessage(ctx, "berhasil menambahkan category")
}

// @Summary Get All Category
// @Description Endpoint untuk menampilkan semua data category
// @Tags category
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer <jwt>"
// @Success 200 {object} swagger.ResponseGetAllCategory
// @Failure 400 {object} swagger.ResponseGetAllCategory
// @Router /api/categories/ [get]
func GetAllCategory(ctx *gin.Context) {
	data, err := categoryService.GetAllCategory(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessageWithData(ctx, "berhasil mendapatkan data category", data)
}

// @Summary Get Category By Id
// @Description Endpoint untuk menampilkan data category berdasarkan category id
// @Tags category
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer <jwt>"
// @Success 200 {object} swagger.ResponseGetCategoryById
// @Failure 400 {object} swagger.ResponseGetCategoryById
// @Router /api/categories/:id [get]
func GetCategoryById(ctx *gin.Context) {
	data, err := categoryService.GetCategoryById(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessageWithData(ctx, "berhasil mendapatkan data category", data)
}

// @Summary Get Books By Category Id
// @Description Endpoint untuk menampilkan data buku berdasarkan category id
// @Tags category
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer <jwt>"
// @Success 200 {object} swagger.ResponseGetBookInCategory
// @Failure 400 {object} swagger.ResponseGetBookInCategory
// @Router /api/categories/:id/books [get]
func GetBookInCategory(ctx *gin.Context) {
	data, err := categoryService.GetBookInCategories(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	id := ctx.Param("id")
	helper.PrintSuccessResponseMessageWithData(ctx, fmt.Sprintf("berhasil mendapatkan data buku dengan category %s", id), data)
}

// @Summary Delete Category
// @Description Endpoint untuk menghapus category menggunakan category id
// @Tags category
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer <jwt>"
// @Success 200 {object} swagger.ResponseDeleteCategory
// @Failure 400 {object} swagger.ResponseDeleteCategory
// @Router /api/categories/:id [delete]
func DeleteCategory(ctx *gin.Context) {
	err := categoryService.DeleteCategory(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessage(ctx, "berhasil menghapus data category")
}
