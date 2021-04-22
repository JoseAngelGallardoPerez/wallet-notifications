package di

import (
	"log"

	"github.com/Confialink/wallet-pkg-service_names"
	"github.com/inconshreveable/log15"
	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"

	"github.com/Confialink/wallet-notifications/internal/config"
	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/db/dao"
	"github.com/Confialink/wallet-notifications/internal/eventmanager"
	"github.com/Confialink/wallet-notifications/internal/http"
	"github.com/Confialink/wallet-notifications/internal/service"
	push_token "github.com/Confialink/wallet-notifications/internal/service/push-token"
	"github.com/Confialink/wallet-notifications/internal/service/workers"
)

// container is the implementation of the Container interface.
type container struct {
	config           *config.Configuration
	httpService      *http.Service
	router           *http.Router
	repository       db.RepositoryInterface
	daoPushToken     *dao.PushToken
	servicePushToken *push_token.Service
	handlerPushToken *http.PushTokenHandler
	dbConnection     *gorm.DB
	response         *http.ResponseService
	serializer       *http.Serializer
	notifier         *service.Notifier
	eventManager     *eventmanager.EventManager
	loggerService    log15.Logger
	worker           *workers.Worker
}

// Container represents a dependency injection container.
var Container *container

func init() {
	Container = new(container)

	// Retrieve config options.
	config.InitConfig(Container.LoggerService().New("service", "configReader"))
	Container.config = config.GetConf()
}

// Config returns config
func (c *container) Config() *config.Configuration {
	return c.config
}

// Router creates new router if not exists and return
func (c *container) Router() *http.Router {
	if nil == c.router {
		c.router = http.NewRouter(c.HttpService(), c.Config().GetCors(), c.LoggerService(), c.HandlerPushToken())
	}
	return c.router
}

// HttpService creates new http service if not exists and return
func (c *container) HttpService() *http.Service {
	if nil == c.httpService {
		c.httpService = http.NewService(c.Repository(), c.Response(), c.Notifier(), c.EventManager(), c.Serializer(), c.LoggerService())
	}
	return c.httpService
}

// Repository creates new repository if not exists and return
func (c *container) Repository() db.RepositoryInterface {
	if nil == c.repository {
		c.repository = db.NewRepository(c.DbConnection())
	}
	return c.repository
}

// DbConnection creates new DB connection if not exists and return
func (c *container) DbConnection() *gorm.DB {
	var err error
	if nil == c.dbConnection {
		c.dbConnection, err = db.CreateConnection(c.Config().GetDatabase())
		if nil != err {
			log.Fatalf("Could not connect to DB: %v", err)
		}
	}
	return c.dbConnection
}

// Response creates new response if not exists and return
func (c *container) Response() *http.ResponseService {
	if nil == c.response {
		c.response = http.NewResponseService()
	}
	return c.response
}

// Notifier creates new mail service if not exists and return
func (c *container) Notifier() *service.Notifier {
	if nil == c.notifier {
		c.notifier = service.NewNotifier(c.Repository(), c.LoggerService())
	}
	return c.notifier
}

func (c *container) EventManager() *eventmanager.EventManager {
	if nil == c.eventManager {
		storage := eventmanager.NewMemoryStorage()
		dispatcher := eventmanager.NewDispatcher()
		repo := c.Repository()
		notifier := c.Notifier()

		c.eventManager = eventmanager.NewEventManager(
			storage, dispatcher, notifier, repo, c.LoggerService().New("service", "EventManager"),
		)

		c.eventManager.Listen()
	}
	return c.eventManager
}

func (c *container) LoggerService() log15.Logger {
	if c.loggerService == nil {
		c.loggerService = log15.New("service", service_names.Notifications.Internal)
	}
	return c.loggerService
}

func (c *container) Serializer() *http.Serializer {
	if nil == c.serializer {
		c.serializer = http.NewSerializer()
	}
	return c.serializer
}

func (c *container) HandlerPushToken() *http.PushTokenHandler {
	if nil == c.handlerPushToken {
		c.handlerPushToken = http.NewPushTokenHandler(c.Response(), c.ServicePushToken(), c.Repository(), c.LoggerService())
	}
	return c.handlerPushToken
}

func (c *container) ServicePushToken() *push_token.Service {
	if nil == c.servicePushToken {
		c.servicePushToken = push_token.NewService(c.DaoPushToken(), c.Config().PushTokenTTL)
	}
	return c.servicePushToken
}

func (c *container) DaoPushToken() *dao.PushToken {
	if nil == c.daoPushToken {
		c.daoPushToken = dao.NewPushToken(c.DbConnection())
	}
	return c.daoPushToken
}

func (c *container) Worker() *workers.Worker {
	if nil == c.worker {
		c.worker = workers.NewWorker(gocron.NewScheduler(), c.LoggerService(), c.ServicePushToken())
	}
	return c.worker
}
