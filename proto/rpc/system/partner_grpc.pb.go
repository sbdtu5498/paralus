// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpcv3

import (
	context "context"
	v3 "github.com/RafayLabs/rcloud-base/proto/types/systempb/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PartnerClient is the client API for Partner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartnerClient interface {
	CreatePartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error)
	GetPartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error)
	GetInitPartner(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*v3.Partner, error)
	UpdatePartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error)
	DeletePartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error)
}

type partnerClient struct {
	cc grpc.ClientConnInterface
}

func NewPartnerClient(cc grpc.ClientConnInterface) PartnerClient {
	return &partnerClient{cc}
}

func (c *partnerClient) CreatePartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error) {
	out := new(v3.Partner)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.Partner/CreatePartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerClient) GetPartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error) {
	out := new(v3.Partner)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.Partner/GetPartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerClient) GetInitPartner(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*v3.Partner, error) {
	out := new(v3.Partner)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.Partner/GetInitPartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerClient) UpdatePartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error) {
	out := new(v3.Partner)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.Partner/UpdatePartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerClient) DeletePartner(ctx context.Context, in *v3.Partner, opts ...grpc.CallOption) (*v3.Partner, error) {
	out := new(v3.Partner)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.Partner/DeletePartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartnerServer is the server API for Partner service.
// All implementations should embed UnimplementedPartnerServer
// for forward compatibility
type PartnerServer interface {
	CreatePartner(context.Context, *v3.Partner) (*v3.Partner, error)
	GetPartner(context.Context, *v3.Partner) (*v3.Partner, error)
	GetInitPartner(context.Context, *EmptyRequest) (*v3.Partner, error)
	UpdatePartner(context.Context, *v3.Partner) (*v3.Partner, error)
	DeletePartner(context.Context, *v3.Partner) (*v3.Partner, error)
}

// UnimplementedPartnerServer should be embedded to have forward compatible implementations.
type UnimplementedPartnerServer struct {
}

func (UnimplementedPartnerServer) CreatePartner(context.Context, *v3.Partner) (*v3.Partner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePartner not implemented")
}
func (UnimplementedPartnerServer) GetPartner(context.Context, *v3.Partner) (*v3.Partner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartner not implemented")
}
func (UnimplementedPartnerServer) GetInitPartner(context.Context, *EmptyRequest) (*v3.Partner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInitPartner not implemented")
}
func (UnimplementedPartnerServer) UpdatePartner(context.Context, *v3.Partner) (*v3.Partner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePartner not implemented")
}
func (UnimplementedPartnerServer) DeletePartner(context.Context, *v3.Partner) (*v3.Partner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePartner not implemented")
}

// UnsafePartnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartnerServer will
// result in compilation errors.
type UnsafePartnerServer interface {
	mustEmbedUnimplementedPartnerServer()
}

func RegisterPartnerServer(s grpc.ServiceRegistrar, srv PartnerServer) {
	s.RegisterService(&Partner_ServiceDesc, srv)
}

func _Partner_CreatePartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Partner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServer).CreatePartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.Partner/CreatePartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServer).CreatePartner(ctx, req.(*v3.Partner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Partner_GetPartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Partner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServer).GetPartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.Partner/GetPartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServer).GetPartner(ctx, req.(*v3.Partner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Partner_GetInitPartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServer).GetInitPartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.Partner/GetInitPartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServer).GetInitPartner(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Partner_UpdatePartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Partner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServer).UpdatePartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.Partner/UpdatePartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServer).UpdatePartner(ctx, req.(*v3.Partner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Partner_DeletePartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Partner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServer).DeletePartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.Partner/DeletePartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServer).DeletePartner(ctx, req.(*v3.Partner))
	}
	return interceptor(ctx, in, info, handler)
}

// Partner_ServiceDesc is the grpc.ServiceDesc for Partner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Partner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rafay.dev.rpc.v3.Partner",
	HandlerType: (*PartnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePartner",
			Handler:    _Partner_CreatePartner_Handler,
		},
		{
			MethodName: "GetPartner",
			Handler:    _Partner_GetPartner_Handler,
		},
		{
			MethodName: "GetInitPartner",
			Handler:    _Partner_GetInitPartner_Handler,
		},
		{
			MethodName: "UpdatePartner",
			Handler:    _Partner_UpdatePartner_Handler,
		},
		{
			MethodName: "DeletePartner",
			Handler:    _Partner_DeletePartner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/system/partner.proto",
}
