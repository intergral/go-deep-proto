// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.21.6
// source: deepproto/proto/poll/v1/poll.proto

package v1

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

// PollConfigClient is the client API for PollConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PollConfigClient interface {
	//Call this function as often as is required to ensure config is up to date.
	Poll(ctx context.Context, in *PollRequest, opts ...grpc.CallOption) (*PollResponse, error)
}

type pollConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewPollConfigClient(cc grpc.ClientConnInterface) PollConfigClient {
	return &pollConfigClient{cc}
}

func (c *pollConfigClient) Poll(ctx context.Context, in *PollRequest, opts ...grpc.CallOption) (*PollResponse, error) {
	out := new(PollResponse)
	err := c.cc.Invoke(ctx, "/deepproto.proto.poll.v1.PollConfig/poll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PollConfigServer is the server API for PollConfig service.
// All implementations must embed UnimplementedPollConfigServer
// for forward compatibility
type PollConfigServer interface {
	//Call this function as often as is required to ensure config is up to date.
	Poll(context.Context, *PollRequest) (*PollResponse, error)
	mustEmbedUnimplementedPollConfigServer()
}

// UnimplementedPollConfigServer must be embedded to have forward compatible implementations.
type UnimplementedPollConfigServer struct {
}

func (UnimplementedPollConfigServer) Poll(context.Context, *PollRequest) (*PollResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Poll not implemented")
}
func (UnimplementedPollConfigServer) mustEmbedUnimplementedPollConfigServer() {}

// UnsafePollConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PollConfigServer will
// result in compilation errors.
type UnsafePollConfigServer interface {
	mustEmbedUnimplementedPollConfigServer()
}

func RegisterPollConfigServer(s grpc.ServiceRegistrar, srv PollConfigServer) {
	s.RegisterService(&PollConfig_ServiceDesc, srv)
}

func _PollConfig_Poll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollConfigServer).Poll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deepproto.proto.poll.v1.PollConfig/poll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollConfigServer).Poll(ctx, req.(*PollRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PollConfig_ServiceDesc is the grpc.ServiceDesc for PollConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PollConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "deepproto.proto.poll.v1.PollConfig",
	HandlerType: (*PollConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "poll",
			Handler:    _PollConfig_Poll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deepproto/proto/poll/v1/poll.proto",
}