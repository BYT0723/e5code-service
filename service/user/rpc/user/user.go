// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package user

import (
	"context"

	"e5Code-Service/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddUserReq        = pb.AddUserReq
	AddUserRsp        = pb.AddUserRsp
	DeleteUserReq     = pb.DeleteUserReq
	DeleteUserRsp     = pb.DeleteUserRsp
	GetPermissionReq  = pb.GetPermissionReq
	GetPermissionRsp  = pb.GetPermissionRsp
	GetUserByEmailReq = pb.GetUserByEmailReq
	GetUserReq        = pb.GetUserReq
	GetUserRsp        = pb.GetUserRsp
	LoginReq          = pb.LoginReq
	LoginRsp          = pb.LoginRsp
	SetPermissionReq  = pb.SetPermissionReq
	SetPermissionRsp  = pb.SetPermissionRsp
	UpdateUserReq     = pb.UpdateUserReq
	UpdateUserRsp     = pb.UpdateUserRsp

	User interface {
		GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserRsp, error)
		GetUserByEmail(ctx context.Context, in *GetUserByEmailReq, opts ...grpc.CallOption) (*GetUserRsp, error)
		AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserRsp, error)
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRsp, error)
		DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserRsp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRsp, error)
		GetPermission(ctx context.Context, in *GetPermissionReq, opts ...grpc.CallOption) (*GetPermissionRsp, error)
		SetPermission(ctx context.Context, in *SetPermissionReq, opts ...grpc.CallOption) (*SetPermissionRsp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserRsp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) GetUserByEmail(ctx context.Context, in *GetUserByEmailReq, opts ...grpc.CallOption) (*GetUserRsp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetUserByEmail(ctx, in, opts...)
}

func (m *defaultUser) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserRsp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.AddUser(ctx, in, opts...)
}

func (m *defaultUser) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRsp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUser) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserRsp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRsp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) GetPermission(ctx context.Context, in *GetPermissionReq, opts ...grpc.CallOption) (*GetPermissionRsp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetPermission(ctx, in, opts...)
}

func (m *defaultUser) SetPermission(ctx context.Context, in *SetPermissionReq, opts ...grpc.CallOption) (*SetPermissionRsp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.SetPermission(ctx, in, opts...)
}