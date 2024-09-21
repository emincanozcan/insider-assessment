package models

type MessageStatus int

const (
	MessageStatusPending MessageStatus = iota
	MessageStatusSending
	MessageStatusSent
)
