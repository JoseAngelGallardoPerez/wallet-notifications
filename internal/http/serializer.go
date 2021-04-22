package http

import "github.com/Confialink/wallet-notifications/internal/db"

type Serializer struct {
}

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (s *Serializer) Serialize(settingEntity *db.Settings) *db.PublicSettings {
	serializedModel := &db.PublicSettings{
		Name:  settingEntity.Name,
		Value: settingEntity.Value,
	}

	return serializedModel
}
