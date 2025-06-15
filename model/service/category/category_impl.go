package category

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sb-67-go-quiz-salman-input-data-buku/helper"
	helperData "github.com/sb-67-go-quiz-salman-input-data-buku/helper/data"
	errorHelper "github.com/sb-67-go-quiz-salman-input-data-buku/helper/error"
	"github.com/sb-67-go-quiz-salman-input-data-buku/model/repository/category"
	"github.com/sb-67-go-quiz-salman-input-data-buku/structs"
)

func NewCategoryService(repo category.CategoryRepositoryInterface) CategoryServiceInterface {
	return &CategoryService{
		repo,
	}
}

func (c *CategoryService) CreateCategory(ctx *gin.Context) error {
	var newCategory structs.Category
	err := ctx.ShouldBindJSON(&newCategory)
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

	exists := helperData.IsCategoryExists(newCategory.Name)
	if exists {
		return errorHelper.EncodeError("data category sudah ada")
	}

	err = c.repo.CreateCategory(newCategory, dataJwt.Username)
	if err != nil {
		return errorHelper.EncodeError(err.Error())
	}

	return nil
}

func (c *CategoryService) GetAllCategory(ctx *gin.Context) (listCategory []structs.Category, errorResult error) {
	listCategory, err := c.repo.GetAllCategory()
	if err != nil {
		return listCategory, errorHelper.EncodeError(err.Error())
	}

	return listCategory, nil
}

func (c *CategoryService) GetCategoryById(ctx *gin.Context) (dataCategory structs.Category, errorResult error) {
	id := ctx.Param("id")

	dataCategory, err := c.repo.GetCategoryById(id)
	if err != nil {
		return dataCategory, errorHelper.EncodeError(err.Error())
	}

	return dataCategory, nil
}

func (c *CategoryService) GetBookInCategories(ctx *gin.Context) (listBook []structs.Book, errorResult error) {
	id := ctx.Param("id")
	listBook, err := c.repo.GetBookInCategories(id)
	if err != nil {
		return listBook, errorHelper.EncodeError(err.Error())
	}

	return listBook, nil
}

func (c *CategoryService) DeleteCategory(ctx *gin.Context) error {
	id := ctx.Param("id")

	err := c.repo.DeleteCategory(id)
	if err != nil {
		return errorHelper.EncodeError(err.Error())
	}

	return nil
}
