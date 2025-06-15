package structs

type Category struct {
	Id         int
	Name       string `json:"name" binding:"required"`
	CreatedAt  string
	CreatedBy  string
	ModifiedAt string
	ModifiedBy string
}
