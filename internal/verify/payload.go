package verify

type SendResponse struct {
	Success bool    `json:"success"`
	Data    any     `json:"data"`
	Errors  []error `json:"errors"`
}

type SendRequest struct {
	Email string `json:"email" validate:"required,email"`
}
