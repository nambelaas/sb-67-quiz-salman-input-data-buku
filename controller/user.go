package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sb-67-go-quiz-salman-input-data-buku/helper"
	"github.com/sb-67-go-quiz-salman-input-data-buku/middleware"
	repository "github.com/sb-67-go-quiz-salman-input-data-buku/model/repository/user"
	service "github.com/sb-67-go-quiz-salman-input-data-buku/model/service/user"
)

var (
	userRepo    = repository.NewUserRepository()
	userService = service.NewUserService(userRepo)
)

func RegisterUser(ctx *gin.Context) {
	err := userService.RegisterUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessage(ctx, "berhasil mendaftarkan user")
}

func LoginUser(ctx *gin.Context) {
	result, err := userService.LoginUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	data, _ := result.(map[string]any)
	id, _ := data["id"].(int)
	username, _ := data["username"].(string)

	token := middleware.GenerateJwtToken(id, username)

	helper.PrintSuccessResponseToken(ctx, token)
}

func UpdateUser(ctx *gin.Context) {
	err := userService.UpdateUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessage(ctx, "berhasil update user")
}
