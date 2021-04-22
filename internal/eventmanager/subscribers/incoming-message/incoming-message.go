package incomingmessage

import (
	"log"

	"github.com/Confialink/wallet-notifications/internal/service"

	"github.com/Confialink/wallet-notifications/internal/db"
)

// EventName is the event name
const EventName = "IncomingMessage"

var (
	// TemplateTitle is the template title
	TemplateTitle = "Incoming Messages"

	// PushTemplate is the template for push notifications
	PushTemplate = &db.Template{
		Subject: "New message received",
		Content: "You have received a new message from {{.PrivateMessageAuthor}}",
	}
	defaultMethods = []service.Method{
		service.MethodEmail,
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
		return
	}

	if template.IsEnabled() {
		template.ApplyTemplateData(data.TemplateData)

		methods := data.GetNotifiers(defaultMethods)
		for _, method := range methods {
			tpl := template
			receiver := service.NewReceiver(data.GetToByMethod(method), method)

			if method == service.MethodPushNotification {
				tpl = PushTemplate
				tpl.ApplyTemplateData(data.TemplateData)
			}

			if method == service.MethodEmail {
				if !tpl.IsEnabled() {
					continue
				}

				setting, err := e.repo.FindUserOptionByUIDAndName(data.To, "email_notification_when_internal_message")
				if err != nil {
					log.Printf("Can't find settings from UID: %s", data.To)
					return
				}

				if setting.IsActive == "0" {
					continue
				}
			}

			err := e.notifier.SendToReceiver(tpl, receiver, data.TemplateData)
			if err != nil {
				log.Printf("[%s] Can't send %s to user: %v", EventName, method, err)
			}
		}
	}
}
