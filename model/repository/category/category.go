package category

import "github.com/sb-67-go-quiz-salman-input-data-buku/structs"

type CategoryRepositoryInterface interface {
	CreateCategory(category structs.Category, responsibleUser string) error
	GetAllCategory() (result []structs.Category, err error)
	GetCategoryById(id string) (result structs.Category, err error)
	UpdateCategory(id string, category structs.Category, responsibleUser string) error
	DeleteCategory(id string) error
}

type CategoryRepository struct{}
