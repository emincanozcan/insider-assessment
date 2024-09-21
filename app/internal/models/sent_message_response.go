package models

type SentMessageResponseModel struct {
	InternalId int    `json:"id"`
	Recipient  string `json:"recipient"`
	Content    string `json:"content"`
}
