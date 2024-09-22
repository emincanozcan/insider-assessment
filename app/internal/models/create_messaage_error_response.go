package models

type CreateMessageErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
