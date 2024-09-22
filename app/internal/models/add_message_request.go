package models

import (
	"errors"
	"strings"
)

type AddMessageRequest struct {
	Content   string `json:"content"`
	Recipient string `json:"recipient"`
}

func (r *AddMessageRequest) Trim() {
	r.Content = strings.TrimSpace(r.Content)
	r.Recipient = strings.TrimSpace(r.Recipient)
}

func (r *AddMessageRequest) Validate() error {
	if len(strings.TrimSpace(r.Content)) > 120 {
		return errors.New("Invalid char length for content. Max allowed length is 120.")
	}

	if len(strings.TrimSpace(r.Recipient)) > 80 {
		return errors.New("Invalid char length for recipient. Max allowed length is 80.")
	}
	return nil
}
