package structs

import (
	"time"
)

type Category struct {
	Id         int
	Name       string `json:"name" binding:"required"`
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt *time.Time
	ModifiedBy *string
}
