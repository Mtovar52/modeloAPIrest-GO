// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ploto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SendGridEmailServiceClient is the client API for SendGridEmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendGridEmailServiceClient interface {
	SendEmail(ctx context.Context, in *Send, opts ...grpc.CallOption) (*SendResponse, error)
}

type sendGridEmailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSendGridEmailServiceClient(cc grpc.ClientConnInterface) SendGridEmailServiceClient {
	return &sendGridEmailServiceClient{cc}
}

func (c *sendGridEmailServiceClient) SendEmail(ctx context.Context, in *Send, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/sendgrid.SendGridEmailService/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendGridEmailServiceServer is the server API for SendGridEmailService service.
// All implementations must embed UnimplementedSendGridEmailServiceServer
// for forward compatibility
type SendGridEmailServiceServer interface {
	SendEmail(context.Context, *Send) (*SendResponse, error)
	mustEmbedUnimplementedSendGridEmailServiceServer()
}

// UnimplementedSendGridEmailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSendGridEmailServiceServer struct {
}

func (UnimplementedSendGridEmailServiceServer) SendEmail(context.Context, *Send) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedSendGridEmailServiceServer) mustEmbedUnimplementedSendGridEmailServiceServer() {}

// UnsafeSendGridEmailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendGridEmailServiceServer will
// result in compilation errors.
type UnsafeSendGridEmailServiceServer interface {
	mustEmbedUnimplementedSendGridEmailServiceServer()
}

func RegisterSendGridEmailServiceServer(s grpc.ServiceRegistrar, srv SendGridEmailServiceServer) {
	s.RegisterService(&SendGridEmailService_ServiceDesc, srv)
}

func _SendGridEmailService_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Send)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendGridEmailServiceServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sendgrid.SendGridEmailService/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendGridEmailServiceServer).SendEmail(ctx, req.(*Send))
	}
	return interceptor(ctx, in, info, handler)
}

// SendGridEmailService_ServiceDesc is the grpc.ServiceDesc for SendGridEmailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendGridEmailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sendgrid.SendGridEmailService",
	HandlerType: (*SendGridEmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _SendGridEmailService_SendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ploto/sendGrid.proto",
}
