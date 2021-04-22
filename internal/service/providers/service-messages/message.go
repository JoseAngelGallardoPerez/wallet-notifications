package servicemessages

import (
	"context"
	"errors"
	"net/http"

	"github.com/Confialink/wallet-messages/rpc/messages"
	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/service/providers"
	"github.com/Confialink/wallet-notifications/internal/srvdiscovery"
)

const CODE = "message"

type Options struct {
	to                     []string
	from                   string
	deleteAfterRead        bool
	doNotDuplicateIfExists bool
}

type client struct {
	options Options
	body    string
	subject string
	logger  log15.Logger
}

var _ providers.ByProfile = &client{}

func New(options Options, logger log15.Logger) *client {
	return &client{options: options, logger: logger}
}

func Load(settings []*db.Settings, logger log15.Logger) (interface{}, error) {
	return New(Options{}, logger), nil
}

func (c client) Code() string {
	return CODE
}

func (c client) SetSubject(subject string) providers.ByProfile {
	c.subject = subject
	return &c
}

func (c client) SetBody(body string) providers.ByProfile {
	c.body = body
	return &c
}

func (c client) To(to []string) providers.ByProfile {
	c.options.to = to
	return &c
}

func (c client) From(from string) providers.ByProfile {
	c.options.from = from
	return &c
}

func (c client) DeleteAfterRead(deleteAfterRead bool) providers.ByProfile {
	c.options.deleteAfterRead = deleteAfterRead
	return &c
}

func (c client) SetDoNotDuplicateIfExists(doNotDuplicateIfExists bool) providers.ByProfile {
	c.options.doNotDuplicateIfExists = doNotDuplicateIfExists
	return &c
}

func (c *client) Send() error {
	if len(c.options.to) == 0 {
		return errors.New("Missing to")
	}

	client, err := getMessageSenderProtobufClient()
	if nil != err {
		return err
	}

	for _, recipientId := range c.options.to {
		req := &messages.SendMessageReq{
			Subject:         c.subject,
			Message:         c.body,
			SenderId:        c.options.from,
			RecipientId:     recipientId,
			DeleteAfterRead: c.options.deleteAfterRead,
			DoNotDuplicateIfExists: c.options.doNotDuplicateIfExists,
		}

		_, err = client.SendMessage(context.Background(), req)
		if nil != err {
			return err
		}
	}

	return nil
}

func getMessageSenderProtobufClient() (messages.MessageSender, error) {
	messagesUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNameMessages)
	if nil != err {
		return nil, err
	}
	return messages.NewMessageSenderProtobufClient(messagesUrl.String(), http.DefaultClient), nil
}
