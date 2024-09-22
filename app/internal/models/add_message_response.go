package models

type AddMessageResponse struct {
	ID        int    `json:"id"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}
