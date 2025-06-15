package structs

import (
	"time"
)

type User struct {
	Id         int
	Username   string
	Password   string
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt *time.Time
	ModifiedBy *string
}

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
