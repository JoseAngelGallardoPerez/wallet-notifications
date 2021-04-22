package service

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/service/providers"
	"github.com/Confialink/wallet-notifications/internal/service/providers/firebase"
	"github.com/Confialink/wallet-notifications/internal/service/providers/log"
	"github.com/Confialink/wallet-notifications/internal/service/providers/plivo"
	servicemessages "github.com/Confialink/wallet-notifications/internal/service/providers/service-messages"
	"github.com/Confialink/wallet-notifications/internal/service/providers/smtp"
	"github.com/Confialink/wallet-notifications/internal/service/providers/twilio"
)

const (
	SettingNameSmsProvider = "sms_provider"
)

var methodToProviderLoader = map[Method]providers.Loader{
	MethodEmail:            smtp.Load,
	MethodPushNotification: firebase.Load,
	MethodInternalMessage:  servicemessages.Load,
	MethodSystemLog:        log.Load,
}

var availableSmsProviders = map[string]providers.Loader{
	plivo.CODE:  plivo.Load,
	twilio.CODE: twilio.Load,
}

type Notifier struct {
	repo   db.RepositoryInterface
	logger log15.Logger
}

func NewNotifier(repo db.RepositoryInterface, logger log15.Logger) *Notifier {
	return &Notifier{repo, logger}
}

func (n *Notifier) SendToReceiver(tpl *db.Template, receiver *Receiver, data *db.TemplateData) error {
	if receiver.To == nil || len(receiver.To) == 0 {
		return nil
	}
	settings, _ := n.repo.GetAll()

	via, err := loadNotifier(receiver.NotificationMethod, settings, n.logger)
	if err != nil {
		return err
	}

	var res error

	switch v := via.(type) {
	case providers.ByEmail:
		var body string
		body, err = renderMail(tpl, settings, data)
		if err != nil {
			res = fmt.Errorf("failed to render mail. Reason: %s", err)
			break
		}
		res = v.To(receiver.To[0], receiver.To[1:]...).
			SetSubject(tpl.GetSubject()).
			SetBody(body).
			Send()
		if res != nil {
			failEmailService := NewFailEmail(n, n.logger.New("service", "FailEmail"))
			failEmailService.notifyAdminsAboutFailEmail(data)
		}

		break
	case providers.BySMS:
		res = v.To(receiver.To[0], receiver.To[1:]...).
			SetBody(tpl.GetContent()).
			Send()
		break
	case providers.ByChat:
		res = v.To(receiver.To[0], receiver.To[1:]...).
			SetBody(tpl.GetContent()).
			Send()
		break
	case providers.ByPush:
		res = v.To(receiver.To).
			SetSubject(tpl.GetSubject()).
			SetBody(tpl.GetContent()).
			SetEntityType(data.EntityType).
			SetEntityID(data.EntityID).
			SetMessageUnreadCount(data.MessageUnreadCount).
			Send()
		break
	case providers.ByProfile:
		res = v.To(receiver.To).
			From(data.SenderID).
			DeleteAfterRead(data.Tan != "").
			SetSubject(tpl.GetSubject()).
			SetDoNotDuplicateIfExists(tpl.GetSubject() == FailEmailSubject && data.SenderID == "").
			SetBody(tpl.GetContent()).
			Send()
		break
	default:
		res = fmt.Errorf(`invalid notifier "%s"`, receiver.NotificationMethod)
	}

	if nil != res {
		n.logger.Error("can not send notification to receiver", "error", res)
	}

	return res
}

type TemplateData struct {
	Body      string
	Logo      string
	Signature string
	SiteName  string
}

func loadNotifier(method Method, settings []*db.Settings, logger log15.Logger) (interface{}, error) {
	if method == MethodSMS {
		loader, err := LoadSmsProvider(settings)
		if err != nil {
			return nil, err
		}
		return loader(settings, logger)
	}

	if loader, exist := methodToProviderLoader[method]; exist {
		return loader(settings, logger)
	}
	return nil, fmt.Errorf("unknown notifier %s", method)
}

func renderMail(template *db.Template, settings []*db.Settings, templateData *db.TemplateData) (string, error) {
	logo, _ := db.GetSettingValue("logo_url", settings)
	mailSignature, _ := db.GetSettingValue("mail_signature", settings)

	if logo != "" {
		data := TemplateData{Logo: logo}
		var buf bytes.Buffer
		if err := logoTemplate.Execute(&buf, data); err != nil {
			return "", err
		}
		mailSignature = strings.Replace(mailSignature, "[Logo]", buf.String(), -1)
	}

	systemSettings, err := getSystemSettings()
	if nil != err {
		return "", err
	}

	templateData.ApplySystemSettings(systemSettings)

	signature, err := templateData.PrepareText(mailSignature)
	if err != nil {
		return "", err
	}

	siteName := templateData.SiteName

	body := strings.Replace(template.GetContent(), "\n", "<br/>", -1)
	data := TemplateData{
		Body:      body,
		Signature: signature,
		Logo:      logo,
		SiteName:  siteName,
	}

	var buf bytes.Buffer
	if err := mailTemplate.Execute(&buf, data); err != nil {
		return "", err
	}
	mail := buf.String()
	return mail, nil
}

func LoadSmsProvider(settings []*db.Settings) (providers.Loader, error) {
	smsProviderCode, err := db.GetSettingValue(SettingNameSmsProvider, settings)
	if err != nil {
		return nil, err
	}

	smsLoader, ok := availableSmsProviders[smsProviderCode]
	if !ok {
		return nil, fmt.Errorf("unknown sms provider: `%s`", smsProviderCode)
	}

	return smsLoader, nil
}

func GetNotificationNamesFromUserSettings() []string {
	return []string{
		"email_notification_dormant_profile_admin",
		"email_notification_unread_news_available",
		"email_notification_when_easytransac_transaction_fail",
		"email_notification_when_funds_added",
		"email_notification_when_internal_message",
		"email_notification_when_login_fails",
		"email_notification_when_new_file_uploaded",
		"email_notification_when_transfer_request_created",
		"email_notification_when_transfer_request_created",
		"internal_notification_when_back_to_pending",
		"internal_notification_when_cancel_pending",
		"internal_notification_when_cancel_processed",
		"internal_notification_when_executed",
		"internal_notification_when_cancel",
		"internal_notification_when_processed",
		"internal_notification_when_processed_was_executed",
		"internal_notification_when_received_transfer",
	}
}
