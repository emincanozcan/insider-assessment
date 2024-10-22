package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		panic("Port is required!")
	}

	router := http.NewServeMux()
	router.HandleFunc("POST /api/receive-message", webhookHandler)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
  log.Println("Webhook Server: Listening on port: " + port)
	server.ListenAndServe()
}

type Request struct {
	To      string `json:"to"`
	Content string `json:"content"`
}
type Response struct {
	MessageId string `json:"messageId"`
	Message   string `json:"message"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("Received message to: %s, content: %s\n", req.To, req.Content)

	res := Response{
		MessageId: uuid.New().String(),
		Message:   "Accepted",
	}

	log.Printf("Responding with MessageId: %s\n", res.MessageId)

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
