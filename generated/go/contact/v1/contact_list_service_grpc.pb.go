// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: contact/v1/contact_list_service.proto

package contact

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

// ContactListClient is the client API for ContactList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContactListClient interface {
	// AddContacts adds a batch of contacts to a user's contact list
	AddContacts(ctx context.Context, in *AddContactsRequest, opts ...grpc.CallOption) (*AddContactsResponse, error)
	// RemoveContacts removes a batch of contacts from a user's contact list
	RemoveContacts(ctx context.Context, in *RemoveContactsRequest, opts ...grpc.CallOption) (*RemoveContactsResponse, error)
	// GetContacts gets a subset of contacts from a user's contact list
	GetContacts(ctx context.Context, in *GetContactsRequest, opts ...grpc.CallOption) (*GetContactsResponse, error)
}

type contactListClient struct {
	cc grpc.ClientConnInterface
}

func NewContactListClient(cc grpc.ClientConnInterface) ContactListClient {
	return &contactListClient{cc}
}

func (c *contactListClient) AddContacts(ctx context.Context, in *AddContactsRequest, opts ...grpc.CallOption) (*AddContactsResponse, error) {
	out := new(AddContactsResponse)
	err := c.cc.Invoke(ctx, "/code.contact.v1.ContactList/AddContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactListClient) RemoveContacts(ctx context.Context, in *RemoveContactsRequest, opts ...grpc.CallOption) (*RemoveContactsResponse, error) {
	out := new(RemoveContactsResponse)
	err := c.cc.Invoke(ctx, "/code.contact.v1.ContactList/RemoveContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactListClient) GetContacts(ctx context.Context, in *GetContactsRequest, opts ...grpc.CallOption) (*GetContactsResponse, error) {
	out := new(GetContactsResponse)
	err := c.cc.Invoke(ctx, "/code.contact.v1.ContactList/GetContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactListServer is the server API for ContactList service.
// All implementations must embed UnimplementedContactListServer
// for forward compatibility
type ContactListServer interface {
	// AddContacts adds a batch of contacts to a user's contact list
	AddContacts(context.Context, *AddContactsRequest) (*AddContactsResponse, error)
	// RemoveContacts removes a batch of contacts from a user's contact list
	RemoveContacts(context.Context, *RemoveContactsRequest) (*RemoveContactsResponse, error)
	// GetContacts gets a subset of contacts from a user's contact list
	GetContacts(context.Context, *GetContactsRequest) (*GetContactsResponse, error)
	mustEmbedUnimplementedContactListServer()
}

// UnimplementedContactListServer must be embedded to have forward compatible implementations.
type UnimplementedContactListServer struct {
}

func (UnimplementedContactListServer) AddContacts(context.Context, *AddContactsRequest) (*AddContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddContacts not implemented")
}
func (UnimplementedContactListServer) RemoveContacts(context.Context, *RemoveContactsRequest) (*RemoveContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveContacts not implemented")
}
func (UnimplementedContactListServer) GetContacts(context.Context, *GetContactsRequest) (*GetContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContacts not implemented")
}
func (UnimplementedContactListServer) mustEmbedUnimplementedContactListServer() {}

// UnsafeContactListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContactListServer will
// result in compilation errors.
type UnsafeContactListServer interface {
	mustEmbedUnimplementedContactListServer()
}

func RegisterContactListServer(s grpc.ServiceRegistrar, srv ContactListServer) {
	s.RegisterService(&ContactList_ServiceDesc, srv)
}

func _ContactList_AddContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactListServer).AddContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.contact.v1.ContactList/AddContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactListServer).AddContacts(ctx, req.(*AddContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactList_RemoveContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactListServer).RemoveContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.contact.v1.ContactList/RemoveContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactListServer).RemoveContacts(ctx, req.(*RemoveContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactList_GetContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactListServer).GetContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.contact.v1.ContactList/GetContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactListServer).GetContacts(ctx, req.(*GetContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContactList_ServiceDesc is the grpc.ServiceDesc for ContactList service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContactList_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "code.contact.v1.ContactList",
	HandlerType: (*ContactListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddContacts",
			Handler:    _ContactList_AddContacts_Handler,
		},
		{
			MethodName: "RemoveContacts",
			Handler:    _ContactList_RemoveContacts_Handler,
		},
		{
			MethodName: "GetContacts",
			Handler:    _ContactList_GetContacts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact/v1/contact_list_service.proto",
}
