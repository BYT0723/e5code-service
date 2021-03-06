// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: ci.proto

package pb

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

// CiClient is the client API for Ci service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CiClient interface {
	GetBuildPlan(ctx context.Context, in *GetBuildPlanReq, opts ...grpc.CallOption) (*GetBuildPlanRsp, error)
	AddBuildPlan(ctx context.Context, in *AddBuildPlanReq, opts ...grpc.CallOption) (*AddBuildPlanRsp, error)
	UpdateBuildPlan(ctx context.Context, in *UpdateBuildPlanReq, opts ...grpc.CallOption) (*UpdateBuildPlanRsp, error)
	DeleteBuildPlan(ctx context.Context, in *DeleteBuildPlanReq, opts ...grpc.CallOption) (*DeleteBuildPlanRsp, error)
	ListBuildPlan(ctx context.Context, in *ListBuildPlanReq, opts ...grpc.CallOption) (*ListBuildPlanRsp, error)
	BuildImage(ctx context.Context, in *BuildReq, opts ...grpc.CallOption) (Ci_BuildImageClient, error)
	GetImage(ctx context.Context, in *GetImageReq, opts ...grpc.CallOption) (*GetImageRsp, error)
	ListImage(ctx context.Context, in *ListImageReq, opts ...grpc.CallOption) (*ListImageRsp, error)
	DeleteImage(ctx context.Context, in *DeleteImageReq, opts ...grpc.CallOption) (*DeleteImageRsp, error)
}

type ciClient struct {
	cc grpc.ClientConnInterface
}

func NewCiClient(cc grpc.ClientConnInterface) CiClient {
	return &ciClient{cc}
}

