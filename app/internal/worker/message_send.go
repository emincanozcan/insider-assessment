package worker

import (
	"context"
	"time"

	"github.com/emincanozcan/insider-assessment/internal/service"
)

type MessageSendJob struct {
	svc       *service.MessageService
	interval  time.Duration
	batchSize int32
	ticker    *time.Ticker
}

var singletonInstance *MessageSendJob = nil

func GetMessageSendJob() *MessageSendJob {
	return singletonInstance
}

func InitMessageSendJob(svc *service.MessageService, interval time.Duration, batchSize int32) *MessageSendJob {
	if singletonInstance == nil {
		singletonInstance = &MessageSendJob{
			svc:       svc,
			interval:  interval,
			batchSize: batchSize,
		}
	}

	return singletonInstance
}

func (job *MessageSendJob) StartBackgroundJob() {
	job.svc.SendPendingMessages(context.Background(), job.batchSize)
	job.ticker = time.NewTicker(job.interval)
	go job.run()
}

func (job *MessageSendJob) run() {
	for range job.ticker.C {
		job.svc.SendPendingMessages(context.Background(), job.batchSize)
	}
}

func (job *MessageSendJob) Start() {
	if job.ticker == nil {
		job.svc.SendPendingMessages(context.Background(), job.batchSize)
		job.ticker = time.NewTicker(job.interval)
		go job.run()
	}
}

func (job *MessageSendJob) Stop() {
	if job.ticker != nil {
		job.ticker.Stop()
		job.ticker = nil
	}
}
