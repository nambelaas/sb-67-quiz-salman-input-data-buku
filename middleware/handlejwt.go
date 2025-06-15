package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sb-67-go-quiz-salman-input-data-buku/helper"
	"github.com/sb-67-go-quiz-salman-input-data-buku/structs"
	"github.com/spf13/viper"
)

func CheckJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := GetJwtTokenFromHeader(c)
		if err != nil {
			helper.PrintErrorResponse(c, http.StatusBadRequest, err)
		}

		token, err := jwt.ParseWithClaims(tokenString, &structs.ClaimJwt{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("JWT.SignatureKey")), nil
		})

		if err != nil {
			helper.PrintErrorResponse(c, http.StatusBadRequest, errors.New("token tidak valid"))
		}

		claims, ok := token.Claims.(*structs.ClaimJwt)
		if !ok || !token.Valid {
			helper.PrintErrorResponse(c, http.StatusBadRequest, errors.New("token gagal diklaim"))
		}

		if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
			helper.PrintErrorResponse(c, http.StatusBadRequest, errors.New("token kadaluarsa, silahkan login kembali"))
		}

		c.Set("auth", claims)
		c.Next()
	}
}

func GetJwtTokenFromHeader(c *gin.Context) (tokenString string, err error) {
	authHeader := c.Request.Header.Get("Authorization")

	if checkIsStringEmpty(authHeader) {
		return tokenString, errors.New("header authorization dibutuhkan")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return tokenString, errors.New("format header authorization tidak sesuai")
	}

	return parts[1], nil
}

func checkIsStringEmpty(input string) bool {
	return input == ""
}

func GenerateJwtToken(id int, username string) string {
	claims := structs.ClaimJwt{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    viper.GetString("JWT.Issuer"),
		},
		UserId: id,
		Username: username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	res, _ := token.SignedString([]byte(viper.GetString("JWT.SignatureKey")))

	return res
}