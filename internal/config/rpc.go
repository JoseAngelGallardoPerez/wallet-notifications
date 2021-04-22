package config

import (
	"github.com/Confialink/wallet-pkg-env_config"
)

// RPCConfiguration is rpc config model
type RPCConfiguration struct {
	NotificationsServerPort string
}

// GetNotificationsServerPort returns rpc port for userserver
func (s *RPCConfiguration) GetNotificationsServerPort() string {
	return s.NotificationsServerPort
}

// Init initializes enviroment variables
func (s *RPCConfiguration) Init() error {
	s.NotificationsServerPort = env_config.Env("VELMIE_WALLET_NOTIFICATIONS_RPC_PORT", "")
	return nil
}
