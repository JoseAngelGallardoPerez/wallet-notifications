package changepassword

import (
	"log"

	"github.com/Confialink/wallet-notifications/internal/service"

	"github.com/Confialink/wallet-notifications/internal/db"
)

const EventName = "ChangePassword"

var (
	templateTitle  = "Username & password change"
	defaultMethods = []service.Method{
		service.MethodEmail,
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

	template, err := e.repo.FindOneByTitleAndScope(templateTitle, db.ScopeUser)
	if err != nil {
		log.Printf("[%s] Can't find template: %v", EventName, err)
		return
	}

	template.ApplyTemplateData(data.TemplateData)

	methods := data.GetNotifiers(defaultMethods)
	for _, method := range methods {

		receiver := service.NewReceiver(data.GetToByMethod(method), method)

		err = e.notifier.SendToReceiver(template, receiver, data.TemplateData)
		if err != nil {
			log.Printf("[%s] Can't send %s to user: %v", EventName, method, err)
		}
	}
}
