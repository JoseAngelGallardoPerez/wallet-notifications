package usersserver

import (
	"fmt"
	"log"
	"net/http"

	pb "github.com/Confialink/wallet-notifications/rpc/proto/notifications"
	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-notifications/internal/config"
	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/eventmanager"
	push_token "github.com/Confialink/wallet-notifications/internal/service/push-token"
	server "github.com/Confialink/wallet-notifications/rpc/internal/notificationsserver"
)

// NotificationsServer implements the notifications service
type NotificationsServer struct {
	EventManager       *eventmanager.EventManager
	SettingsRepository db.RepositoryInterface
	Logger             log15.Logger
	Config             *config.Configuration
	pushTokenService   *push_token.Service
}

func NewNotificationsServer(
	eventManager *eventmanager.EventManager,
	settingsRepository db.RepositoryInterface,
	logger log15.Logger,
	config *config.Configuration,
	pushTokenService *push_token.Service,
) *NotificationsServer {
	return &NotificationsServer{EventManager: eventManager, SettingsRepository: settingsRepository, Logger: logger, Config: config, pushTokenService: pushTokenService}
}

// Init initializes users rpc server
func (n *NotificationsServer) Init() {

	srv := server.NewNotificationsHandlerServer(n.EventManager, n.SettingsRepository, n.Logger, n.pushTokenService)
	twirpHandler := pb.NewNotificationHandlerServer(srv, nil)

	mux := http.NewServeMux()
	mux.Handle(pb.NotificationHandlerPathPrefix, twirpHandler)

	log.Printf("Starting rpc on port: %s ", n.Config.GetRPC().GetNotificationsServerPort())

	go http.ListenAndServe(fmt.Sprintf(":%s", n.Config.GetRPC().GetNotificationsServerPort()), mux)
}
