package log

import (
	"errors"

	"github.com/Confialink/wallet-notifications/internal/service/providers"

	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/inconshreveable/log15"
)

const CODE = "stdout"

type Options struct {
	To []string
}

type client struct {
	options Options
	body    string
	logger  log15.Logger
}

var _ providers.ByChat = &client{}

func New(options Options, logger log15.Logger) *client {
	return &client{options: options, logger: logger}
}

func Load(settings []*db.Settings, logger log15.Logger) (interface{}, error) {
	return New(Options{}, logger), nil
}

func (c client) Code() string {
	return CODE
}

func (c client) SetBody(body string) providers.ByChat {
	c.body = body
	return &c
}

func (c client) To(to string, cc ...string) providers.ByChat {
	c.options.To = append([]string{to}, cc...)
	return &c
}

func (c *client) Send() error {
	if len(c.options.To) == 0 {
		return errors.New("Missing to")
	}

	c.logger.Info(
		"send notification",
		"to",
		c.options.To,
		"body",
		c.body,
	)

	return nil
}
