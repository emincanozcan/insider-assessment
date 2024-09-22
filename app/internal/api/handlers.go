package api

import (
	"encoding/json"
	"net/http"

	"github.com/emincanozcan/insider-assessment/internal/models"
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

// @Summary Get sent messages
// @Description Retrieve all sent messages
// @Tags messages
// @Produce json
// @Success 200 {array} models.SentMessageResponseModel
// @Failure 500 {string} string "Internal Server Error"
// @Router /messages/sent [get]
func (h *Handler) GetSentMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := h.messageService.GetSentMessages(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

// @Summary Create a new message
// @Description Create a new message with content and recipient.
// @Tags messages
// @Produce json
// @Param message body models.AddMessageRequest true "Message request payload"
// @Success 200 {object} models.AddMessageResponse "Success response"
// @Failure 400 {object} map[string]string "Bad request response"
// @Router /messages [post]
func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var messageReq models.AddMessageRequest

	err := json.NewDecoder(r.Body).Decode(&messageReq)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid payload",
			"error":   "The payload is not a valid json",
		})
		return
	}

	// NOTE: Service only return validation errors for this case.
	msg, err := h.messageService.Create(r.Context(), &messageReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid payload",
			"error":   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(msg)
}

// @Summary Start message processing
// @Description Start the message sending job
// @Tags processing
// @Produce json
// @Success 200 {object} map[string]string
// @Router /messages/processing/start [post]
func (h *Handler) StartProcessing(job *worker.MessageSendJob) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		job.Start()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Processing started"})
	}
}

// @Summary Stop message processing
// @Description Stop the message sending job
// @Tags processing
// @Produce json
// @Success 200 {object} map[string]string
// @Router /messages/processing/stop [post]
func (h *Handler) StopProcessing(job *worker.MessageSendJob) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		job.Stop()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Processing stopped"})
	}
}
