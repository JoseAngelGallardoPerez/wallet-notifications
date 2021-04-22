package incomingtransaction

import (
	"log"

	"github.com/Confialink/wallet-notifications/internal/service"

	"github.com/Confialink/wallet-notifications/internal/db"
)

const EventName = "IncomingTransaction"

var (
	TemplateTitle  = "Incoming transaction"
	defaultMethods = []service.Method{
		service.MethodEmail,
		service.MethodInternalMessage,
		service.MethodPushNotification,
	}
)

// Event is event struct
type Event struct {
	notifier *service.Notifier
	repo     db.RepositoryInterface
}

// New creates new subscriber
func New(notifier *service.Notifier, repo db.RepositoryInterface) service.Subscriber {
	return &Event{notifier, repo}
}

// Call event callback
func (e *Event) Call(data *service.CallData) {
	template, err := e.repo.FindOneByTitleAndScope(TemplateTitle, db.ScopeUser)
	if err != nil {
		log.Printf("[%s] Can't find template: %v", EventName, err)
		return
	}

	_, err = template.ApplyTemplateData(data.TemplateData)
	if err != nil {
		log.Printf("[%s] Can't apply template data: %v", EventName, err)
	}

	methods := data.GetNotifiers(defaultMethods)
	for _, method := range methods {
		tpl := template
		receiver := service.NewReceiver(data.GetToByMethod(method), method)

		if method == service.MethodEmail {
			if !tpl.IsEnabled() {
				continue
			}

			setting, err := e.repo.FindUserOptionByUIDAndName(data.To, "email_notification_when_funds_added")
			if err != nil {
				log.Printf("Can't find settings from UID: %s", data.To)
				return
			}

			if setting.IsActive == "0" {
				continue
			}

		} else if method == service.MethodInternalMessage {
			template = &db.Template{
				Subject: "You have received a transfer from another user.",
				Content: "Another user has made a transfer to your account #[AccountNumber].",
			}
			_, err = template.ApplyTemplateData(data.TemplateData)
			if err != nil {
				log.Printf("[%s] Can't apply template data: %v", EventName, err)
			}

			setting, err := e.repo.FindUserOptionByUIDAndName(data.To, "internal_notification_when_received_transfer")
			if err != nil {
				log.Printf("Can't find settings from UID: %s", data.To)
				return
			}

			if setting.IsActive == "0" {
				continue
			}
		}

		err := e.notifier.SendToReceiver(template, receiver, data.TemplateData)
		if err != nil {
			log.Printf("[%s] Can't send %s to user: %v", EventName, method, err)
		}
	}
}
