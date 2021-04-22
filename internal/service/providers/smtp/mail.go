package smtp

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/service/providers"
	"github.com/Confialink/wallet-pkg-gomail"
	"github.com/inconshreveable/log15"
)

const CODE = "smtp"

type Options struct {
	status   string
	username string
	password string
	host     string
	port     int
	protocol string
	from     string
	fromName string
	to       []string
}

type client struct {
	options *Options
	subject string
	body    string
	logger  log15.Logger
}

var _ providers.ByEmail = &client{}

// New creates new client
func New(options *Options, logger log15.Logger) *client {
	return &client{options: options, logger: logger}
}

// Load init client
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

func (c client) From(from string) providers.ByEmail {
	c.options.from = from
	return &c
}

func (c client) To(to string, cc ...string) providers.ByEmail {
	c.options.to = append([]string{to}, cc...)
	return &c
}

func (c client) SetSubject(subject string) providers.ByEmail {
	c.subject = subject
	return &c
}

func (c client) SetBody(body string) providers.ByEmail {
	c.body = body
	return &c
}

// Send sends email
func (c *client) Send() error {
	if len(c.options.to) == 0 {
		return errors.New("missing to")
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", mail.FormatAddress(c.options.from, c.options.fromName))
	mail.SetHeader("To", c.options.to...)
	mail.SetHeader("Subject", c.subject)
	mail.SetBody("text/html", c.body)

	var d *gomail.Dialer
	if c.options.status != "disabled" && c.options.username != "" && c.options.password != "" {
		d = gomail.NewDialer(c.options.host, c.options.port, c.options.username, c.options.password)
	} else {
		d = &gomail.Dialer{Host: c.options.host, Port: c.options.port}
	}
	return d.DialAndSend(mail)
}

// prepareSettings init options
func prepareSettings(settings []*db.Settings) (*Options, error) {

	status, err := db.GetSettingValue("smtp_status", settings)
	if err != nil {
		return nil, err
	}

	username, err := db.GetSettingValue("smtp_username", settings)
	if err != nil {
		return nil, err
	}

	password, err := db.GetSettingValue("smtp_password", settings)
	if err != nil {
		return nil, err
	}

	host, err := db.GetSettingValue("smtp_host", settings)
	if err != nil {
		return nil, err
	}

	port, err := db.GetSettingValue("smtp_port", settings)
	if err != nil {
		return nil, err
	}

	protocol, err := db.GetSettingValue("smtp_protocol", settings)
	if err != nil {
		return nil, err
	}

	from, err := db.GetSettingValue("email_from", settings)
	if err != nil {
		return nil, err
	}

	fromName, err := db.GetSettingValue("email_from_name", settings)
	if err != nil {
		return nil, err
	}

	return &Options{
		status:   status,
		username: username,
		password: password,
		host:     host,
		port:     stringToUint64(port),
		protocol: protocol,
		from:     from,
		fromName: fromName,
	}, nil
}

func stringToUint64(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
