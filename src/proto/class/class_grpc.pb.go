// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: src/proto/class/class.proto

package classpb

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

// ClassServiceClient is the client API for ClassService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClassServiceClient interface {
	GetDetailClass(ctx context.Context, in *GetDetailClassRequest, opts ...grpc.CallOption) (*GetDetailClassResponse, error)
}

type classServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClassServiceClient(cc grpc.ClientConnInterface) ClassServiceClient {
	return &classServiceClient{cc}
}

func (c *classServiceClient) GetDetailClass(ctx context.Context, in *GetDetailClassRequest, opts ...grpc.CallOption) (*GetDetailClassResponse, error) {
	out := new(GetDetailClassResponse)
	err := c.cc.Invoke(ctx, "/class.ClassService/GetDetailClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClassServiceServer is the server API for ClassService service.
// All implementations must embed UnimplementedClassServiceServer
// for forward compatibility
type ClassServiceServer interface {
	GetDetailClass(context.Context, *GetDetailClassRequest) (*GetDetailClassResponse, error)
	mustEmbedUnimplementedClassServiceServer()
}

// UnimplementedClassServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClassServiceServer struct {
}

func (UnimplementedClassServiceServer) GetDetailClass(context.Context, *GetDetailClassRequest) (*GetDetailClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailClass not implemented")
}
func (UnimplementedClassServiceServer) mustEmbedUnimplementedClassServiceServer() {}

// UnsafeClassServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClassServiceServer will
// result in compilation errors.
type UnsafeClassServiceServer interface {
	mustEmbedUnimplementedClassServiceServer()
}

func RegisterClassServiceServer(s grpc.ServiceRegistrar, srv ClassServiceServer) {
	s.RegisterService(&ClassService_ServiceDesc, srv)
}

func _ClassService_GetDetailClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).GetDetailClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/class.ClassService/GetDetailClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).GetDetailClass(ctx, req.(*GetDetailClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClassService_ServiceDesc is the grpc.ServiceDesc for ClassService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClassService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "class.ClassService",
	HandlerType: (*ClassServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDetailClass",
			Handler:    _ClassService_GetDetailClass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/proto/class/class.proto",
}
