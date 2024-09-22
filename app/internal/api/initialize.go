package api

import (
	"log"
	"net/http"

	_ "github.com/emincanozcan/insider-assessment/docs" // Generated by Swag CLI
	"github.com/emincanozcan/insider-assessment/internal/service"
	"github.com/emincanozcan/insider-assessment/internal/worker"
	"github.com/swaggo/http-swagger"
)

func InitializeApi(messageSendJob *worker.MessageSendJob, messageService *service.MessageService, port string) {
	handler := NewHandler(messageService)
	router := http.NewServeMux()

	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	router.HandleFunc("GET /messages/sent", handler.GetSentMessages)
	router.HandleFunc("POST /messages", handler.CreateMessage)
	router.HandleFunc("POST /messages/processing/start", handler.StartProcessing)
	router.HandleFunc("POST /messages/processing/stop", handler.StopProcessing)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Println("App server listening on port: " + port)
	server.ListenAndServe()
}
