package api

import (
	"net/http"

	"github.com/emincanozcan/insider-assessment/internal/service"
	"github.com/emincanozcan/insider-assessment/internal/worker"
)

func InitializeApi(messageSendJob *worker.MessageSendJob, messageService *service.MessageService, port string) {
	handler := NewHandler(messageService)
	router := http.NewServeMux()
	router.HandleFunc("GET /api/messages/sent", handler.GetSentMessages)
	router.HandleFunc("GET /api/message-sending/start", handler.StartProcessing(messageSendJob))
	router.HandleFunc("GET /api/message-sending/stop", handler.StopProcessing(messageSendJob))

	// NOTE: this route added for testing purposes, it adds 10 new dummy messages as unsent to the databaase.
	router.HandleFunc("GET /api/add-test-messages", handler.AddTestMessages)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	server.ListenAndServe()
}
