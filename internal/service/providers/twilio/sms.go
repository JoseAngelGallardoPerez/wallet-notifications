package twilio

import (
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/inconshreveable/log15"
	"github.com/kevinburke/rest"
	twilioPkg "github.com/kevinburke/twilio-go"

	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/errcodes"
	"github.com/Confialink/wallet-notifications/internal/service/providers"
)

const (
	CODE = "twilio"

	settingNameAccountSid = "twilio_account_sid"
	settingNameAuthToken  = "twilio_auth_token"
	settingNameSmsFrom    = "twilio_sms_from"
)

var (
	twilioErrorTitles = map[string]string{
		"Authentication Error - invalid username": errcodes.InvalidAPIKeys,
	}
	twilioErrorId = map[string]string{
		"20003": errcodes.InvalidAPIKeys,
		"20404": errcodes.InvalidAPIKeys,
	}
)

var defaultClient = &http.Client{
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
	Timeout: 10 * time.Second,
}

type Options struct {
	AccountSid string
	AuthToken  string
	SmsFrom    string
	To         string
}

type Client struct {
	options *Options
	body    string
	logger  log15.Logger
}

func New(options *Options, logger log15.Logger) *Client {
	return &Client{options: options, logger: logger}
}

func Load(settings []*db.Settings, logger log15.Logger) (interface{}, error) {
	options, err := prepareSettings(settings)
	if err != nil {
		return nil, err
	}
	return New(options, logger), nil
}

func (s Client) Code() string {
	return CODE
}

func (s Client) From(from string) providers.BySMS {
	s.options.SmsFrom = from
	return &s
}

func (s Client) SetBody(body string) providers.BySMS {
	s.body = body
	return &s
}

func (s Client) To(to string, cc ...string) providers.BySMS {
	s.options.To = to
	return &s
}

func (s *Client) Send() error {
	if len(s.options.To) == 0 {
		return errors.New("missing to")
	}

	client := twilioPkg.NewClient(s.options.AccountSid, s.options.AuthToken, defaultClient)
	_, err := client.Messages.SendMessage(s.options.SmsFrom, s.options.To, s.body, nil)
	if err != nil {
		s.logger.Error("cannot send sms", "err", err)
		return s.attemptMakePublicErr(err)
	}

	return nil
}

// prepareSettings init options
func prepareSettings(settings []*db.Settings) (*Options, error) {
	accSid, err := db.GetSettingValue(settingNameAccountSid, settings)
	if err != nil {
		return nil, err
	}

	authToken, err := db.GetSettingValue(settingNameAuthToken, settings)
	if err != nil {
		return nil, err
	}

	from, err := db.GetSettingValue(settingNameSmsFrom, settings)
	if err != nil {
		return nil, err
	}

	return &Options{
		AccountSid: accSid,
		AuthToken:  authToken,
		SmsFrom:    from,
		To:         "",
	}, nil
}

// attemptMakePublicErr tries to convert received error to a public error
func (s *Client) attemptMakePublicErr(err error) error {
	er, ok := err.(*rest.Error)
	if !ok {
		return err
	}

	errCode, ok := twilioErrorTitles[er.Title]
	if ok {
		return errcodes.CreatePublicError(errCode, er.Title)
	}

	errCode, ok = twilioErrorId[er.ID]
	if ok {
		return errcodes.CreatePublicError(errCode, er.Title)
	}

	return err
}
