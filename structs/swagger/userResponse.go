package swagger

type ResponseRegister struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

type ResponseLogin struct {
	Success bool   `json:"Success"`
	Token   string `json:"Token"`
}

type ResponseUpdate struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}