package api

import (
	"encoding/json"
	"net/http"

	"github.com/emincanozcan/insider-assessment/internal/service"
)

type Handler struct {
	messageService *service.MessageService
}

func NewHandler(messageService *service.MessageService) *Handler {
	return &Handler{
		messageService: messageService,
	}
}

func (h *Handler) GetSentMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := h.messageService.GetSentMessages(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(messages)
}
