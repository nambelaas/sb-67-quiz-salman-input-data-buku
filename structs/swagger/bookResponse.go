package swagger

import "github.com/sb-67-go-quiz-salman-input-data-buku/structs"

type ResponseCreateBook struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

type ResponseGetAllBook struct {
	Success bool           `json:"Success"`
	Message string         `json:"Message"`
	Data    []structs.Book `json:"Data"`
}

type ResponseGetBookById struct {
	Success bool         `json:"Success"`
	Message string       `json:"Message"`
	Data    structs.Book `json:"Data"`
}

type ResponseDeleteBook struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}
