package structs

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserJwtData struct {
	UserId    int64
	Username  string
	Role      string
	LoginAt   time.Time
	ExpiredAt time.Time
}

type ClaimJwt struct {
	UserId   int
	Username string
	jwt.RegisteredClaims
}
