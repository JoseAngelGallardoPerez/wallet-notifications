package plivo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	walletErrors "github.com/Confialink/wallet-pkg-errors"

	"github.com/Confialink/wallet-notifications/internal/errcodes"
	"github.com/Confialink/wallet-notifications/internal/service/providers"

	"github.com/inconshreveable/log15"
	"github.com/shopspring/decimal"

	"github.com/Confialink/wallet-notifications/internal/db"
)

const (
	CODE                      = "plivo"
	singleSmsUrlTemplate      = "https://api.plivo.com/v1/Account/%v/Message/"
	accountDetailsUrlTemplate = "https://api.plivo.com/v1/Account/%v/"
)

type Options struct {
	AuthID     string
	AuthToken  string
	From       string
	To         []string
	MinBalance string
}

type Client struct {
	options *Options
	body    string
	logger  log15.Logger
}

type Account struct {
	AccountType  string `json:"account_type,omitempty"`
	Address      string `json:"address,omitempty"`
	ApiId        string `json:"api_id,omitempty"`
	AuthId       string `json:"auth_id,omitempty"`
	AutoRecharge bool   `json:"auto_recharge,omitempty"`
	BillingMode  string `json:"billing_mode,omitempty"`
	CashCredits  string `json:"cash_credits,omitempty"`
	City         string `json:"city,omitempty"`
	Name         string `json:"name,omitempty"`
	ResourceUri  string `json:"resource_uri,omitempty"`
	State        string `json:"state,omitempty"`
	Timezone     string `json:"timezone,omitempty"`
	Error        string `json:"error,omitempty"`
}

var _ providers.BySMS = &Client{}

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

func (c Client) Code() string {
	return CODE
}

func (c Client) From(from string) providers.BySMS {
	c.options.From = from
	return &c
}

func (c Client) SetBody(body string) providers.BySMS {
	c.body = body
	return &c
}

func (c Client) To(to string, cc ...string) providers.BySMS {
	c.options.To = append([]string{to}, cc...)
	return &c
}

func (c *Client) Send() error {
	if len(c.options.To) == 0 {
		return errors.New("Missing to")
	}

	hc := &http.Client{Timeout: time.Second * 10}

	url := fmt.Sprintf(singleSmsUrlTemplate, c.options.AuthID)
	params := struct {
		Src  string `json:"src,omitempty"`
		Dst  string `json:"dst,omitempty"`
		Text string `json:"text,omitempty"`
	}{
		c.options.From,
		"",
		c.body,
	}
	for _, dst := range c.options.To {
		params.Dst = dst
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(params); err != nil {
			return err
		}

		req, err := http.NewRequest("POST", url, buf)
		if err != nil {
			return err
		}

		req.SetBasicAuth(c.options.AuthID, c.options.AuthToken)
		req.Header.Add("Content-Type", "application/json")

		acc, err := c.getAccountDetails()
		if err != nil {
			return err
		}

		currentBalance, err := decimal.NewFromString(acc.CashCredits)
		if err != nil {
			return err
		}

		minBalance, err := decimal.NewFromString(c.options.MinBalance)

		if currentBalance.LessThan(minBalance) {
			return errors.New("insufficient funds, plivo balance is less than minimum balance")
		}

		resp, err := hc.Do(req)
		if err != nil {
			return err
		}

		respBody := struct {
			Error string `json:"error"`
		}{}
		if err = json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
			_ = resp.Body.Close()
			return err
		}
		if respBody.Error != "" {
			_ = resp.Body.Close()
			return &walletErrors.PublicError{
				Code:       errcodes.ErrorSendingTestSms,
				Title:      respBody.Error,
				HttpStatus: http.StatusBadRequest,
			}
		}

		if err := resp.Body.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) GetAccountDetails() (*Account, error) {
	return c.getAccountDetails()
}

// prepareSettings init options
func prepareSettings(settings []*db.Settings) (*Options, error) {

	authID, err := db.GetSettingValue("plivo_auth_id", settings)
	if err != nil {
		return nil, err
	}

	authToken, err := db.GetSettingValue("plivo_auth_token", settings)
	if err != nil {
		return nil, err
	}

	from, err := db.GetSettingValue("sms_from", settings)
	if err != nil {
		return nil, err
	}

	minBalance, err := db.GetSettingValue("plivo_min_balance", settings)
	if err != nil {
		return nil, err
	}

	return &Options{
		AuthID:     authID,
		AuthToken:  authToken,
		From:       from,
		MinBalance: minBalance,
	}, nil
}

// getAccountDetails get account details from plivo api
func (c *Client) getAccountDetails() (*Account, error) {
	hc := &http.Client{Timeout: time.Second * 5}

	url := fmt.Sprintf(accountDetailsUrlTemplate, c.options.AuthID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.options.AuthID, c.options.AuthToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body := make([]byte, 128)
		n, _ := resp.Body.Read(body)
		title := fmt.Sprintf("plivo responded with: %s", string(body[:n]))
		return nil, errcodes.CreatePublicError(errcodes.InvalidAPIKeys, title)
	}

	var acc Account
	if err = json.NewDecoder(resp.Body).Decode(&acc); err != nil {
		return nil, errcodes.CreatePublicError(errcodes.CanNotRetrieveProviderDetails, "failed to parse response body: "+err.Error())
	}

	if acc.Error != "" {
		return nil, errors.New(acc.Error)
	}

	return &acc, nil
}
