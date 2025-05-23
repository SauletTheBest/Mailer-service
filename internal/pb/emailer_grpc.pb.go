// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: emailer.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	EmailerService_SendEmail_FullMethodName = "/pb.EmailerService/SendEmail"
)

// EmailerServiceClient is the client API for EmailerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailerServiceClient interface {
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error)
}

type emailerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailerServiceClient(cc grpc.ClientConnInterface) EmailerServiceClient {
	return &emailerServiceClient{cc}
}

func (c *emailerServiceClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendEmailResponse)
	err := c.cc.Invoke(ctx, EmailerService_SendEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailerServiceServer is the server API for EmailerService service.
// All implementations must embed UnimplementedEmailerServiceServer
// for forward compatibility.
type EmailerServiceServer interface {
	SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error)
	mustEmbedUnimplementedEmailerServiceServer()
}

// UnimplementedEmailerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEmailerServiceServer struct{}

func (UnimplementedEmailerServiceServer) SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedEmailerServiceServer) mustEmbedUnimplementedEmailerServiceServer() {}
func (UnimplementedEmailerServiceServer) testEmbeddedByValue()                        {}

// UnsafeEmailerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailerServiceServer will
// result in compilation errors.
type UnsafeEmailerServiceServer interface {
	mustEmbedUnimplementedEmailerServiceServer()
}

func RegisterEmailerServiceServer(s grpc.ServiceRegistrar, srv EmailerServiceServer) {
	// If the following call pancis, it indicates UnimplementedEmailerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EmailerService_ServiceDesc, srv)
}

func _EmailerService_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailerServiceServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailerService_SendEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailerServiceServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailerService_ServiceDesc is the grpc.ServiceDesc for EmailerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EmailerService",
	HandlerType: (*EmailerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _EmailerService_SendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emailer.proto",
}
