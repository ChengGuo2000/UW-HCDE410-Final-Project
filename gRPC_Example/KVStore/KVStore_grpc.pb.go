// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: KVStore/KVStore.proto

package KVStore

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

const (
	KVStore_Get_FullMethodName = "/KVStore.KVStore/Get"
	KVStore_Put_FullMethodName = "/KVStore.KVStore/Put"
)

// KVStoreClient is the client API for KVStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KVStoreClient interface {
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	Put(ctx context.Context, in *KVPair, opts ...grpc.CallOption) (*Success, error)
}

type kVStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewKVStoreClient(cc grpc.ClientConnInterface) KVStoreClient {
	return &kVStoreClient{cc}
}

func (c *kVStoreClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, KVStore_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVStoreClient) Put(ctx context.Context, in *KVPair, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, KVStore_Put_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVStoreServer is the server API for KVStore service.
// All implementations should embed UnimplementedKVStoreServer
// for forward compatibility
type KVStoreServer interface {
	Get(context.Context, *Key) (*Value, error)
	Put(context.Context, *KVPair) (*Success, error)
}

// UnimplementedKVStoreServer should be embedded to have forward compatible implementations.
type UnimplementedKVStoreServer struct {
}

func (UnimplementedKVStoreServer) Get(context.Context, *Key) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKVStoreServer) Put(context.Context, *KVPair) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}

// UnsafeKVStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KVStoreServer will
// result in compilation errors.
type UnsafeKVStoreServer interface {
	mustEmbedUnimplementedKVStoreServer()
}

func RegisterKVStoreServer(s grpc.ServiceRegistrar, srv KVStoreServer) {
	s.RegisterService(&KVStore_ServiceDesc, srv)
}

func _KVStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KVStore_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStoreServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVStore_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStoreServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KVStore_Put_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStoreServer).Put(ctx, req.(*KVPair))
	}
	return interceptor(ctx, in, info, handler)
}

// KVStore_ServiceDesc is the grpc.ServiceDesc for KVStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KVStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KVStore.KVStore",
	HandlerType: (*KVStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _KVStore_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _KVStore_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "KVStore/KVStore.proto",
}
