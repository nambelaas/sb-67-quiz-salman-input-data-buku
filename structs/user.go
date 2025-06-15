package structs

type User struct {
	Id         int
	Username   string
	Password   string
	CreatedAt  string
	CreatedBy  string
	ModifiedAt string
	ModifiedBy string
}

type AuthRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}