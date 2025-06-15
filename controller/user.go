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

// @Summary Register User
// @Description Endpoint untuk registrasi user baru
// @Tags user
// @Accept  json
// @Produce  json
// @Param   Data  body  structs.AuthRequest  true  "Data user"
// @Success 200 {object} swagger.ResponseRegister
// @Failure 400 {object} swagger.ResponseRegister
// @Router /api/users/register [post]
func RegisterUser(ctx *gin.Context) {
	err := userService.RegisterUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessage(ctx, "berhasil mendaftarkan user")
}

// @Summary Login User
// @Description Endpoint untuk login user
// @Tags user
// @Accept  json
// @Produce  json
// @Param   Data  body  structs.AuthRequest  true  "Data user"
// @Success 200 {object} swagger.ResponseLogin
// @Failure 400 {object} swagger.ResponseLogin
// @Router /api/users/login [post]
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

// @Summary Update User
// @Description Endpoint untuk update data user
// @Tags user
// @Accept  json
// @Produce  json
// @Param   Data  body  structs.AuthRequest  true  "Data user"
// @Success 200 {object} swagger.ResponseUpdate
// @Failure 400 {object} swagger.ResponseUpdate
// @Router /api/users/update/:id [put]
func UpdateUser(ctx *gin.Context) {
	err := userService.UpdateUser(ctx)
	if err != nil {
		helper.PrintErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.PrintSuccessResponseMessage(ctx, "berhasil update user")
}
