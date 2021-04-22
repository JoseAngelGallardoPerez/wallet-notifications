package africastalking

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/service/providers"

	"github.com/inconshreveable/log15"
)

const (
	CODE = "africatakling"
)

// Service is a model
type Service struct {
	options *Options
	logger  log15.Logger
}

type Options struct {
	Body      string
	AuthToken string
	From      string
	To        []string
	UserName  string
	APIKey    string
	ShortCode string
	Env       string
}

// NewService returns a new service
func NewService(options *Options, logger log15.Logger) Service {
	return Service{options, logger}
}

func Load(settings []*db.Settings, logger log15.Logger) (interface{}, error) {
	options, err := prepareSettings(settings)
	if err != nil {
		return nil, err
	}
	return NewService(options, logger), nil
}

func (service Service) newPostRequest(url string, values url.Values, headers map[string]string) (*http.Response, error) {
	reader := strings.NewReader(values.Encode())

	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Length", strconv.Itoa(reader.Len()))
	req.Header.Set("apikey", service.options.APIKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	return client.Do(req)
}

// prepareSettings init options
func prepareSettings(settings []*db.Settings) (*Options, error) {

	username, err := db.GetSettingValue("africastalking_username", settings)
	if err != nil {
		return nil, err
	}

	apiCode, err := db.GetSettingValue("africastalking_api_code", settings)
	if err != nil {
		return nil, err
	}

	shortCode, err := db.GetSettingValue("africastalking_short_code", settings)
	if err != nil {
		return nil, err
	}

	env, err := db.GetSettingValue("africastalking_env", settings)
	if err != nil {
		return nil, err
	}

	return &Options{
		UserName:  username,
		ShortCode: shortCode,
		APIKey:    apiCode,
		Env:       env,
	}, nil
}

func (service Service) Code() string {
	return CODE
}

func (service Service) From(from string) providers.BySMS {
	service.options.From = from
	return &service
}

func (service Service) SetBody(body string) providers.BySMS {
	service.options.Body = body
	return &service
}

func (service Service) To(to string, cc ...string) providers.BySMS {
	service.options.To = append([]string{to}, cc...)
	return &service
}

func (service Service) Send() error {
	values := url.Values{}
	values.Set("username", service.options.UserName)
	values.Set("to", strings.Join(service.options.To, ","))
	values.Set("message", service.options.Body)
	if service.options.ShortCode != "" {
		// set from = "" to avoid this
		values.Set("from", service.options.ShortCode)
	}

	smsURL := GetSmsURL(service.options.Env)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	res, err := service.newPostRequest(smsURL, values, headers)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var smsMessageResponse SendMessageResponse
	if err := json.NewDecoder(res.Body).Decode(&smsMessageResponse); err != nil {
		return errors.New("unable to parse sms response")
	}
	return nil
}
