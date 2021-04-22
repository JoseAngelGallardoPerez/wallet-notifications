package service

import (
	"context"
	"log"
	"net/http"

	"github.com/Confialink/wallet-pkg-env_mods"
	settingspb "github.com/Confialink/wallet-settings/rpc/proto/settings"
	userpb "github.com/Confialink/wallet-users/rpc/proto/users"

	"github.com/Confialink/wallet-notifications/internal/db"
	push_token "github.com/Confialink/wallet-notifications/internal/service/push-token"
	"github.com/Confialink/wallet-notifications/internal/srvdiscovery"
)

var PushTokenService *push_token.Service

// CallData is call data structure
type CallData struct {
	To           string
	Methods      []Method
	TemplateData *db.TemplateData
}

// Subscriber is subscriber interface
type Subscriber interface {
	Call(data *CallData)
}

// NewCallData is factory method for the call data
func NewCallData(to string, methods []Method, data *db.TemplateData) *CallData {
	settings, err := getSystemSettings()
	if nil != err {
		log.Printf("failed to retrieve system settings: %s", err.Error())
		return &CallData{}
	}
	data.ApplySystemSettings(settings)

	if to != "" {
		user, err := getUserByUID(to) // TODO: could be as uid or email or empty
		if nil != err {
			log.Fatalf(err.Error())
		}

		data.ApplyUser(user)
	}

	return &CallData{to, methods, data}
}

// GetToByMethod returns to by notifier
func (c *CallData) GetToByMethod(method Method) []string {
	switch method {
	case MethodEmail:
		return str2arr(c.TemplateData.Email)
	case MethodSMS:
		return str2arr(c.TemplateData.PhoneNumber)
	case MethodPushNotification:
		return c.GetPushTokens()
	default:
		return str2arr(c.To)
	}
}

// GetPushTokens returns push tokens
func (c *CallData) GetPushTokens() []string {
	// We use not all tokens, because we should not send notifications to the out of date devices.
	pushTokens, err := PushTokenService.NotExpiredTokensByUser(c.To)
	if err != nil {
		log.Println(err.Error())
		return make([]string, 0)
	}
	tokens := make([]string, 0, len(pushTokens))
	for _, token := range pushTokens {
		tokens = append(tokens, token.PushToken)
	}

	return tokens
}

// GetNotifiers returns notifiers
func (c *CallData) GetNotifiers(defaultMethods []Method) []Method {
	var notifiers []Method
	if len(c.Methods) > 0 {
		notifiers = c.Methods
	} else {
		notifiers = defaultMethods
	}

	if env_mods.IsDebugMode() {
		notifiers = append(notifiers, MethodSystemLog)
	}

	return notifiers
}

func str2arr(str string) []string {
	var to []string
	to = append(to, str)
	return to
}

// getUserByUID retrieve user by to
func getUserByUID(uid string) (*userpb.User, error) {
	client, err := GetUserHandlerProtobufClient()
	if nil != err {
		return nil, err
	}

	resp, err := client.GetByUID(context.Background(), &userpb.Request{UID: uid})
	if err != nil {
		return nil, err
	}

	return resp.GetUser(), nil
}

func getSystemSettings() ([]*settingspb.Setting, error) {
	client, err := GetSettingsHandlerProtobufClient()
	if nil != err {
		return nil, err
	}

	response, err := client.List(context.Background(), &settingspb.Request{Path: "regional/general/%"})
	if nil != err {
		return nil, err
	}

	return response.Settings, nil
}

func GetSettingsHandlerProtobufClient() (settingspb.SettingsHandler, error) {
	settingsUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNameSettings)
	if nil != err {
		return nil, err
	}
	return settingspb.NewSettingsHandlerProtobufClient(settingsUrl.String(), http.DefaultClient), nil
}

func GetUserHandlerProtobufClient() (userpb.UserHandler, error) {
	usersUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNameUsers)
	if err != nil {
		return nil, err
	}
	return userpb.NewUserHandlerProtobufClient(usersUrl.String(), http.DefaultClient), nil
}
