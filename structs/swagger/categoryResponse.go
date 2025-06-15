package swagger

import "github.com/sb-67-go-quiz-salman-input-data-buku/structs"

type ResponseCreateCategory struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

type ResponseGetAllCategory struct {
	Success bool               `json:"Success"`
	Message string             `json:"Message"`
	Data    []structs.Category `json:"Data"`
}

type ResponseGetCategoryById struct {
	Success bool             `json:"Success"`
	Message string           `json:"Message"`
	Data    structs.Category `json:"Data"`
}

type ResponseGetBookInCategory struct {
	Success bool           `json:"Success"`
	Message string         `json:"Message"`
	Data    []structs.Book `json:"Data"`
}

type ResponseDeleteCategory struct {
	Success bool           `json:"Success"`
	Message string         `json:"Message"`
}
