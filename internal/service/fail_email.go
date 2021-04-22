package service

import (
	"context"

	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-notifications/internal/db"
)

const FailEmailSubject = "Important: Outgoing emails failure"

type FailEmail struct {
	notifier *Notifier
	logger   log15.Logger
}

func NewFailEmail(notifier *Notifier, logger log15.Logger) *FailEmail {
	return &FailEmail{notifier, logger}
}

func (s *FailEmail) notifyAdminsAboutFailEmail(templateData *db.TemplateData) {
	data := NewCallData("", []Method{MethodInternalMessage}, templateData)

	template := &db.Template{}
	template.Subject = FailEmailSubject
	template.Content = "There appears to be a problem with the system’s SMTP settings. As a consequence, outgoing emails are failing to send. Please revise the SMTP configuration at Settings > System Emails > Common Settings."
	template.ApplyTemplateData(data.TemplateData)

	client, _ := GetUserHandlerProtobufClient()
	response, err := client.GetByRoleName(context.Background(), &userpb.Request{RoleName: "admin"})
	if nil != err {
		s.logger.Error("can not receive admins", "error", err)
	}

	var to []string
	for _, u := range response.Users {
		to = append(to, u.UID)
	}

	receiver := NewReceiver(to, MethodInternalMessage)
	_ = s.notifier.SendToReceiver(template, receiver, data.TemplateData)
}
