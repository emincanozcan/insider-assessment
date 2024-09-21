package api

import (
	"encoding/json"
	"net/http"

	"github.com/emincanozcan/insider-assessment/internal/service"
	"github.com/emincanozcan/insider-assessment/internal/worker"
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

func (h *Handler) AddTestMessages(w http.ResponseWriter, r *http.Request) {
	h.messageService.AddTestMessages(r.Context())
	w.Write([]byte("10 new messages added."))
}

func (h *Handler) StartProcessing(job *worker.MessageSendJob) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		job.Start()
		w.Write([]byte("Processing started."))
	}
}

func (h *Handler) StopProcessing(job *worker.MessageSendJob) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		job.Stop()
		w.Write([]byte("Processing stopped."))
	}
}
