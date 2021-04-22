package models

import (
	"time"
)

const (
	OsTypeIos     = "ios"
	OsTypeAndroid = "android"
	OsTypeOther   = "other"
)

type PushToken struct {
	PushToken string    `gorm:"primary_key:yes;column:push_token;unique_index" json:"pushToken"`
	Os        string    `gorm:"column:os" json:"os"`
	Name      string    `gorm:"column:name;" json:"name"`
	UID       string    `gorm:"column:uid" json:"uid"`
	DeviceId  string    `gorm:"column:device_id" json:"deviceId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (d *PushToken) IsExists() bool {
	return len(d.PushToken) > 0
}
