package db

import (
	"bytes"
	"fmt"
	htmlTemplate "html/template"
	"log"
	"strings"
	"time"

	settingspb "github.com/Confialink/wallet-settings/rpc/proto/settings"
	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
)

const (
	StatusTemplateEnabled  = "enabled"
	StatusTemplateDisabled = "disabled"

	ScopeUser  = "user"
	ScopeAdmin = "admin"

	IsEditableTemplateYes = 1
	IsEditableTemplateNo  = 0

	KeyEmailFrom = "email_from"
	KeyPushStatus = "push_status"
)

type Settings struct {
	Id          uint32    `gorm:"primary_key:yes;column:id;unique_index" json:"id"`
	Name        string    `gorm:"column:name;unique_index" json:"name" binding:"required,max=255"`
	Value       string    `gorm:"column:value;" json:"value" binding:"max=255"`
	Description string    `gorm:"column:description;" json:"description" binding:"max=255"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UserSettings struct {
	Id               uint32 `gorm:"primary_key:yes;column:id;unique_index" json:"id"`
	UID              string `gorm:"column:uid;unique_index" json:"-"`
	NotificationName string `gorm:"column:notification_name;unique_index" json:"notificationName" binding:"required,max=255"`
	IsActive         string `gorm:"not null; default:0" json:"isActive"`
}

type PublicSettings struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Template struct {
	Id        uint32    `gorm:"primary_key:yes;column:id;unique_index" json:"id" binding:"required"`
	Title     string    `gorm:"column:title" json:"title" binding:"required,max=255"`
	Scope     string    `gorm:"column:scope" json:"scope"`
	Legend    string    `gorm:"column:legend" json:"legend" binding:"max=255"`
	Subject   string    `gorm:"column:subject" json:"subject" binding:"required,max=255"`
	Content   string    `gorm:"column:content" json:"content" binding:"required"`
	Status    string    `gorm:"column:status" json:"status" json.enum:"enabled,disabled"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Sort      uint32    `json:"sort"`
}

type TemplateData struct {
	UserName                       string
	Email                          string
	PhoneNumber                    string
	SmsPhoneNumber                 string
	FirstName                      string
	LastName                       string
	SiteName                       string
	SiteLoginURL                   string
	SiteLoginUrl                   string
	Logo                           string
	OneTimeLoginURL                string
	PrivateMessageRecipient        string
	PrivateMessageAuthor           string
	PrivateMessageURL              string
	PrivateMessageRecipientEditURL string
	Reason                         string
	Link                           string
	DocumentName                   string
	Tan                            string
	SiteURL                        string
	SiteUrl                        string
	PasswordRecoveryURL            string
	PasswordRecoveryUrl            string
	Password                       string
	EntityType                     string
	EntityID                       uint64
	MessageUnreadCount             uint64
	SenderID                       string
	VerificationLink               string
	AccountNumber                  string
	TransactionId                  uint64
	ConfirmationCode               string
	SetPasswordConfirmationCode    string
	SetPasswordOneTimeURL          string
	SetPasswordOneTimeUrl          string
	RequestID                      uint64
	Count                          uint64
	InvoiceID                      string
	SupplierCompany                string
	FunderCompany                  string
	Date                           string
	PlatformAdmin                  string
	StaffFirstName                 string
	OwnerFirstName                 string
	OwnerLastName                  string
}

// GetSettingValue is helper function that returns setting value by name
func GetSettingValue(name string, settings []*Settings) (string, error) {
	for i := range settings {
		n := settings[i].Name
		if n == name {
			return settings[i].Value, nil
		}
	}
	return "", fmt.Errorf("setting `%s` not found", name)
}

// IsEnabled checks if template active
func (t *Template) IsEnabled() bool {
	if t.Status == StatusTemplateEnabled {
		return true
	}
	return false
}

func (t *Template) GetSubject() string {
	return t.Subject
}

func (t *Template) PrepareSubject(templateData *TemplateData) (*Template, error) {
	subject, err := templateData.parseText(t.Subject)
	if err != nil {
		log.Fatalf(err.Error())
	}

	t.Subject = subject
	return t, nil
}

func (t *Template) GetContent() string {
	return t.Content
}

func (t *Template) PrepareContent(templateData *TemplateData) (*Template, error) {
	content, err := templateData.parseText(t.Content)
	if err != nil {
		return nil, err
	}

	t.Content = content
	return t, nil
}

func (t *Template) ApplyTemplateData(templateData *TemplateData) (*Template, error) {
	t, err := t.PrepareSubject(templateData)
	if err != nil {
		return nil, err
	}
	t, err = t.PrepareContent(templateData)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (td *TemplateData) parseText(text string) (string, error) {
	//replace "[value]" -> "{{.value}}"
	replacer := strings.NewReplacer("[", "{{.", "]", "}}")
	text = replacer.Replace(text)
	t, err := htmlTemplate.New("text").Parse(text)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, td); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// ApplySystemSettings applied system settings to template data
func (td *TemplateData) ApplySystemSettings(settings []*settingspb.Setting) error {
	for _, s := range settings {
		if s.Path == "regional/general/site_name" {
			td.SiteName = s.Value
		}

		if s.Path == "regional/general/site_url" {
			baseurl := s.Value
			td.SiteURL = baseurl
			td.SiteUrl = baseurl
			td.SiteLoginURL = fmt.Sprintf("%s/log-in", baseurl)
			td.SiteLoginUrl = fmt.Sprintf("%s/log-in", baseurl)
			td.PasswordRecoveryURL = fmt.Sprintf("%s/password-recovery", baseurl)
			td.PasswordRecoveryUrl = fmt.Sprintf("%s/password-recovery", baseurl)
			td.VerificationLink = fmt.Sprintf("%s/idverification/%s", baseurl, td.VerificationLink)
			td.SetPasswordOneTimeURL = fmt.Sprintf("%s/password-set/%s", baseurl, td.SetPasswordConfirmationCode)
			td.SetPasswordOneTimeUrl = fmt.Sprintf("%s/password-set/%s", baseurl, td.SetPasswordConfirmationCode)
		}
	}
	return nil
}

// ApplyUser applied user info to template data
func (td *TemplateData) ApplyUser(user *userpb.User) error {
	if td.UserName == "" {
		td.UserName = user.Username
	}
	td.Email = user.Email
	td.PhoneNumber = user.PhoneNumber
	td.SmsPhoneNumber = user.SmsPhoneNumber
	td.FirstName = user.FirstName
	td.LastName = user.LastName
	return nil
}

// PrepareText applied template data to text
func (td *TemplateData) PrepareText(text string) (string, error) {
	result, err := td.parseText(text)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return result, nil
}
