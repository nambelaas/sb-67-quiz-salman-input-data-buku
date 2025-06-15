package category

import (
	"github.com/gin-gonic/gin"
	"github.com/sb-67-go-quiz-salman-input-data-buku/model/repository/category"
	"github.com/sb-67-go-quiz-salman-input-data-buku/structs"
)

type CategoryServiceInterface interface {
	CreateCategory(ctx *gin.Context) error
	GetAllCategory(ctx *gin.Context) (listCategory []structs.Category, errorResult error)
	GetCategoryById(ctx *gin.Context) (dataCategory structs.Category, errorResult error)
	UpdateCategory(ctx *gin.Context) error
	DeleteCategory(ctx *gin.Context) error
}

type CategoryService struct {
	repo category.CategoryRepositoryInterface
}
