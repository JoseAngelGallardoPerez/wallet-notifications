package firebase

import (
	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-notifications/internal/service/providers"
)

// This firebase client does not send notifications, it just writes to log
type blackHole struct {
	options *Options
	logger  log15.Logger
}

func NewBlackHole(options *Options, logger log15.Logger) *blackHole {
	return &blackHole{options: options, logger: logger.New("firebaseClient", "BlackHole")}
}

func (s *blackHole) Code() string {
	return "BlackHole"
}

func (s *blackHole) From(from string) providers.ByPush {
	return s
}

func (s *blackHole) SetBody(body string) providers.ByPush {
	return s
}

func (s *blackHole) SetSubject(subject string) providers.ByPush {
	return s
}

func (s *blackHole) SetEntityType(entityType string) providers.ByPush {
	return s
}

func (s *blackHole) SetEntityID(entityID uint64) providers.ByPush {
	return s
}

func (s *blackHole) SetMessageUnreadCount(messageUnreadCount uint64) providers.ByPush {
	return s
}

func (s *blackHole) To(to []string) providers.ByPush {
	s.options.to = to
	return s
}

// Send does not send a notification, it writes to log in debug mode
func (s *blackHole) Send() error {
	if len(s.options.to) == 0 {
		return nil
	}
	for _, receiver := range s.options.to {
		s.logger.Info(
			"push a firebase notification into a black hole",
			"token",
			receiver,
		)
	}

	return nil
}
