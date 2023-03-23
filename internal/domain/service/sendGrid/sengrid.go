package sendgrid_service

import (
	pbSend "api-gateway/internal/infra/proto/sendGrid"
	"context"
	"fmt"

	"api-gateway/internal/infra/grpc"
)

type ServiceSend struct {
	grpc pbSend.SendGridEmailServiceClient
}

func NewServiceSend() ServiceSend {

	return ServiceSend{
		grpc: grpc.GetPublisherSendGrid(),
	}
}

func (s *ServiceSend) SendEmail(ctx context.Context, send *pbSend.Send) *pbSend.SendResponse {
	response, err := s.grpc.SendEmail(ctx, send)
	if err != nil {
		fmt.Println(err)
	}

	return response
}
