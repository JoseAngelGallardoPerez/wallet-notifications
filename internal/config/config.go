package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Confialink/wallet-pkg-env_config"
	"github.com/inconshreveable/log15"
)

// Configuration is microservice config model
type Configuration struct {
	Server       *ServerConfiguration
	Database     *env_config.Db
	Cors         *env_config.Cors
	RPC          *RPCConfiguration
	PushTokenTTL time.Duration
}

// Create a new config instance.
var conf *Configuration

// init initializes conf variables
func InitConfig(logger log15.Logger) {
	defaultConfigReader := env_config.NewReader("notifications")

	// Configure server.
	server := &ServerConfiguration{}
	server.Init()

	// Configure database settings.
	database := defaultConfigReader.ReadDbConfig()

	// Configure cors settings.
	cors := defaultConfigReader.ReadCorsConfig()

	// Configure rpc settings.
	rpc := &RPCConfiguration{}
	rpc.Init()

	conf = &Configuration{
		Server:       server,
		Database:     database,
		Cors:         cors,
		RPC:          rpc,
		PushTokenTTL: pushTokenTTL(),
	}

	validateConfig(conf, logger)
}

// GetConf returns the conf Configuration struct.
func GetConf() *Configuration {
	return conf
}

// GetServer returns server configuration
func (s *Configuration) GetServer() *ServerConfiguration {
	return s.Server
}

// GetDatabase returns database configuration
func (s *Configuration) GetDatabase() *env_config.Db {
	return s.Database
}

// GetCors returns cors configuration
func (s *Configuration) GetCors() *env_config.Cors {
	return s.Cors
}

// GetRPC returns RPC configuration
func (s *Configuration) GetRPC() *RPCConfiguration {
	return s.RPC
}

func validateConfig(cfg *Configuration, logger log15.Logger) {
	validator := env_config.NewValidator(logger)
	validator.ValidateCors(cfg.Cors, logger)
	validator.ValidateDb(cfg.Database, logger)
	validator.CriticalIfEmpty(cfg.Server.GetPort(), "VELMIE_WALLET_NOTIFICATIONS_SERVER_PORT", logger)
	validator.CriticalIfEmpty(cfg.RPC.GetNotificationsServerPort(), "VELMIE_WALLET_NOTIFICATIONS_RPC_PORT", logger)
}

func pushTokenTTL() time.Duration {
	env := "VELMIE_WALLET_NOTIFICATIONS_PUSH_TOKEN_TTL_DAYS"
	pushTokenTTLDays := env_config.Env(env, "14") // 2 weeks by default
	pushTokenTTLDaysInt, err := strconv.Atoi(pushTokenTTLDays)
	if err != nil {
		panic(fmt.Sprintf("Please check the `%s`. It must be an integer value. Error: %s", env, err.Error()))
	}
	if pushTokenTTLDaysInt < 1 {
		panic(fmt.Sprintf("Please check the `%s`. It must be a positive value", env))
	}

	return time.Duration((time.Hour * 24).Nanoseconds() * int64(pushTokenTTLDaysInt))
}
