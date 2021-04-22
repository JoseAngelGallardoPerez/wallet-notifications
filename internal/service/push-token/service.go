package push_token

import (
	"time"

	list_params "github.com/Confialink/wallet-pkg-list_params"
	"github.com/pkg/errors"

	"github.com/Confialink/wallet-notifications/internal/db/dao"
	"github.com/Confialink/wallet-notifications/internal/db/models"
	"github.com/Confialink/wallet-notifications/internal/validators"
)

const (
	coefficientForOldPushTokens = 1.5
)

type Service struct {
	dao                   *dao.PushToken
	pushTokenTTL          time.Duration
	expiredTokensDuration time.Duration
}

func NewService(dao *dao.PushToken, pushTokenTTL time.Duration) *Service {
	// We calculate a duration to delete expired tokens by a special formula: we add a coefficient to store
	// expired tokens some more than expiration time is.
	d := time.Duration(int64(float64(pushTokenTTL.Nanoseconds()) * coefficientForOldPushTokens))
	return &Service{dao, pushTokenTTL, d}
}

// AddOrUpdate add a new push token to the DB or update existing one
func (s *Service) AddOrUpdate(in *validators.AddPushToken, userID string) (*models.PushToken, error) {
	if in.Os != models.OsTypeIos && in.Os != models.OsTypeAndroid {
		in.Os = models.OsTypeOther
	}
	pushToken, err := s.dao.FindOneByPushToken(in.PushToken)
	if err != nil {
		return nil, errors.Wrap(err, "cannot find push token")
	}

	// we update the PushToken data every time because many users can login into one device
	pushToken.UID = userID
	pushToken.Name = in.Name
	pushToken.Os = in.Os
	pushToken.DeviceId = in.DeviceId

	if pushToken.IsExists() {
		if err := s.dao.Save(pushToken); err != nil {
			return nil, errors.Wrap(err, "cannot save push token")
		}
		return pushToken, nil
	}

	pushToken.PushToken = in.PushToken
	if err := s.dao.Create(pushToken); err != nil {
		return nil, errors.Wrap(err, "cannot create push token")
	}

	return pushToken, nil
}

// RemovePushToken removes provided push token from the DB
func (s *Service) RemovePushToken(pushToken string) error {
	if err := s.dao.Delete(pushToken); err != nil {
		return errors.Wrap(err, "cannot remove push token")
	}

	return nil
}

// FindOne returns one push token
func (s *Service) FindOne(pushToken string) (*models.PushToken, error) {
	return s.dao.FindOneByPushToken(pushToken)
}

// NotExpiredTokensByUser returns not expired push tokens for a user
func (s *Service) NotExpiredTokensByUser(userID string) ([]*models.PushToken, error) {
	deadLine := time.Now().Local().Add(-s.pushTokenTTL)
	pushTokens, err := s.dao.FindNotExpiredByUID(userID, deadLine)
	if err != nil {
		return nil, errors.Wrap(err, "cannot find not expired push tokens")
	}

	return pushTokens, nil
}

// DeleteExpiredTokens deletes expired push tokens from the DB
func (s *Service) DeleteExpiredTokens() error {
	deadLine := time.Now().Local().Add(-s.expiredTokensDuration)
	if err := s.dao.DeleteByDeadline(deadLine); err != nil {
		return errors.Wrap(err, "cannot remove expired push tokens")
	}

	return nil
}

func (s *Service) List(params *list_params.ListParams) ([]*models.PushToken, error) {
	return s.dao.List(params)
}

func (s *Service) Count(params *list_params.ListParams) (uint64, error) {
	return s.dao.Count(params)
}
