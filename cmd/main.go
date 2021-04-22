package main

import (
	"log"

	"github.com/Confialink/wallet-pkg-env_mods"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"

	"github.com/Confialink/wallet-notifications/internal/app/di"
	rpcserver "github.com/Confialink/wallet-notifications/rpc/cmd/server/notificationsserver"
)

func main() {

	c := di.Container
	serverConfig := c.Config().GetServer()

	ginMode := env_mods.GetMode(serverConfig.GetEnv())
	gin.SetMode(ginMode)

	// Run the rpc server.
	rpc := rpcserver.NewNotificationsServer(
		c.EventManager(),
		c.Repository(),
		c.LoggerService().New("service", "rpc"),
		c.Config(),
		c.ServicePushToken(),
	)
	rpc.Init()

	if err := c.Worker().Start(gocron.NewScheduler()); err != nil {
		log.Fatalf("cannot run worker: %s", err.Error())
	}

	// Register routes to be used.
	ginRouter := c.Router().RegisterRoutes()

	log.Printf("Starting API on port: %s", serverConfig.GetPort())

	// Run the server
	ginRouter.Run(":" + serverConfig.GetPort())
}
