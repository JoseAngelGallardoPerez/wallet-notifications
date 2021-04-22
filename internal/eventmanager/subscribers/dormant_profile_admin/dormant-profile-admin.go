package dormant_profile_admin

import (
	"context"
	"log"

	"github.com/Confialink/wallet-notifications/internal/service"

	"github.com/Confialink/wallet-notifications/internal/db"
	userspb "github.com/Confialink/wallet-users/rpc/proto/users"
)

const EventName = "DormantProfileAdmin"

var (
	TemplateTitle  = "Dormant profile"
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

	template, err := e.repo.FindOneByTitleAndScope(TemplateTitle, db.ScopeAdmin)
	if err != nil {
		log.Printf("[%s] Can't find template: %v", EventName, err)
		return
	}

	template.ApplyTemplateData(data.TemplateData)

	client, _ := service.GetUserHandlerProtobufClient()
	response, err := client.GetByRoleName(context.Background(), &userspb.Request{RoleName: "admin"})
	if nil != err {
		log.Printf("Can't send notification: %v", err)
		return
	}
	var to []string
	for _, u := range response.Users {
		// we exclude initiator
		if u.UID != data.To {
			setting, err := e.repo.FindUserOptionByUIDAndName(u.UID, "email_notification_dormant_profile_admin")
			if err != nil {
				log.Printf("Can't find settings from UID: %s", u.UID)
				continue
			}

			if setting.IsActive == "0" {
				continue
			}
			to = append(to, u.Email)
		}
	}

	methods := data.GetNotifiers(defaultMethods)
	for _, method := range methods {
		tpl := template

		if nil != err {
			log.Printf("Can't send notification: %v", err)
			return
		}

		if method == service.MethodEmail {
			if !tpl.IsEnabled() {
				continue
			}
		}

		receiver := service.NewReceiver(to, method)
		err = e.notifier.SendToReceiver(template, receiver, data.TemplateData)
		if err != nil {
			log.Printf("[%s] Can't send mail to user: %v", EventName, err)
		}
	}
}
