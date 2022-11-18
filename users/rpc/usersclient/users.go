// Code generated by goctl. DO NOT EDIT!
// Source: users.proto

package usersclient

import (
	"context"

	"gcloud/users/rpc/users"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginReply                  = users.LoginReply
	LoginRequest                = users.LoginRequest
	RefreshAuthorizationReply   = users.RefreshAuthorizationReply
	RefreshAuthorizationRequest = users.RefreshAuthorizationRequest
	RegisterCountReply          = users.RegisterCountReply
	RegisterCountRequest        = users.RegisterCountRequest
	UserDetailReply             = users.UserDetailReply
	UserDetailRequest           = users.UserDetailRequest
	UserRegisterReply           = users.UserRegisterReply
	UserRegisterRequest         = users.UserRegisterRequest
	UserUpdateReply             = users.UserUpdateReply
	UserUpdateRequest           = users.UserUpdateRequest

	Users interface {
		UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateReply, error)
		RefreshAuthorization(ctx context.Context, in *RefreshAuthorizationRequest, opts ...grpc.CallOption) (*RefreshAuthorizationReply, error)
		RegisterCount(ctx context.Context, in *RegisterCountRequest, opts ...grpc.CallOption) (*RegisterCountReply, error)
		UserDetail(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetailReply, error)
		UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
		UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterReply, error)
	}

	defaultUsers struct {
		cli zrpc.Client
	}
)

func NewUsers(cli zrpc.Client) Users {
	return &defaultUsers{
		cli: cli,
	}
}

func (m *defaultUsers) UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateReply, error) {
	client := users.NewUsersClient(m.cli.Conn())
	return client.UserUpdate(ctx, in, opts...)
}

func (m *defaultUsers) RefreshAuthorization(ctx context.Context, in *RefreshAuthorizationRequest, opts ...grpc.CallOption) (*RefreshAuthorizationReply, error) {
	client := users.NewUsersClient(m.cli.Conn())
	return client.RefreshAuthorization(ctx, in, opts...)
}

func (m *defaultUsers) RegisterCount(ctx context.Context, in *RegisterCountRequest, opts ...grpc.CallOption) (*RegisterCountReply, error) {
	client := users.NewUsersClient(m.cli.Conn())
	return client.RegisterCount(ctx, in, opts...)
}

func (m *defaultUsers) UserDetail(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetailReply, error) {
	client := users.NewUsersClient(m.cli.Conn())
	return client.UserDetail(ctx, in, opts...)
}

func (m *defaultUsers) UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	client := users.NewUsersClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

func (m *defaultUsers) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterReply, error) {
	client := users.NewUsersClient(m.cli.Conn())
	return client.UserRegister(ctx, in, opts...)
}
