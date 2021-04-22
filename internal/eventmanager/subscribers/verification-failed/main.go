package verificationfailed

import (
	"context"
	"log"

	"github.com/Confialink/wallet-notifications/internal/service"

	"github.com/Confialink/wallet-notifications/internal/db"
	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
)

const EventName = "VerificationFailed"

var (
	TemplateTitle = "User ID Validation Failed"

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
	// Send mail to admins
	template, err := e.repo.FindOneByTitleAndScope(TemplateTitle, db.ScopeAdmin)
	if err != nil {
		log.Printf("[%s] Can't find template: %v", EventName, err)
	}

	if template.IsEnabled() {
		template.ApplyTemplateData(data.TemplateData)

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

		err = e.notifier.SendToReceiver(template, receiver, data.TemplateData)
		if err != nil {
			log.Printf("[%s] Can't send mail to user: %v", EventName, err)
		}
	}
}
