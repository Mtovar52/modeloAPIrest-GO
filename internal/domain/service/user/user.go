package service_user

import (
	validate "api-gateway/internal/domain/validation"
	pbUser "api-gateway/internal/infra/proto/user"
	"context"
	"fmt"

	"api-gateway/internal/infra/grpc"
)

type ServiceUser struct {
	grpc pbUser.UserServiceClient
}

func NewServiceUser() ServiceUser {

	return ServiceUser{
		grpc: grpc.GetPublisherUser(),
	}
}

func (s *ServiceUser) CreateUser(user *pbUser.User) *pbUser.UserResponse {

	paswords := validate.ValidatePaswords(user)
	if paswords != nil {
		return paswords
	}

	crypt := validate.ValidateNamePasswCREATE(user)
	if crypt != nil {
		return crypt
	}

	conditions := validate.ValidateConditions(user)

	if conditions != nil {
		return conditions
	}

	response, err := s.grpc.CreateUser(context.Background(), user) //respuesta del GRPC
	if err != nil {
		fmt.Println(err)
	}

	return response
}

func (s *ServiceUser) UpdateUser(userRequest *pbUser.UpdateRequestUser) *pbUser.UserResponse {
	crypt := validate.ValidatePasswUserUPDATE(userRequest.User)
	if crypt != nil {
		return crypt
	}

	response, err := s.grpc.UpdateUser(context.Background(), userRequest)
	if err != nil {
		fmt.Println(err)
	}

	return response
}

func (s *ServiceUser) ListUser(offset int32) *pbUser.ListAllResponse {
	response, err := s.grpc.ListUser(context.Background(), &pbUser.ListRequestUser{Offset: offset})
	if err != nil {
		fmt.Println(err)
	}
	return response
}

func (s *ServiceUser) DeleteUser(ID int64) *pbUser.UserResponse {
	response, err := s.grpc.DeleteUser(context.Background(), &pbUser.DeleteRequestUser{Id: int64(ID)})
	if err != nil {
		fmt.Println("FATAL ERROR XC", err)
	}
	return response
}

func (s *ServiceUser) GetByIdUser(ID *pbUser.GetById) *pbUser.UserLogin {

	response, err := s.grpc.GetByIdUser(context.Background(), ID)

	if err != nil {
		fmt.Println("err response getById ", err)
	}
	return response
}

func (s *ServiceUser) Authservice(ctx context.Context, in *pbUser.FindVerifRequest) *pbUser.UserLogin {
	response, err := s.grpc.FindUserByEmailAndNick(ctx, in)
	fmt.Println(err)
	return response

}
