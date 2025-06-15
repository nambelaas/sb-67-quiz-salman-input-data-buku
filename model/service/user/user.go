package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sb-67-go-quiz-salman-input-data-buku/model/repository/user"
)

type UserServiceInterface interface {
	RegisterUser(ctx *gin.Context) error
	LoginUser(ctx *gin.Context) (any, error)
	UpdateUser(ctx *gin.Context) error
}

type UserService struct {
	repo user.UserRepositoryInterface
}
