package newtransferrequest

import (
	"context"
	"log"

	"github.com/Confialink/wallet-notifications/internal/service"

	"github.com/Confialink/wallet-notifications/internal/db"
	userspb "github.com/Confialink/wallet-users/rpc/proto/users"
)

const EventName = "NewTransferRequest"

var (
	TemplateTitle  = "New transfer request"
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
		log.Printf("[%s] can't find template: %v", EventName, err)
		return
	}

	if template.IsEnabled() {
		template, _ = template.ApplyTemplateData(data.TemplateData)

		recipients := e.getAdmins()

		for _, method := range data.GetNotifiers(defaultMethods) {
			to := e.getToByMethod(recipients, method)

			receiver := service.NewReceiver(to, method)

			err = e.notifier.SendToReceiver(template, receiver, data.TemplateData)
			if err != nil {
				log.Printf("[%s] Can't send mail to user: %v", EventName, err)
			}
		}

	}
}

func (e *Event) getToByMethod(users []*userspb.User, method service.Method) (to []string) {
	for _, u := range users {
		if method == service.MethodEmail {
			if e.isEmailNotificationActive(u.UID) {
				to = append(to, u.Email)
			}
		} else {
			to = append(to, u.UID)
		}
	}
	return
}

func (e *Event) getAdmins() []*userspb.User {
	client, _ := service.GetUserHandlerProtobufClient()
	res, err := client.GetByRoleName(context.Background(), &userspb.Request{RoleName: "admin"})
	if err != nil {
		log.Printf("can't get list of admins: %v", err)
		return nil
	}
	return res.Users
}

func (e *Event) isEmailNotificationActive(uid string) bool {
	s, err := e.repo.FindUserOptionByUIDAndName(uid, "email_notification_when_transfer_request_created")
	if err != nil {
		log.Printf("can't find user settings by uid: %s", uid)
		return false
	}
	return "0" != s.IsActive
}
