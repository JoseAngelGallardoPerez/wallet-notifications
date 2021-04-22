package failedloginattempts

import (
	"context"
	"log"

	"github.com/Confialink/wallet-notifications/internal/service"

	"github.com/Confialink/wallet-notifications/internal/db"
	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
)

const EventName = "FailedLoginAttempts"

var (
	TemplateTitle      = "Blocked-failed login attempts"
	AdminTemplateTitle = "Blocked-failed login attempts (Admin)"
	defaultMethods     = []service.Method{
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
	// Send mail to user
	template, err := e.repo.FindOneByTitleAndScope(TemplateTitle, db.ScopeUser)
	if err != nil {
		log.Printf("[%s] Can't find template: %v", EventName, err)
		return
	}

	if template.IsEnabled() {
		template.ApplyTemplateData(data.TemplateData)

		methods := data.GetNotifiers(defaultMethods)
		for _, method := range methods {
			tpl := template
			receiver := service.NewReceiver(data.GetToByMethod(method), method)

			if method == service.MethodEmail {
				if !tpl.IsEnabled() {
					continue
				}

				setting, err := e.repo.FindUserOptionByUIDAndName(data.To, "email_notification_when_login_fails")
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

	// Send mail to admins
	template2, err := e.repo.FindOneByTitleAndScope(AdminTemplateTitle, db.ScopeAdmin)
	if err != nil {
		log.Printf("[%s] Can't find template: %v", EventName, err)
	}

	if template2.IsEnabled() {
		template2.ApplyTemplateData(data.TemplateData)

		client, _ := service.GetUserHandlerProtobufClient()
		response, err := client.GetByRoleName(context.Background(), &userpb.Request{RoleName: "admin"})
		if nil != err {
			log.Printf("Can't send notification: %v", err)
			return
		}

		var to []string
		for _, u := range response.Users {
			to = append(to, u.Email)
		}

		method := service.MethodEmail
		receiver := service.NewReceiver(to, method)

		err = e.notifier.SendToReceiver(template2, receiver, data.TemplateData)
		if err != nil {
			log.Printf("[%s] Can't send mail to user: %v", EventName, err)
		}
	}
}
