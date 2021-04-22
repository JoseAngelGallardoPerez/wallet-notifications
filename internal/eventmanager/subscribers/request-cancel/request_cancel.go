package requestcancel

import (
	"fmt"
	"log"

	"github.com/Confialink/wallet-notifications/internal/service"

	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-notifications/internal/db"
)

const EventName = "RequestCancel"

var (
	defaultMethods = []service.Method{
		service.MethodInternalMessage,
	}
)

// Event is event struct
type Event struct {
	notifier *service.Notifier
	logger   log15.Logger
	repo     db.RepositoryInterface
}

// New creates new subscriber
func New(notifier *service.Notifier, logger log15.Logger, repo db.RepositoryInterface) service.Subscriber {
	return &Event{notifier, logger, repo}
}

// Call event callback
func (e *Event) Call(data *service.CallData) {
	template := &db.Template{
		Subject: "Your transfer request has been cancelled.",
		Content: "Your pending request with ID [RequestID] has been cancelled by an Administrator.",
	}

	_, err := template.ApplyTemplateData(data.TemplateData)
	if err != nil {
		e.logger.Error(fmt.Sprintf("[%s] Can't apply template data: %v", EventName, err))
	}

	methods := data.GetNotifiers(defaultMethods)
	for _, method := range methods {
		receiver := service.NewReceiver(data.GetToByMethod(method), method)

		if method == service.MethodInternalMessage {

			setting, err := e.repo.FindUserOptionByUIDAndName(data.To, "internal_notification_when_cancel")
			if err != nil {
				log.Printf("Can't find settings from UID: %s", data.To)
				return
			}

			if setting.IsActive == "0" {
				continue
			}
		}

		err = e.notifier.SendToReceiver(template, receiver, data.TemplateData)
		if err != nil {
			e.logger.Error(fmt.Sprintf("[%s] Can't send %s to user: %v", EventName, method, err))
		}
	}
}
