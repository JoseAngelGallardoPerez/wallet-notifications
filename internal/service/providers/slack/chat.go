package plivo

import (
	"errors"

	"github.com/Confialink/wallet-notifications/internal/service/providers"

	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/inconshreveable/log15"
	// "github.com/nlopes/slack"
)

const CODE = "slack"

type Options struct {
	AuthToken string   `envconfig:"AUTH_TOKEN" required:"true"`
	Channel   []string `envconfig:"CHANNEL"`
}

type client struct {
	options *Options
	body    string
	logger  log15.Logger
}

var _ providers.ByChat = &client{}

func New(options *Options, logger log15.Logger) *client {
	return &client{options: options, logger: logger}
}

func Load(settings []*db.Settings, logger log15.Logger) (interface{}, error) {
	options, err := prepareSettings(settings)
	if err != nil {
		return nil, err
	}
	return New(options, logger), nil
}

func (c client) Code() string {
	return CODE
}

func (c client) SetBody(body string) providers.ByChat {
	c.body = body
	return &c
}

func (c client) To(to string, cc ...string) providers.ByChat {
	c.options.Channel = append([]string{to})
	return &c
}

func (c *client) Send() error {
	if len(c.options.Channel) == 0 {
		return errors.New("Missing to")
	}

	// s := slack.New(c.options.AuthToken)
	// for _, channel := range c.options.Channel {
	// 	if _, _, err := s.PostMessage(channel, c.body, slack.PostMessageParameters{}); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

// prepareSettings init options
func prepareSettings(settings []*db.Settings) (*Options, error) {

	// TODO

	return &Options{}, nil
}
