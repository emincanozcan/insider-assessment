package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/emincanozcan/insider-assessment/internal/database/sqlc"
	"github.com/emincanozcan/insider-assessment/internal/models"
	"github.com/emincanozcan/insider-assessment/pkg/webhook_client"
	"github.com/redis/go-redis/v9"
)

type MessageService struct {
	sqlcQueries   *sqlc.Queries
	webhookClient *webhook_client.Client
	redisClient   *redis.Client
}

func NewMessageService(queries *sqlc.Queries, redisClient *redis.Client, webhookClient *webhook_client.Client) *MessageService {
	return &MessageService{
		sqlcQueries:   queries,
		webhookClient: webhookClient,
		redisClient:   redisClient,
	}
}

func (s *MessageService) SendPendingMessages(ctx context.Context, maxSize int32) {
	log.Println("Message sending to the webhook service started at: " + time.Now().String())
	messages, err := s.sqlcQueries.GetPendingMessagesAndMarkAsSending(ctx, maxSize)
	if err != nil {
		fmt.Printf("failed to get pending messages: %s", err.Error())
		return
	}

	var wg sync.WaitGroup

	for _, msg := range messages {
		wg.Add(1)
		go func(msg *sqlc.Message) {
			defer wg.Done()
			s.processMessage(ctx, msg)
		}(&msg)
	}

	wg.Wait()
}

func (s *MessageService) processMessage(ctx context.Context, msg *sqlc.Message) error {
	log.Printf("Processing message id: %d, to: %s, content: %s\n", msg.ID, msg.Recipient, msg.Content)
	res, err := s.webhookClient.Send(msg.Recipient, msg.Content)

	if err != nil {
		log.Printf("Main server: Error in http send request, rollback result to resend it in the future. " + err.Error())
		s.sqlcQueries.MarkMessageAsNotSent(ctx, msg.ID)
		return fmt.Errorf("failed to send message: %w", err)
	}

	err = s.sqlcQueries.MarkMessageAsSent(ctx, msg.ID)
	if err != nil {
		return fmt.Errorf("failed to update message status to sent: %w", err)
	}
	s.redisClient.Set(ctx, "message:sent:"+res.MessageID, time.Now().UTC().UnixMilli(), time.Duration(30*24)*time.Hour)
	if err != nil {
		log.Println("failed to store in redis: " + err.Error())
	}
	return nil
}

func (s *MessageService) GetSentMessages(ctx context.Context) ([]models.SentMessageResponseModel, error) {
	lastMessages, err := s.sqlcQueries.ListSentMessages(ctx, 1000)
	if err != nil {
		return nil, err
	}

	list := []models.SentMessageResponseModel{}

	for _, message := range lastMessages {
		list = append(list, models.SentMessageResponseModel{
			InternalId: int(message.ID),
			Recipient:  message.Recipient,
			Content:    message.Content,
		})
	}

	return list, nil
}

func (s *MessageService) Create(ctx context.Context, req *models.AddMessageRequest) (*models.AddMessageResponse, error) {
	req.Trim()
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	msg, err := s.sqlcQueries.CreateMessage(ctx, sqlc.CreateMessageParams{
		Content:   req.Content,
		Recipient: req.Recipient,
	})

	return &models.AddMessageResponse{
		ID:        int(msg.ID),
		Recipient: msg.Recipient,
		Content:   msg.Content,
	}, nil
}
func (s *MessageService) AddTestMessages(ctx context.Context) {
	now := time.Now().UTC().String()
	for i := 0; i < 10; i++ {
		s.sqlcQueries.CreateMessage(ctx, sqlc.CreateMessageParams{
			Content:   "Some random message, created at: " + now + "idx:" + strconv.Itoa(i),
			Recipient: "emincan@emincanozcan.com",
		})
	}
}
