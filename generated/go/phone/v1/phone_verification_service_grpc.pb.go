// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: phone/v1/phone_verification_service.proto

package phone

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
	PhoneVerification_SendVerificationCode_FullMethodName     = "/code.phone.v1.PhoneVerification/SendVerificationCode"
	PhoneVerification_CheckVerificationCode_FullMethodName    = "/code.phone.v1.PhoneVerification/CheckVerificationCode"
	PhoneVerification_GetAssociatedPhoneNumber_FullMethodName = "/code.phone.v1.PhoneVerification/GetAssociatedPhoneNumber"
)

// PhoneVerificationClient is the client API for PhoneVerification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhoneVerificationClient interface {
	// SendVerificationCode sends a verification code to the provided phone number
	// over SMS. If an active verification is already taking place, the existing code
	// will be resent.
	SendVerificationCode(ctx context.Context, in *SendVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error)
	// CheckVerificationCode validates a verification code. On success, a one-time use
	// token to link an owner account is provided.
	CheckVerificationCode(ctx context.Context, in *CheckVerificationCodeRequest, opts ...grpc.CallOption) (*CheckVerificationCodeResponse, error)
	// GetAssociatedPhoneNumber gets the latest verified phone number linked to an owner account.
	GetAssociatedPhoneNumber(ctx context.Context, in *GetAssociatedPhoneNumberRequest, opts ...grpc.CallOption) (*GetAssociatedPhoneNumberResponse, error)
}

type phoneVerificationClient struct {
	cc grpc.ClientConnInterface
}

func NewPhoneVerificationClient(cc grpc.ClientConnInterface) PhoneVerificationClient {
	return &phoneVerificationClient{cc}
}

func (c *phoneVerificationClient) SendVerificationCode(ctx context.Context, in *SendVerificationCodeRequest, opts ...grpc.CallOption) (*SendVerificationCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendVerificationCodeResponse)
	err := c.cc.Invoke(ctx, PhoneVerification_SendVerificationCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phoneVerificationClient) CheckVerificationCode(ctx context.Context, in *CheckVerificationCodeRequest, opts ...grpc.CallOption) (*CheckVerificationCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckVerificationCodeResponse)
	err := c.cc.Invoke(ctx, PhoneVerification_CheckVerificationCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phoneVerificationClient) GetAssociatedPhoneNumber(ctx context.Context, in *GetAssociatedPhoneNumberRequest, opts ...grpc.CallOption) (*GetAssociatedPhoneNumberResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAssociatedPhoneNumberResponse)
	err := c.cc.Invoke(ctx, PhoneVerification_GetAssociatedPhoneNumber_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhoneVerificationServer is the server API for PhoneVerification service.
// All implementations must embed UnimplementedPhoneVerificationServer
// for forward compatibility.
type PhoneVerificationServer interface {
	// SendVerificationCode sends a verification code to the provided phone number
	// over SMS. If an active verification is already taking place, the existing code
	// will be resent.
	SendVerificationCode(context.Context, *SendVerificationCodeRequest) (*SendVerificationCodeResponse, error)
	// CheckVerificationCode validates a verification code. On success, a one-time use
	// token to link an owner account is provided.
	CheckVerificationCode(context.Context, *CheckVerificationCodeRequest) (*CheckVerificationCodeResponse, error)
	// GetAssociatedPhoneNumber gets the latest verified phone number linked to an owner account.
	GetAssociatedPhoneNumber(context.Context, *GetAssociatedPhoneNumberRequest) (*GetAssociatedPhoneNumberResponse, error)
	mustEmbedUnimplementedPhoneVerificationServer()
}

// UnimplementedPhoneVerificationServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPhoneVerificationServer struct{}

func (UnimplementedPhoneVerificationServer) SendVerificationCode(context.Context, *SendVerificationCodeRequest) (*SendVerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerificationCode not implemented")
}
func (UnimplementedPhoneVerificationServer) CheckVerificationCode(context.Context, *CheckVerificationCodeRequest) (*CheckVerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckVerificationCode not implemented")
}
func (UnimplementedPhoneVerificationServer) GetAssociatedPhoneNumber(context.Context, *GetAssociatedPhoneNumberRequest) (*GetAssociatedPhoneNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssociatedPhoneNumber not implemented")
}
func (UnimplementedPhoneVerificationServer) mustEmbedUnimplementedPhoneVerificationServer() {}
func (UnimplementedPhoneVerificationServer) testEmbeddedByValue()                           {}

// UnsafePhoneVerificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhoneVerificationServer will
// result in compilation errors.
type UnsafePhoneVerificationServer interface {
	mustEmbedUnimplementedPhoneVerificationServer()
}

func RegisterPhoneVerificationServer(s grpc.ServiceRegistrar, srv PhoneVerificationServer) {
	// If the following call pancis, it indicates UnimplementedPhoneVerificationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PhoneVerification_ServiceDesc, srv)
}

func _PhoneVerification_SendVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVerificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneVerificationServer).SendVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhoneVerification_SendVerificationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneVerificationServer).SendVerificationCode(ctx, req.(*SendVerificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhoneVerification_CheckVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckVerificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneVerificationServer).CheckVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhoneVerification_CheckVerificationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneVerificationServer).CheckVerificationCode(ctx, req.(*CheckVerificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhoneVerification_GetAssociatedPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssociatedPhoneNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneVerificationServer).GetAssociatedPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhoneVerification_GetAssociatedPhoneNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneVerificationServer).GetAssociatedPhoneNumber(ctx, req.(*GetAssociatedPhoneNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PhoneVerification_ServiceDesc is the grpc.ServiceDesc for PhoneVerification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhoneVerification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "code.phone.v1.PhoneVerification",
	HandlerType: (*PhoneVerificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendVerificationCode",
			Handler:    _PhoneVerification_SendVerificationCode_Handler,
		},
		{
			MethodName: "CheckVerificationCode",
			Handler:    _PhoneVerification_CheckVerificationCode_Handler,
		},
		{
			MethodName: "GetAssociatedPhoneNumber",
			Handler:    _PhoneVerification_GetAssociatedPhoneNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "phone/v1/phone_verification_service.proto",
}
