package firebase

import (
	"github.com/Confialink/wallet-pkg-env_mods"
	"github.com/appleboy/go-fcm"
	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"

	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/service/providers"
)

const (
	CODE = "firebase"

	FCM_MESSAGE_PRIORITY = "high"

	FCM_MESSAGE_NOTIFICATION_ICON         = "fcm_push_icon"
	FCM_MESSAGE_NOTIFICATION_CLICK_ACTION = "FCM_PLUGIN_ACTIVITY"
)

type Options struct {
	authID             string
	authToken          string
	isEnabled          bool
	from               string
	to                 []string
	entityType         string
	entityID           uint64
	messageUnreadCount uint64
}

type client struct {
	options *Options
	subject string
	body    string
	logger  log15.Logger
}

var _ providers.ByPush = &client{}

func New(options *Options, logger log15.Logger) *client {
	return &client{options: options, logger: logger}
}

func Load(settings []*db.Settings, logger log15.Logger) (interface{}, error) {
	options, err := prepareSettings(settings)
	if err != nil {
		return nil, err
	}

	if !options.isEnabled {
		return NewBlackHole(options, logger), nil
	}

	return New(options, logger), nil
}

func (c client) Code() string {
	return CODE
}

func (c client) From(from string) providers.ByPush {
	c.options.from = from
	return &c
}

func (c client) SetBody(body string) providers.ByPush {
	c.body = body
	return &c
}

func (c client) SetSubject(subject string) providers.ByPush {
	c.subject = subject
	return &c
}

func (c client) SetEntityType(entityType string) providers.ByPush {
	c.options.entityType = entityType
	return &c
}

func (c client) SetEntityID(entityID uint64) providers.ByPush {
	c.options.entityID = entityID
	return &c
}

func (c client) SetMessageUnreadCount(messageUnreadCount uint64) providers.ByPush {
	c.options.messageUnreadCount = messageUnreadCount
	return &c
}

func (c client) To(to []string) providers.ByPush {
	c.options.to = to
	return &c
}

func (c *client) Send() error {
	if len(c.options.to) == 0 {
		return errors.New("Missing to")
	}

	for _, receiver := range c.options.to {
		push := &fcm.Message{
			To:               receiver,
			Priority:         FCM_MESSAGE_PRIORITY,
			ContentAvailable: true,
			Notification: &fcm.Notification{
				Title:       c.subject,
				Body:        c.body,
				Icon:        FCM_MESSAGE_NOTIFICATION_ICON,
				ClickAction: FCM_MESSAGE_NOTIFICATION_CLICK_ACTION,
			},
			Data: map[string]interface{}{
				"body":                   c.body,
				"title":                  c.subject,
				"type":                   c.options.entityType,
				"id":                     c.options.entityID,
				"message_unreaded_count": c.options.messageUnreadCount,
			},
		}

		if env_mods.IsDebugMode() {
			c.logger.Info(
				"firebase request",
				"push",
				push,
			)
		}

		if err := c.sendRequest(push); err != nil {
			return err
		}
	}

	return nil
}

// sendRequest send the message and receive the response without retries.
func (c *client) sendRequest(push *fcm.Message) error {
	fcmClient, err := c.getFcmClient()
	if err != nil {
		return err
	}
	resp, err := fcmClient.Send(push)
	if err != nil {
		return errors.Wrap(err, "cannot send a push notification")
	}

	if resp.Failure > 0 {
		err := errors.New("cannot send a push notification")
		for _, res := range resp.Results {
			if res.Error != nil {
				err = errors.Wrap(err, res.Error.Error())
			}
		}
		return err
	}

	if env_mods.IsDebugMode() {
		c.logger.Info(
			"firebase response",
			"response",
			resp,
		)
	}

	return nil
}

// getFcmClient creates a FCM client to send the message.
func (c *client) getFcmClient() (*fcm.Client, error) {
	client, err := fcm.NewClient(c.options.authToken)
	if err != nil {
		return nil, errors.Wrap(err, "can not make fcm client")
	}
	return client, nil
}

// prepareSettings init options
func prepareSettings(settings []*db.Settings) (*Options, error) {
	pushStatus, err := db.GetSettingValue("push_status", settings)
	if err != nil {
		return nil, err
	}

	isEnabled := pushStatus == "enabled" || pushStatus == "1" || pushStatus == "yes"
	if !isEnabled {
		return &Options{isEnabled: false}, nil
	}

	//appID, err := db.GetSettingValue("google_firebase_app_id", settings)
	//if err != nil {
	//	return nil, err
	//}
	authToken, err := db.GetSettingValue("google_firebase_token", settings)
	if err != nil {
		return nil, err
	}

	return &Options{
		//authID:    appID,
		authToken: authToken,
		isEnabled: true,
	}, nil
}
