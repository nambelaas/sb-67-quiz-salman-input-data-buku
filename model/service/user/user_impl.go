package user

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sb-67-go-quiz-salman-input-data-buku/helper"
	helperData "github.com/sb-67-go-quiz-salman-input-data-buku/helper/data"
	errorHelper "github.com/sb-67-go-quiz-salman-input-data-buku/helper/error"
	"github.com/sb-67-go-quiz-salman-input-data-buku/model/repository/user"
	"github.com/sb-67-go-quiz-salman-input-data-buku/structs"
)

func NewUserService(repo user.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		repo,
	}
}

func (u *UserService) RegisterUser(ctx *gin.Context) error {
	var newUser structs.AuthRequest
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		result := errorHelper.ValidationCheck(err)
		if result != nil {
			encode, _ := json.Marshal(result)
			return errors.New(string(encode))
		}

		return errorHelper.EncodeError(err.Error())
	}

	exists := helperData.IsUserExists(newUser)
	if exists {
		return errorHelper.EncodeError("data user sudah ada")
	}

	err = u.repo.RegisterUser(newUser)
	if err != nil {
		return errorHelper.EncodeError(err.Error())
	}

	return nil
}

func (u *UserService) LoginUser(ctx *gin.Context) (any, error) {
	var newUser structs.AuthRequest
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		return nil, errorHelper.EncodeError(err.Error())
	}

	id, err := u.repo.LoginUser(newUser.Username, newUser.Password)
	if err != nil {
		return nil, errorHelper.EncodeError(err.Error())
	}

	idInt := id.(int)
	return map[string]any{
		"id":       idInt,
		"username": newUser.Username,
	}, nil
}

func (u *UserService) UpdateUser(ctx *gin.Context) error {
	var updateUser structs.AuthRequest
	err := ctx.ShouldBindJSON(&updateUser)
	if err != nil {
		return errorHelper.EncodeError(err.Error())
	}

	var id = ctx.Param("id")
	dataJwt, err := helper.GetJwtData(ctx)
	if err != nil {
		return errorHelper.EncodeError(err.Error())
	}

	err = u.repo.UpdateUser(id, updateUser, dataJwt.Username)
	if err != nil {
		return errorHelper.EncodeError(err.Error())
	}

	return nil
}
