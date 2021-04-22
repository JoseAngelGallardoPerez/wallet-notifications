package service

import "errors"

type Method string

const (
	MethodSMS              = Method("sms")
	MethodEmail            = Method("email")
	MethodInternalMessage  = Method("internal_message")
	MethodPushNotification = Method("push_notification")
	MethodSystemLog        = Method("system_log")
)

var knownMethods = map[string]Method{
	"sms":               MethodSMS,
	"email":             MethodEmail,
	"internal_message":  MethodInternalMessage,
	"push_notification": MethodPushNotification,
	"system_log":        MethodSystemLog,
}

func MethodFromString(method string) (Method, error) {
	if knownMethod, exist := knownMethods[method]; exist {
		return knownMethod, nil
	}
	return Method(""), errors.New("unknown notifier method " + method)
}
