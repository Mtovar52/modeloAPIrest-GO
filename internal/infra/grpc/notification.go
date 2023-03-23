package grpc

import (
	pb_SendGrid "api-gateway/internal/infra/proto/sendGrid"

	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"api-gateway/cmd/config"
)

var once sync.Once
var err error

var con *grpc.ClientConn

var publisherSendGrid pb_SendGrid.SendGridEmailServiceClient

func NewCon() {
	once.Do(func() {
		con, err = grpc.Dial(config.GetConfig().Domain.GrpcNotification, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic("error al conectarse al grpc")
		}

		publisherSendGrid = pb_SendGrid.NewSendGridEmailServiceClient(con)

	})
}

func GetPublisherSendGrid() pb_SendGrid.SendGridEmailServiceClient {
	return publisherSendGrid
}
