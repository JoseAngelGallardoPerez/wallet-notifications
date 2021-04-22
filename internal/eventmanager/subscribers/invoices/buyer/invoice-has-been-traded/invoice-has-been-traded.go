package invoice_has_been_traded

import (
	"fmt"
	"log"

	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/service"

	"github.com/inconshreveable/log15"
)

const (
	EventName     = "InvoiceHasBeenTraded"
	TemplateTitle = "Invoice Has Been Traded"
)

var (
	defaultMethods = []service.Method{
		service.MethodEmail,
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
	return &Event{
		notifier: notifier,
		logger:   logger,
		repo:     repo,
	}
}

// Call event callback
func (e *Event) Call(data *service.CallData) {
	// Send mail buyer and supplier
	template, err := e.repo.FindOneByTitleAndScope(TemplateTitle, db.ScopeUser)
	if err != nil {
		log.Printf("[%s] Can't find template: %v", EventName, err)
	}

	if template.IsEnabled() {
		_, err := template.ApplyTemplateData(data.TemplateData)
		if err != nil {
			e.logger.Error(fmt.Sprintf("[%s] Can't apply template data: %v", EventName, err))
		}

		methods := data.GetNotifiers(defaultMethods)
		for _, method := range methods {
			receiver := service.NewReceiver(data.GetToByMethod(method), method)

			setting, err := e.repo.FindUserOptionByUIDAndName(data.To, "internal_notification_when_executed")
			if err != nil {
				log.Printf("Can't find settings from UID: %s", data.To)
				return
			}

			if setting.IsActive == "0" {
				continue
			}

			err = e.notifier.SendToReceiver(template, receiver, data.TemplateData)
			if err != nil {
				e.logger.Error(fmt.Sprintf("[%s] Can't send %s to user: %v", EventName, method, err))
			}
		}
	}
}
