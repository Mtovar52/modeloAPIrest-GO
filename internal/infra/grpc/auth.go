package grpc

import (
	pb_user "api-gateway/internal/infra/proto/user"

	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"api-gateway/cmd/config"
)

var once2 sync.Once

var err2 error

var con2 *grpc.ClientConn

var publisherUser pb_user.UserServiceClient

func NewCon_Auth() {
	once2.Do(func() {
		con2, err2 = grpc.Dial(config.GetConfig().Domain.GrpcAuth, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err2 != nil {
			panic("error al conectarse al grpc")
		}

		publisherUser = pb_user.NewUserServiceClient(con2)
	})
}

func GetPublisherUser() pb_user.UserServiceClient {
	return publisherUser
}
