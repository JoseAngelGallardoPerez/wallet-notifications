package newmoneyrequest

import (
	"fmt"

	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/service"
)

const (
	EventName = "NewMoneyRequest"
)

var (
	// PushTemplate is the template for push notifications
	PushTemplate = db.Template{
		Subject: "New money request",
		Content: "You have received a new money request from {{.OwnerFirstName}} {{.OwnerLastName}}",
	}

	defaultMethods = []service.Method{
		service.MethodPushNotification,
	}
)

// Event is event struct
type Event struct {
	notifier *service.Notifier
	logger   log15.Logger
}

// New creates new subscriber
func New(notifier *service.Notifier, logger log15.Logger) service.Subscriber {
	return &Event{notifier, logger}
}

// Call event callback
func (e *Event) Call(data *service.CallData) {
	methods := data.GetNotifiers(defaultMethods)
	for _, method := range methods {
		template := e.GetTemplateByMethod(method)

		_, err := template.ApplyTemplateData(data.TemplateData)
		if err != nil {
			e.logger.Error(fmt.Sprintf("[%s] Can't apply template data: %v", EventName, err))
		}

		receiver := service.NewReceiver(data.GetToByMethod(method), method)

		err = e.notifier.SendToReceiver(template, receiver, data.TemplateData)
		if err != nil {
			e.logger.Error(fmt.Sprintf("[%s] Can't send mail to user: %v", EventName, err))
		}
	}
}

func (e *Event) GetTemplateByMethod(method service.Method) *db.Template {
	tpl := PushTemplate
	return &tpl
}
