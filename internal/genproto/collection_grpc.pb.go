// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: protos/collection.proto

package genproto

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

// CollectionServiceClient is the client API for CollectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectionServiceClient interface {
	// COLLECTION MANAGEMENT
	// ListCollections gets all collections of a user...
	ListCollections(ctx context.Context, in *ListCollectionsRequest, opts ...grpc.CallOption) (CollectionService_ListCollectionsClient, error)
	// CreateCollection creates a collection
	CreateCollection(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*Collection, error)
	// GetCollection get a given collection given a string (the collection's ID)
	GetCollection(ctx context.Context, in *GetCollectionRequest, opts ...grpc.CallOption) (*Collection, error)
}

type collectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectionServiceClient(cc grpc.ClientConnInterface) CollectionServiceClient {
	return &collectionServiceClient{cc}
}

func (c *collectionServiceClient) ListCollections(ctx context.Context, in *ListCollectionsRequest, opts ...grpc.CallOption) (CollectionService_ListCollectionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CollectionService_ServiceDesc.Streams[0], "/collection.CollectionService/ListCollections", opts...)
	if err != nil {
		return nil, err
	}
	x := &collectionServiceListCollectionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CollectionService_ListCollectionsClient interface {
	Recv() (*Collection, error)
	grpc.ClientStream
}

type collectionServiceListCollectionsClient struct {
	grpc.ClientStream
}

func (x *collectionServiceListCollectionsClient) Recv() (*Collection, error) {
	m := new(Collection)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *collectionServiceClient) CreateCollection(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*Collection, error) {
	out := new(Collection)
	err := c.cc.Invoke(ctx, "/collection.CollectionService/CreateCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) GetCollection(ctx context.Context, in *GetCollectionRequest, opts ...grpc.CallOption) (*Collection, error) {
	out := new(Collection)
	err := c.cc.Invoke(ctx, "/collection.CollectionService/GetCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectionServiceServer is the server API for CollectionService service.
// All implementations must embed UnimplementedCollectionServiceServer
// for forward compatibility
type CollectionServiceServer interface {
	// COLLECTION MANAGEMENT
	// ListCollections gets all collections of a user...
	ListCollections(*ListCollectionsRequest, CollectionService_ListCollectionsServer) error
	// CreateCollection creates a collection
	CreateCollection(context.Context, *Collection) (*Collection, error)
	// GetCollection get a given collection given a string (the collection's ID)
	GetCollection(context.Context, *GetCollectionRequest) (*Collection, error)
	mustEmbedUnimplementedCollectionServiceServer()
}

// UnimplementedCollectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCollectionServiceServer struct {
}

func (UnimplementedCollectionServiceServer) ListCollections(*ListCollectionsRequest, CollectionService_ListCollectionsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCollections not implemented")
}
func (UnimplementedCollectionServiceServer) CreateCollection(context.Context, *Collection) (*Collection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (UnimplementedCollectionServiceServer) GetCollection(context.Context, *GetCollectionRequest) (*Collection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollection not implemented")
}
func (UnimplementedCollectionServiceServer) mustEmbedUnimplementedCollectionServiceServer() {}

// UnsafeCollectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectionServiceServer will
// result in compilation errors.
type UnsafeCollectionServiceServer interface {
	mustEmbedUnimplementedCollectionServiceServer()
}

func RegisterCollectionServiceServer(s grpc.ServiceRegistrar, srv CollectionServiceServer) {
	s.RegisterService(&CollectionService_ServiceDesc, srv)
}

func _CollectionService_ListCollections_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListCollectionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CollectionServiceServer).ListCollections(m, &collectionServiceListCollectionsServer{stream})
}

type CollectionService_ListCollectionsServer interface {
	Send(*Collection) error
	grpc.ServerStream
}

type collectionServiceListCollectionsServer struct {
	grpc.ServerStream
}

func (x *collectionServiceListCollectionsServer) Send(m *Collection) error {
	return x.ServerStream.SendMsg(m)
}

func _CollectionService_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Collection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collection.CollectionService/CreateCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).CreateCollection(ctx, req.(*Collection))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_GetCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).GetCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collection.CollectionService/GetCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).GetCollection(ctx, req.(*GetCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CollectionService_ServiceDesc is the grpc.ServiceDesc for CollectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collection.CollectionService",
	HandlerType: (*CollectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCollection",
			Handler:    _CollectionService_CreateCollection_Handler,
		},
		{
			MethodName: "GetCollection",
			Handler:    _CollectionService_GetCollection_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCollections",
			Handler:       _CollectionService_ListCollections_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/collection.proto",
}
