package tancreate

import (
	"github.com/Confialink/wallet-notifications/internal/srvdiscovery"
	"context"
	"log"

	accountspb "github.com/Confialink/wallet-accounts/rpc/accounts"
	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/service"
	"github.com/Confialink/wallet-pkg-http"
)

const EventName = "TanCreate"

var (
	DefaultTemplate = &db.Template{
		Title:   "Tan created",
		Scope:   db.ScopeUser,
		Legend:  "Tan creation for user",
		Subject: "TANs",
		Content: "Please, copy or print this message, since it is only going to be shown once. \nYour TANs: \n{{.Tan}}",
	}
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
	methods := data.GetNotifiers(defaultMethods)
	for _, method := range methods {
		template := e.getTemplateByMethod(method)
		_, err := template.ApplyTemplateData(data.TemplateData)
		if err != nil {
			log.Printf("[%s] Can't apply %s template data: %v", EventName, method, err)
		}

		receiver := service.NewReceiver(data.GetToByMethod(method), method)

		err = e.notifier.SendToReceiver(template, receiver, data.TemplateData)
		if err != nil {
			log.Printf("[%s] Can't send %s to user: %v", EventName, method, err)
		}
	}
}

func (e *Event) getTemplateByMethod(method service.Method) *db.Template {
	var content string

	switch method {
	case service.MethodEmail, service.MethodInternalMessage:
		content, _ = e.getTemplateContentFromAccountsSettings("tan_message_content")
	case service.MethodSMS:
		content, _ = e.getTemplateContentFromNotificationsSettings("tans_text_for_sms")
	}

	template := DefaultTemplate
	template.Content = content

	return template
}

func (e *Event) getTemplateContentFromNotificationsSettings(name string) (string, error) {
	settings, err := e.repo.GetAll()
	if err != nil {
		log.Printf("[%s] Can't load settings: %v", EventName, err)
	}
	content, err := db.GetSettingValue(name, settings)

	return content, err
}

func (e *Event) getTemplateContentFromAccountsSettings(name string) (string, error) {
	client, err := e.getAccountsClient()
	if err != nil {
		log.Printf("[%s] Can't get accounts client: %v", EventName, err)
		return "", err
	}

	req := accountspb.SettingsByNameReq{
		Name: name,
	}

	resp, err := client.GetSettingsByName(context.Background(), &req)
	if err != nil {
		log.Printf("[%s] Can't load accounts settings: %v", EventName, err)
		return "", err
	}

	return resp.Value, err
}

func (e *Event) getAccountsClient() (accountspb.AccountsProcessor, error) {
	accountsUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNameAccounts)
	if nil != err {
		return nil, err
	}
	return accountspb.NewAccountsProcessorProtobufClient(accountsUrl.String(), http.DefaultClient), nil
}
