package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sb-67-go-quiz-salman-input-data-buku/structs"
)

func GetJwtData(ctx *gin.Context) (result structs.ClaimJwt, err error) {
	auth, exists := ctx.Get("auth")
	if !exists {
		return result, errors.New("data authorization tidak ditemukan")
	}

	dataJwt, ok := auth.(*structs.ClaimJwt)
	if !ok {
		return result, errors.New("data authorization tidak valid")
	}

	result = *dataJwt

	return result, nil
}
