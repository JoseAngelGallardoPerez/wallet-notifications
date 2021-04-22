package notificationsserver

import (
	"context"
	"errors"

	pb "github.com/Confialink/wallet-notifications/rpc/proto/notifications"
	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/eventmanager"
	"github.com/Confialink/wallet-notifications/internal/service"
	push_token "github.com/Confialink/wallet-notifications/internal/service/push-token"
)

// NotificationsHandlerServer implements the notifications service
type NotificationsHandlerServer struct {
	EventManager       *eventmanager.EventManager
	SettingsRepository db.RepositoryInterface
	Logger             log15.Logger
}

func NewNotificationsHandlerServer(
	eventManager *eventmanager.EventManager,
	settingsRepository db.RepositoryInterface,
	logger log15.Logger,
	pushTokenService *push_token.Service,
) *NotificationsHandlerServer {

	// This is a crutch to set a dependency
	service.PushTokenService = pushTokenService

	return &NotificationsHandlerServer{EventManager: eventManager, SettingsRepository: settingsRepository, Logger: logger}
}

func (n *NotificationsHandlerServer) GetSettings(
	ctx context.Context,
	req *pb.SettingsRequest,
) (res *pb.SettingsResponse, err error) {
	logger := n.Logger.New("where", "NotificationsHandlerServer.GetSettings")
	settings := make([]*pb.Setting, 0, len(req.SettingNames))
	for _, name := range req.SettingNames {
		setting, err := n.SettingsRepository.GetByName(name)
		if err != nil {
			logger.Warn("filed to retrieve setting", "setting", name, "error", err)
			continue
		}
		settings = append(settings, &pb.Setting{
			Name:  setting.Name,
			Value: setting.Value,
		})
	}
	return &pb.SettingsResponse{Settings: settings}, nil
}

func (n *NotificationsHandlerServer) GetUserSettings(
	ctx context.Context,
	req *pb.UserSettingsRequest,
) (res *pb.UserSettingsResponse, err error) {
	logger := n.Logger.New("where", "NotificationsHandlerServer.GetUserSettings")
	settings, err := n.SettingsRepository.FindActiveUserOptionsByName(req.NotificationName)
	if err != nil {
		logger.Warn("filed to retrieve user settings", "notificationName", req.NotificationName, "error", err)
		return res, errors.New("filed to retrieve user settings")
	}
	userSettings := make([]*pb.UsersSetting, 0, len(settings))
	for _, setting := range settings {
		userSettings = append(userSettings, &pb.UsersSetting{
			NotificationName: setting.NotificationName,
			Uid:              setting.UID,
		})
	}

	return &pb.UserSettingsResponse{UserSettings: userSettings}, nil
}

func (n *NotificationsHandlerServer) Dispatch(ctx context.Context, req *pb.Request) (res *pb.Response, err error) {
	templateData := req.GetTemplateData()
	td := &db.TemplateData{
		UserName:                       templateData.GetUserName(),
		FirstName:                      templateData.GetFirstName(),
		LastName:                       templateData.GetLastName(),
		OneTimeLoginURL:                templateData.GetOneTimeLoginUrl(),
		PrivateMessageRecipient:        templateData.GetPrivateMessageRecipient(),
		PrivateMessageAuthor:           templateData.GetPrivateMessageAuthor(),
		PrivateMessageURL:              templateData.GetPrivateMessageUrl(),
		PrivateMessageRecipientEditURL: templateData.GetPrivateMessageRecipientEditUrl(),
		Reason:                         templateData.GetReason(),
		Link:                           templateData.GetLink(),
		DocumentName:                   templateData.GetDocumentName(),
		Tan:                            templateData.GetTan(),
		Password:                       templateData.GetPassword(),
		EntityType:                     templateData.GetEntityType(),
		EntityID:                       templateData.GetEntityID(),
		MessageUnreadCount:             templateData.GetMessageUnreadedCount(),
		SenderID:                       templateData.GetSenderID(),
		VerificationLink:               templateData.GetVerificationLink(),
		AccountNumber:                  templateData.GetAccountNumber(),
		TransactionId:                  templateData.GetTransactionId(),
		ConfirmationCode:               templateData.GetConfirmationCode(),
		SetPasswordConfirmationCode:    templateData.GetSetPasswordConfirmationCode(),
		RequestID:                      templateData.GetRequestId(),
		Count:                          templateData.GetCount(),
		InvoiceID:                      templateData.GetInvoiceID(),
		SupplierCompany:                templateData.GetSupplierCompany(),
		FunderCompany:                  templateData.GetFunderCompany(),
		Date:                           templateData.GetDate(),
		PlatformAdmin:                  templateData.GetPlatformAdmin(),
		StaffFirstName:                 templateData.GetStaffFirstName(),
		OwnerFirstName:                 templateData.GetOwnerFirstName(),
		OwnerLastName:                  templateData.GetOwnerLastName(),
	}

	reqMethods := req.GetNotifiers()
	methods := make([]service.Method, len(reqMethods))
	for i, reqMethod := range reqMethods {
		m, err := service.MethodFromString(reqMethod)
		if err != nil {
			return nil, err
		}
		methods[i] = m
	}

	data := service.NewCallData(req.GetTo(), methods, td)

	n.EventManager.Dispatch(req.GetEventName(), *data)
	return &pb.Response{
		Status: "OK",
	}, nil
}
