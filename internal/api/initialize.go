package api

import (
	"log"
	"net/http"

	"github.com/emincanozcan/insider-assessment/internal/service"
)

func InitializeApi(messageService *service.MessageService, port string) {
	handler := NewHandler(messageService)
	router := http.NewServeMux()
	router.HandleFunc("GET /api/messages/sent", handler.GetSentMessages)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Println("App Server listening:" + port)
	server.ListenAndServe()
}
