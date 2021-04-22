package dao

import (
	"time"

	list_params "github.com/Confialink/wallet-pkg-list_params"
	"github.com/Confialink/wallet-pkg-list_params/adapters"
	"github.com/jinzhu/gorm"

	"github.com/Confialink/wallet-notifications/internal/db/models"
)

const (
	pushTokenLimitForOneUser = 500 // FCM has this limit
)

type PushToken struct {
	db *gorm.DB
}

func NewPushToken(db *gorm.DB) *PushToken {
	return &PushToken{db: db}
}

// FindByUID find push tokens by uid
func (r *PushToken) FindByUID(uid string) ([]*models.PushToken, error) {
	var devices []*models.PushToken
	if err := r.db.Where("uid = ?", uid).Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}

// FindNotExpiredByUID returns tokens for a user according to deadline
func (r *PushToken) FindNotExpiredByUID(uid string, deadLine time.Time) ([]*models.PushToken, error) {
	var devices []*models.PushToken
	if err := r.db.Where("uid = ? AND updated_at > ?", uid, deadLine).Limit(pushTokenLimitForOneUser).
		Order("updated_at desc").Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}

// FindOneByPushToken returns a record by the provided pushToken
func (r *PushToken) FindOneByPushToken(pushToken string) (*models.PushToken, error) {
	device := &models.PushToken{}
	err := r.db.Where("push_token = ?", pushToken).FirstOrInit(&device).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return device, nil
}

// Create creates a new record in the DB
func (r *PushToken) Create(record *models.PushToken) error {
	if err := r.db.Create(record).Error; err != nil {
		return err
	}
	return nil
}

// Save updates the record in the DB
func (r *PushToken) Save(record *models.PushToken) error {
	if err := r.db.Save(record).Error; err != nil {
		return err
	}
	return nil
}

// Delete deletes a record by a provided push token
func (r *PushToken) Delete(pushToken string) error {
	if err := r.db.Where("push_token = ?", pushToken).Delete(&models.PushToken{}).Error; err != nil {
		return err
	}
	return nil
}

// DeleteByDeadline deletes push tokens from the DB by deadline
func (r *PushToken) DeleteByDeadline(deadLine time.Time) error {
	if err := r.db.Where("updated_at < ?", deadLine).Delete(&models.PushToken{}).Error; err != nil {
		return err
	}
	return nil
}

// List returns list of push tokens by a list of parameters
func (r *PushToken) List(params *list_params.ListParams) ([]*models.PushToken, error) {
	var records []*models.PushToken
	adapter := adapters.NewGorm(r.db)
	err := adapter.LoadList(&records, params, "push_tokens")

	return records, err
}

// Count returns a count of push tokens by a list of parameters
func (r *PushToken) Count(params *list_params.ListParams) (uint64, error) {
	var count uint64
	str, arguments := params.GetWhereCondition()
	query := r.db.Where(str, arguments...)
	query = query.Joins(params.GetJoinCondition())

	if err := query.Model(&models.PushToken{}).Count(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}
