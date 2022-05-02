// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: protofiles/shakespeare.proto

package protofiles

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

// QuoteTransactionClient is the client API for QuoteTransaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuoteTransactionClient interface {
	MakeQuote(ctx context.Context, in *QuoteRequest, opts ...grpc.CallOption) (*QuoteResponse, error)
}

type quoteTransactionClient struct {
	cc grpc.ClientConnInterface
}

func NewQuoteTransactionClient(cc grpc.ClientConnInterface) QuoteTransactionClient {
	return &quoteTransactionClient{cc}
}

func (c *quoteTransactionClient) MakeQuote(ctx context.Context, in *QuoteRequest, opts ...grpc.CallOption) (*QuoteResponse, error) {
	out := new(QuoteResponse)
	err := c.cc.Invoke(ctx, "/protofiles.QuoteTransaction/MakeQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuoteTransactionServer is the server API for QuoteTransaction service.
// All implementations must embed UnimplementedQuoteTransactionServer
// for forward compatibility
type QuoteTransactionServer interface {
	MakeQuote(context.Context, *QuoteRequest) (*QuoteResponse, error)
	mustEmbedUnimplementedQuoteTransactionServer()
}

// UnimplementedQuoteTransactionServer must be embedded to have forward compatible implementations.
type UnimplementedQuoteTransactionServer struct {
}

func (UnimplementedQuoteTransactionServer) MakeQuote(context.Context, *QuoteRequest) (*QuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeQuote not implemented")
}
func (UnimplementedQuoteTransactionServer) mustEmbedUnimplementedQuoteTransactionServer() {}

// UnsafeQuoteTransactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuoteTransactionServer will
// result in compilation errors.
type UnsafeQuoteTransactionServer interface {
	mustEmbedUnimplementedQuoteTransactionServer()
}

func RegisterQuoteTransactionServer(s grpc.ServiceRegistrar, srv QuoteTransactionServer) {
	s.RegisterService(&QuoteTransaction_ServiceDesc, srv)
}

func _QuoteTransaction_MakeQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteTransactionServer).MakeQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.QuoteTransaction/MakeQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteTransactionServer).MakeQuote(ctx, req.(*QuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuoteTransaction_ServiceDesc is the grpc.ServiceDesc for QuoteTransaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuoteTransaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.QuoteTransaction",
	HandlerType: (*QuoteTransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeQuote",
			Handler:    _QuoteTransaction_MakeQuote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofiles/shakespeare.proto",
}
