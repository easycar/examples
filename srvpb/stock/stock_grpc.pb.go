// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: stock.proto

package stock

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

// StockClient is the client API for Stock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockClient interface {
	TryDeduct(ctx context.Context, in *Req, opts ...grpc.CallOption) (*TryDeductResp, error)
	ConfirmDeduct(ctx context.Context, in *Req, opts ...grpc.CallOption) (*ConfirmDeductResp, error)
	CancelDeduct(ctx context.Context, in *Req, opts ...grpc.CallOption) (*CancelDeductResp, error)
}

type stockClient struct {
	cc grpc.ClientConnInterface
}

func NewStockClient(cc grpc.ClientConnInterface) StockClient {
	return &stockClient{cc}
}

func (c *stockClient) TryDeduct(ctx context.Context, in *Req, opts ...grpc.CallOption) (*TryDeductResp, error) {
	out := new(TryDeductResp)
	err := c.cc.Invoke(ctx, "/stock.Stock/TryDeduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) ConfirmDeduct(ctx context.Context, in *Req, opts ...grpc.CallOption) (*ConfirmDeductResp, error) {
	out := new(ConfirmDeductResp)
	err := c.cc.Invoke(ctx, "/stock.Stock/ConfirmDeduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) CancelDeduct(ctx context.Context, in *Req, opts ...grpc.CallOption) (*CancelDeductResp, error) {
	out := new(CancelDeductResp)
	err := c.cc.Invoke(ctx, "/stock.Stock/CancelDeduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServer is the server API for Stock service.
// All implementations must embed UnimplementedStockServer
// for forward compatibility
type StockServer interface {
	TryDeduct(context.Context, *Req) (*TryDeductResp, error)
	ConfirmDeduct(context.Context, *Req) (*ConfirmDeductResp, error)
	CancelDeduct(context.Context, *Req) (*CancelDeductResp, error)
	mustEmbedUnimplementedStockServer()
}

// UnimplementedStockServer must be embedded to have forward compatible implementations.
type UnimplementedStockServer struct {
}

func (UnimplementedStockServer) TryDeduct(context.Context, *Req) (*TryDeductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TryDeduct not implemented")
}
func (UnimplementedStockServer) ConfirmDeduct(context.Context, *Req) (*ConfirmDeductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmDeduct not implemented")
}
func (UnimplementedStockServer) CancelDeduct(context.Context, *Req) (*CancelDeductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelDeduct not implemented")
}
func (UnimplementedStockServer) mustEmbedUnimplementedStockServer() {}

// UnsafeStockServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServer will
// result in compilation errors.
type UnsafeStockServer interface {
	mustEmbedUnimplementedStockServer()
}

func RegisterStockServer(s grpc.ServiceRegistrar, srv StockServer) {
	s.RegisterService(&Stock_ServiceDesc, srv)
}

func _Stock_TryDeduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).TryDeduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.Stock/TryDeduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).TryDeduct(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_ConfirmDeduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).ConfirmDeduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.Stock/ConfirmDeduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).ConfirmDeduct(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_CancelDeduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).CancelDeduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.Stock/CancelDeduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).CancelDeduct(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// Stock_ServiceDesc is the grpc.ServiceDesc for Stock service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stock_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stock.Stock",
	HandlerType: (*StockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TryDeduct",
			Handler:    _Stock_TryDeduct_Handler,
		},
		{
			MethodName: "ConfirmDeduct",
			Handler:    _Stock_ConfirmDeduct_Handler,
		},
		{
			MethodName: "CancelDeduct",
			Handler:    _Stock_CancelDeduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stock.proto",
}