func (c *ciClient) GetBuildPlan(ctx context.Context, in *GetBuildPlanReq, opts ...grpc.CallOption) (*GetBuildPlanRsp, error) {
	out := new(GetBuildPlanRsp)
	err := c.cc.Invoke(ctx, "/ci.ci/GetBuildPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ciClient) AddBuildPlan(ctx context.Context, in *AddBuildPlanReq, opts ...grpc.CallOption) (*AddBuildPlanRsp, error) {
	out := new(AddBuildPlanRsp)
	err := c.cc.Invoke(ctx, "/ci.ci/AddBuildPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ciClient) UpdateBuildPlan(ctx context.Context, in *UpdateBuildPlanReq, opts ...grpc.CallOption) (*UpdateBuildPlanRsp, error) {
	out := new(UpdateBuildPlanRsp)
	err := c.cc.Invoke(ctx, "/ci.ci/UpdateBuildPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ciClient) DeleteBuildPlan(ctx context.Context, in *DeleteBuildPlanReq, opts ...grpc.CallOption) (*DeleteBuildPlanRsp, error) {
	out := new(DeleteBuildPlanRsp)
	err := c.cc.Invoke(ctx, "/ci.ci/DeleteBuildPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ciClient) ListBuildPlan(ctx context.Context, in *ListBuildPlanReq, opts ...grpc.CallOption) (*ListBuildPlanRsp, error) {
	out := new(ListBuildPlanRsp)
	err := c.cc.Invoke(ctx, "/ci.ci/ListBuildPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ciClient) BuildImage(ctx context.Context, in *BuildReq, opts ...grpc.CallOption) (Ci_BuildImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Ci_ServiceDesc.Streams[0], "/ci.ci/BuildImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &ciBuildImageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ci_BuildImageClient interface {
	Recv() (*BuildRsp, error)
	grpc.ClientStream
}

type ciBuildImageClient struct {
	grpc.ClientStream
}

func (x *ciBuildImageClient) Recv() (*BuildRsp, error) {
	m := new(BuildRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ciClient) GetImage(ctx context.Context, in *GetImageReq, opts ...grpc.CallOption) (*GetImageRsp, error) {
	out := new(GetImageRsp)
	err := c.cc.Invoke(ctx, "/ci.ci/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ciClient) ListImage(ctx context.Context, in *ListImageReq, opts ...grpc.CallOption) (*ListImageRsp, error) {
	out := new(ListImageRsp)
	err := c.cc.Invoke(ctx, "/ci.ci/ListImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ciClient) DeleteImage(ctx context.Context, in *DeleteImageReq, opts ...grpc.CallOption) (*DeleteImageRsp, error) {
	out := new(DeleteImageRsp)
	err := c.cc.Invoke(ctx, "/ci.ci/DeleteImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CiServer is the server API for Ci service.
// All implementations must embed UnimplementedCiServer
// for forward compatibility
type CiServer interface {
	GetBuildPlan(context.Context, *GetBuildPlanReq) (*GetBuildPlanRsp, error)
	AddBuildPlan(context.Context, *AddBuildPlanReq) (*AddBuildPlanRsp, error)
	UpdateBuildPlan(context.Context, *UpdateBuildPlanReq) (*UpdateBuildPlanRsp, error)
	DeleteBuildPlan(context.Context, *DeleteBuildPlanReq) (*DeleteBuildPlanRsp, error)
	ListBuildPlan(context.Context, *ListBuildPlanReq) (*ListBuildPlanRsp, error)
	BuildImage(*BuildReq, Ci_BuildImageServer) error
	GetImage(context.Context, *GetImageReq) (*GetImageRsp, error)
	ListImage(context.Context, *ListImageReq) (*ListImageRsp, error)
	DeleteImage(context.Context, *DeleteImageReq) (*DeleteImageRsp, error)
	mustEmbedUnimplementedCiServer()
}

// UnimplementedCiServer must be embedded to have forward compatible implementations.
type UnimplementedCiServer struct {
}

func (UnimplementedCiServer) GetBuildPlan(context.Context, *GetBuildPlanReq) (*GetBuildPlanRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuildPlan not implemented")
}
func (UnimplementedCiServer) AddBuildPlan(context.Context, *AddBuildPlanReq) (*AddBuildPlanRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBuildPlan not implemented")
}
func (UnimplementedCiServer) UpdateBuildPlan(context.Context, *UpdateBuildPlanReq) (*UpdateBuildPlanRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBuildPlan not implemented")
}
func (UnimplementedCiServer) DeleteBuildPlan(context.Context, *DeleteBuildPlanReq) (*DeleteBuildPlanRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBuildPlan not implemented")
}
func (UnimplementedCiServer) ListBuildPlan(context.Context, *ListBuildPlanReq) (*ListBuildPlanRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuildPlan not implemented")
}
func (UnimplementedCiServer) BuildImage(*BuildReq, Ci_BuildImageServer) error {
	return status.Errorf(codes.Unimplemented, "method BuildImage not implemented")
}
func (UnimplementedCiServer) GetImage(context.Context, *GetImageReq) (*GetImageRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedCiServer) ListImage(context.Context, *ListImageReq) (*ListImageRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImage not implemented")
}
func (UnimplementedCiServer) DeleteImage(context.Context, *DeleteImageReq) (*DeleteImageRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImage not implemented")
}
func (UnimplementedCiServer) mustEmbedUnimplementedCiServer() {}

// UnsafeCiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CiServer will
// result in compilation errors.
type UnsafeCiServer interface {
	mustEmbedUnimplementedCiServer()
}

func RegisterCiServer(s grpc.ServiceRegistrar, srv CiServer) {
	s.RegisterService(&Ci_ServiceDesc, srv)
}

func _Ci_GetBuildPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBuildPlanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CiServer).GetBuildPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.ci/GetBuildPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CiServer).GetBuildPlan(ctx, req.(*GetBuildPlanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ci_AddBuildPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBuildPlanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CiServer).AddBuildPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.ci/AddBuildPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CiServer).AddBuildPlan(ctx, req.(*AddBuildPlanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ci_UpdateBuildPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBuildPlanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CiServer).UpdateBuildPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.ci/UpdateBuildPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CiServer).UpdateBuildPlan(ctx, req.(*UpdateBuildPlanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ci_DeleteBuildPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBuildPlanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CiServer).DeleteBuildPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.ci/DeleteBuildPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CiServer).DeleteBuildPlan(ctx, req.(*DeleteBuildPlanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ci_ListBuildPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBuildPlanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CiServer).ListBuildPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.ci/ListBuildPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CiServer).ListBuildPlan(ctx, req.(*ListBuildPlanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ci_BuildImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BuildReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CiServer).BuildImage(m, &ciBuildImageServer{stream})
}

type Ci_BuildImageServer interface {
	Send(*BuildRsp) error
	grpc.ServerStream
}

type ciBuildImageServer struct {
	grpc.ServerStream
}

func (x *ciBuildImageServer) Send(m *BuildRsp) error {
	return x.ServerStream.SendMsg(m)
}

func _Ci_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CiServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.ci/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CiServer).GetImage(ctx, req.(*GetImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ci_ListImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CiServer).ListImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.ci/ListImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CiServer).ListImage(ctx, req.(*ListImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ci_DeleteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CiServer).DeleteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.ci/DeleteImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CiServer).DeleteImage(ctx, req.(*DeleteImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Ci_ServiceDesc is the grpc.ServiceDesc for Ci service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ci_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ci.ci",
	HandlerType: (*CiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBuildPlan",
			Handler:    _Ci_GetBuildPlan_Handler,
		},
		{
			MethodName: "AddBuildPlan",
			Handler:    _Ci_AddBuildPlan_Handler,
		},
		{
			MethodName: "UpdateBuildPlan",
			Handler:    _Ci_UpdateBuildPlan_Handler,
		},
		{
			MethodName: "DeleteBuildPlan",
			Handler:    _Ci_DeleteBuildPlan_Handler,
		},
		{
			MethodName: "ListBuildPlan",
			Handler:    _Ci_ListBuildPlan_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _Ci_GetImage_Handler,
		},
		{
			MethodName: "ListImage",
			Handler:    _Ci_ListImage_Handler,
		},
		{
			MethodName: "DeleteImage",
			Handler:    _Ci_DeleteImage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BuildImage",
			Handler:       _Ci_BuildImage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ci.proto",
}
