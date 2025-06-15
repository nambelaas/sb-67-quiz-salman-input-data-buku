package user

import "github.com/sb-67-go-quiz-salman-input-data-buku/structs"

type UserRepositoryInterface interface {
	RegisterUser(user structs.AuthRequest) error
	LoginUser(username string, password string) (any, error)
	UpdateUser(id string, user structs.AuthRequest, responsibleUser string) error
}

type UserRepository struct{}
