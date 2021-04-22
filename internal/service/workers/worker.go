package workers

import (
	"fmt"

	"github.com/inconshreveable/log15"
	"github.com/jasonlvhit/gocron"

	push_token "github.com/Confialink/wallet-notifications/internal/service/push-token"
)

const (
	deleteOldPushTokensAt = "23:30"
)

type Job func() error

type Worker struct {
	scheduler        *gocron.Scheduler
	logger           log15.Logger
	pushTokenService *push_token.Service
}

func NewWorker(
	scheduler *gocron.Scheduler,
	logger log15.Logger,
	pushTokenService *push_token.Service,
) *Worker {
	return &Worker{
		scheduler,
		logger.New("service", "Jobs"),
		pushTokenService,
	}
}

// Start initializes jobs and starts scheduler
func (s *Worker) Start(scheduler *gocron.Scheduler) error {
	s.logger.Info("starting scheduler...")

	err := scheduler.Every(1).Day().At(deleteOldPushTokensAt).Do(
		s.prepare("DeleteOldPushTokens", s.pushTokenService.DeleteExpiredTokens),
	)
	if err != nil {
		return err
	}

	scheduler.Start()
	s.logger.Info("scheduler is started")
	return nil
}

// prepare wraps the command into log commands
func (s *Worker) prepare(jobName string, f Job) func() {
	return func() {
		s.logger.Info(fmt.Sprintf("job '%s' is started", jobName))
		err := f()
		if err == nil {
			s.logger.Info(fmt.Sprintf("job '%s' is finished", jobName))
			return
		}

		s.logger.Error("cannot execute job", "job", jobName, "err", err)
	}
}
