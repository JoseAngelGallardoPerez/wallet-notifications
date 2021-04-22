package providers

import (
	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-notifications/internal/db"
)

type Loader func(settings []*db.Settings, logger log15.Logger) (interface{}, error)

type ByEmail interface {
	Code() string
	From(from string) ByEmail
	SetSubject(subject string) ByEmail
	SetBody(body string) ByEmail
	To(to string, cc ...string) ByEmail
	Send() error
}

type BySMS interface {
	Code() string
	From(from string) BySMS
	SetBody(body string) BySMS
	To(to string, cc ...string) BySMS
	Send() error
}

type ByChat interface {
	Code() string
	SetBody(body string) ByChat
	To(to string, cc ...string) ByChat
	Send() error
}

type ByPush interface {
	Code() string
	SetSubject(subject string) ByPush
	SetBody(body string) ByPush
	SetEntityType(entityType string) ByPush
	SetEntityID(entityId uint64) ByPush
	SetMessageUnreadCount(messageUnreadCount uint64) ByPush
	To(to []string) ByPush
	Send() error
}

type ByProfile interface {
	Code() string
	SetSubject(subject string) ByProfile
	SetBody(body string) ByProfile
	To(to []string) ByProfile
	From(from string) ByProfile
	DeleteAfterRead(deleteAfterRead bool) ByProfile
	SetDoNotDuplicateIfExists(doNotDuplicateIfExists bool) ByProfile
	Send() error
}
