package controller

import (
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

func CreateCategory(ctx *gin.Context) {
	err := categoryService.CreateCategory(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessage(ctx, "berhasil menambahkan category")
}

func GetAllCategory(ctx *gin.Context) {
	data, err := categoryService.GetAllCategory(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessageWithData(ctx, "berhasil mendapatkan data category", data)
}

func GetCategoryById(ctx *gin.Context) {
	data, err := categoryService.GetCategoryById(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessageWithData(ctx, "berhasil mendapatkan data category", data)
}

func UpdateCategory(ctx *gin.Context) {
	err := categoryService.UpdateCategory(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessage(ctx, "berhasil update data category")
}

func DeleteCategory(ctx *gin.Context) {
	err := categoryService.DeleteCategory(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessage(ctx, "berhasil menghapus data category")
}
