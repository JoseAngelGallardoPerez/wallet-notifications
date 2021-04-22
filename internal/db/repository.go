package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// RepositoryInterface is interface for repository functionality
// that ought to be implemented manually.
type RepositoryInterface interface {
	GetAll() ([]*Settings, error)
	GetByName(name string) (*Settings, error)
	FirstOrCreate(data *Settings) error
	GetTemplates() ([]*Template, error)
	UpdateTemplate(template *Template) error
	FindByScopeAndEditable(scope string) ([]*Template, error)
	FindOneByTitleAndScope(title, scope string) (*Template, error)
	FindUserOptionByUIDAndName(UID string, notificationName string) (*UserSettings, error)
	FindActiveUserOptionsByName(notificationName string) ([]*UserSettings, error)
	UpdateUserSetting(setting *UserSettings) error
}

// Repository is user repository for CRUD operations.
type Repository struct {
	db *gorm.DB
}

// NewRepository creates new repository
func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{db}
}

// GetAll returns list of settings
func (r *Repository) GetAll() ([]*Settings, error) {
	result := make([]*Settings, 0, 16)
	if err := r.db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) GetByName(name string) (*Settings, error) {
	entity := &Settings{}
	err := r.db.Where("name = ?", name).First(entity).Error
	if nil != err {
		return nil, err
	}
	return entity, nil
}

// FirstOrCreate updates or create if setting don't exists
func (r *Repository) FirstOrCreate(data *Settings) error {
	err := r.db.Where(Settings{Name: data.Name}).
		Assign(Settings{Name: data.Name, Value: data.Value, Description: data.Description}).
		FirstOrCreate(&data).Error
	if nil != err {
		return err
	}
	return nil
}

// GetTemplates returns list of templates
func (r *Repository) GetTemplates() ([]*Template, error) {
	result := make([]*Template, 0, 16)
	if err := r.db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// FindUserOptionByUIDAndName returns user options by UID
func (r *Repository) FindUserOptionByUIDAndName(UID string, notificationName string) (*UserSettings, error) {
	setting := &UserSettings{}
	err := r.db.Where(UserSettings{UID: UID, NotificationName: notificationName}).
		Attrs(UserSettings{IsActive: "1"}).
		FirstOrCreate(&setting).Error
	if nil != err {
		return nil, err
	}
	return setting, nil
}

// find active user option by name
func (r *Repository) FindActiveUserOptionsByName(notificationName string) ([]*UserSettings, error) {
	var settings []*UserSettings
	err := r.db.Where(UserSettings{NotificationName: notificationName, IsActive: "1"}).
		Find(&settings).Error
	if nil != err {
		return nil, err
	}
	return settings, nil
}

// UpdateUserSetting updates an existing template
func (r *Repository) UpdateUserSetting(setting *UserSettings) error {
	if err := r.db.Model(&setting).Where("uid = ? and notification_name = ?", setting.UID, setting.NotificationName).Updates(setting).Error; err != nil {
		return err
	}
	return nil
}

// FindByScope returns templates by scope
func (r *Repository) FindByScopeAndEditable(scope string) ([]*Template, error) {
	var templates []*Template
	if err := r.db.Where("scope = ? AND is_editable = ?", scope, IsEditableTemplateYes).
		Order("sort desc").
		Find(&templates).Error; err != nil {
		return nil, err
	}
	return templates, nil
}

// FindOneByTitleAndScope find template by title and scope
func (r *Repository) FindOneByTitleAndScope(title, scope string) (*Template, error) {
	tmpl := &Template{}
	if err := r.db.Where("title = ? AND scope = ?", title, scope).
		First(&tmpl).Error; err != nil {
		return nil, fmt.Errorf("Could not find template with title `%s` in database", title)
	}
	return tmpl, nil
}

// UpdateTemplate updates an existing template
func (r *Repository) UpdateTemplate(template *Template) error {
	if err := r.db.Model(template).Updates(template).Error; err != nil {
		return err
	}
	return nil
}
