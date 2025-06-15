package helper

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PrintErrorResponse(ctx *gin.Context, code int, err error) {
	var errData map[string]interface{}
	_ = json.Unmarshal([]byte(err.Error()), &errData)

	response := gin.H{
		"Success": false,
	}

	for k, v := range errData {
		response[k] = v
	}

	ctx.AbortWithStatusJSON(code, response)
}

func PrintSuccessResponseToken(ctx *gin.Context, token string) {
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Token":   token,
	})
}

func PrintSuccessResponseMessage(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": message,
	})
}

func PrintSuccessResponseMessageWithData(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": message,
		"Data":    data,
	})
}
