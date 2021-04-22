package users

import (
	"github.com/Confialink/wallet-notifications/internal/srvdiscovery"
	"context"
	"net/http"

	pb "github.com/Confialink/wallet-users/rpc/proto/users"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (u *Service) GetByUID(uid string) (*pb.User, error) {
	req := pb.Request{UID: uid}
	client, err := u.getClient()
	if err != nil {
		return nil, err
	}
	resp, err := client.GetByUID(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}

func (u *Service) getClient() (pb.UserHandler, error) {
	usersUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNameUsers)
	if nil != err {
		return nil, err
	}
	return pb.NewUserHandlerProtobufClient(usersUrl.String(), http.DefaultClient), nil
}
