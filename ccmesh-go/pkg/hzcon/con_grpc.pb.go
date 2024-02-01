// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: con.proto

package hzcon

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MeshClient is the client API for Mesh service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeshClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	Prepare1(ctx context.Context, in *Prepare1Request, opts ...grpc.CallOption) (*Prepare1Response, error)
	Prepare2(ctx context.Context, in *Prepare2Request, opts ...grpc.CallOption) (*empty.Empty, error)
	ServerRead(ctx context.Context, in *ServerReadRequest, opts ...grpc.CallOption) (*ServerReadResponse, error)
	CleanUp(ctx context.Context, in *CleanUpRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ClientRead(ctx context.Context, in *ClientReadRequest, opts ...grpc.CallOption) (*ClientReadResponse, error)
	ClientWrite(ctx context.Context, in *ClientWriteRequest, opts ...grpc.CallOption) (*ClientWriteResponse, error)
}

type meshClient struct {
	cc grpc.ClientConnInterface
}

func NewMeshClient(cc grpc.ClientConnInterface) MeshClient {
	return &meshClient{cc}
}

func (c *meshClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/con.Mesh/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) Prepare1(ctx context.Context, in *Prepare1Request, opts ...grpc.CallOption) (*Prepare1Response, error) {
	out := new(Prepare1Response)
	err := c.cc.Invoke(ctx, "/con.Mesh/Prepare1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) Prepare2(ctx context.Context, in *Prepare2Request, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/con.Mesh/Prepare2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) ServerRead(ctx context.Context, in *ServerReadRequest, opts ...grpc.CallOption) (*ServerReadResponse, error) {
	out := new(ServerReadResponse)
	err := c.cc.Invoke(ctx, "/con.Mesh/ServerRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) CleanUp(ctx context.Context, in *CleanUpRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/con.Mesh/CleanUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) ClientRead(ctx context.Context, in *ClientReadRequest, opts ...grpc.CallOption) (*ClientReadResponse, error) {
	out := new(ClientReadResponse)
	err := c.cc.Invoke(ctx, "/con.Mesh/ClientRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) ClientWrite(ctx context.Context, in *ClientWriteRequest, opts ...grpc.CallOption) (*ClientWriteResponse, error) {
	out := new(ClientWriteResponse)
	err := c.cc.Invoke(ctx, "/con.Mesh/ClientWrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeshServer is the server API for Mesh service.
// All implementations must embed UnimplementedMeshServer
// for forward compatibility
type MeshServer interface {
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	Prepare1(context.Context, *Prepare1Request) (*Prepare1Response, error)
	Prepare2(context.Context, *Prepare2Request) (*empty.Empty, error)
	ServerRead(context.Context, *ServerReadRequest) (*ServerReadResponse, error)
	CleanUp(context.Context, *CleanUpRequest) (*empty.Empty, error)
	ClientRead(context.Context, *ClientReadRequest) (*ClientReadResponse, error)
	ClientWrite(context.Context, *ClientWriteRequest) (*ClientWriteResponse, error)
	mustEmbedUnimplementedMeshServer()
}

// UnimplementedMeshServer must be embedded to have forward compatible implementations.
type UnimplementedMeshServer struct {
}

func (UnimplementedMeshServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedMeshServer) Prepare1(context.Context, *Prepare1Request) (*Prepare1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare1 not implemented")
}
func (UnimplementedMeshServer) Prepare2(context.Context, *Prepare2Request) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare2 not implemented")
}
func (UnimplementedMeshServer) ServerRead(context.Context, *ServerReadRequest) (*ServerReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerRead not implemented")
}
func (UnimplementedMeshServer) CleanUp(context.Context, *CleanUpRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanUp not implemented")
}
func (UnimplementedMeshServer) ClientRead(context.Context, *ClientReadRequest) (*ClientReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientRead not implemented")
}
func (UnimplementedMeshServer) ClientWrite(context.Context, *ClientWriteRequest) (*ClientWriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientWrite not implemented")
}
func (UnimplementedMeshServer) mustEmbedUnimplementedMeshServer() {}

// UnsafeMeshServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeshServer will
// result in compilation errors.
type UnsafeMeshServer interface {
	mustEmbedUnimplementedMeshServer()
}

func RegisterMeshServer(s grpc.ServiceRegistrar, srv MeshServer) {
	s.RegisterService(&Mesh_ServiceDesc, srv)
}

func _Mesh_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/con.Mesh/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_Prepare1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prepare1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).Prepare1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/con.Mesh/Prepare1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).Prepare1(ctx, req.(*Prepare1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_Prepare2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prepare2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).Prepare2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/con.Mesh/Prepare2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).Prepare2(ctx, req.(*Prepare2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_ServerRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).ServerRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/con.Mesh/ServerRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).ServerRead(ctx, req.(*ServerReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_CleanUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).CleanUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/con.Mesh/CleanUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).CleanUp(ctx, req.(*CleanUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_ClientRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).ClientRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/con.Mesh/ClientRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).ClientRead(ctx, req.(*ClientReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_ClientWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientWriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).ClientWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/con.Mesh/ClientWrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).ClientWrite(ctx, req.(*ClientWriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mesh_ServiceDesc is the grpc.ServiceDesc for Mesh service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mesh_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "con.Mesh",
	HandlerType: (*MeshServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Mesh_HealthCheck_Handler,
		},
		{
			MethodName: "Prepare1",
			Handler:    _Mesh_Prepare1_Handler,
		},
		{
			MethodName: "Prepare2",
			Handler:    _Mesh_Prepare2_Handler,
		},
		{
			MethodName: "ServerRead",
			Handler:    _Mesh_ServerRead_Handler,
		},
		{
			MethodName: "CleanUp",
			Handler:    _Mesh_CleanUp_Handler,
		},
		{
			MethodName: "ClientRead",
			Handler:    _Mesh_ClientRead_Handler,
		},
		{
			MethodName: "ClientWrite",
			Handler:    _Mesh_ClientWrite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "con.proto",
}
