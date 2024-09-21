package worker

import (
	"context"
	"time"
	"github.com/emincanozcan/insider-assessment/internal/service"
)

type MessageSendJob struct {
	backgroundJobStarted bool
	running              bool
	svc                  *service.MessageService
	interval             time.Duration
	batchSize            int32
}

var singletonInstance *MessageSendJob = nil

func MakeOrGet(svc *service.MessageService, interval time.Duration, batchSize int32) *MessageSendJob {
	if singletonInstance == nil {
		// this is not 100% safe, but should be okay for our use case.
		singletonInstance = &MessageSendJob{
			running:   true,
			svc:       svc,
			interval:  interval,
			batchSize: batchSize,
		}
	}

	return singletonInstance
}

func (job *MessageSendJob) StartBackgroundJob() {
	if job.backgroundJobStarted {
		// already started ignore
		return
	}
	job.backgroundJobStarted = true

	ticker := time.NewTicker(job.interval)
	defer ticker.Stop()
	for range ticker.C {
		if !job.running {
			continue
		}

		job.svc.SendPendingMessages(context.Background(), job.batchSize)
	}
}

func (job *MessageSendJob) Start() {
	job.running = true
}

func (job *MessageSendJob) Stop() {
	job.running = false
}
