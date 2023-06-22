// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package authservice

import (
	"context"

	"github.com/super667/user/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddRoleForUserReq     = user.AddRoleForUserReq
	AddRoleForUserResp    = user.AddRoleForUserResp
	AuthenticateReq       = user.AuthenticateReq
	AuthenticateResp      = user.AuthenticateResp
	BlackListTokenReq     = user.BlackListTokenReq
	BlackListTokenResp    = user.BlackListTokenResp
	CreatePermReq         = user.CreatePermReq
	CreatePermResp        = user.CreatePermResp
	CreateRoleReq         = user.CreateRoleReq
	CreateRoleResp        = user.CreateRoleResp
	CreateStrategyReq     = user.CreateStrategyReq
	CreateStrategyResp    = user.CreateStrategyResp
	CreateUserReq         = user.CreateUserReq
	CreateUserResp        = user.CreateUserResp
	CreateUserRoleReq     = user.CreateUserRoleReq
	CreateUserRoleResp    = user.CreateUserRoleResp
	DeletePermReq         = user.DeletePermReq
	DeletePermResp        = user.DeletePermResp
	DeleteRoleReq         = user.DeleteRoleReq
	DeleteRoleResp        = user.DeleteRoleResp
	DeleteStrategyReq     = user.DeleteStrategyReq
	DeleteStrategyResp    = user.DeleteStrategyResp
	DeleteUserReq         = user.DeleteUserReq
	DeleteUserResp        = user.DeleteUserResp
	DeleteUserRoleReq     = user.DeleteUserRoleReq
	DeleteUserRoleResp    = user.DeleteUserRoleResp
	GetPermByIdReq        = user.GetPermByIdReq
	GetPermByIdResp       = user.GetPermByIdResp
	GetRoleByIdReq        = user.GetRoleByIdReq
	GetRoleByIdResp       = user.GetRoleByIdResp
	GetStrategyByIdReq    = user.GetStrategyByIdReq
	GetStrategyByIdResp   = user.GetStrategyByIdResp
	GetUserByIdReq        = user.GetUserByIdReq
	GetUserByIdResp       = user.GetUserByIdResp
	GetUserByNameReq      = user.GetUserByNameReq
	GetUserByNameResp     = user.GetUserByNameResp
	GetUserByNumberReq    = user.GetUserByNumberReq
	GetUserByNumberResp   = user.GetUserByNumberResp
	GetUserRoleByIdReq    = user.GetUserRoleByIdReq
	GetUserRoleByIdResp   = user.GetUserRoleByIdResp
	ListAllRoleUsersReq   = user.ListAllRoleUsersReq
	ListAllRoleUsersResp  = user.ListAllRoleUsersResp
	ListPermReq           = user.ListPermReq
	ListPermResp          = user.ListPermResp
	ListRoleForUserReq    = user.ListRoleForUserReq
	ListRoleForUserResp   = user.ListRoleForUserResp
	ListRoleReq           = user.ListRoleReq
	ListRoleResp          = user.ListRoleResp
	ListStrategyReq       = user.ListStrategyReq
	ListStrategyResp      = user.ListStrategyResp
	ListUserReq           = user.ListUserReq
	ListUserResp          = user.ListUserResp
	ListUserRoleReq       = user.ListUserRoleReq
	ListUserRoleResp      = user.ListUserRoleResp
	LoginReq              = user.LoginReq
	LoginResp             = user.LoginResp
	LogoutReq             = user.LogoutReq
	LogoutResp            = user.LogoutResp
	PatchPermReq          = user.PatchPermReq
	PatchPermResp         = user.PatchPermResp
	PatchRoleReq          = user.PatchRoleReq
	PatchRoleResp         = user.PatchRoleResp
	PatchStrategyReq      = user.PatchStrategyReq
	PatchStrategyResp     = user.PatchStrategyResp
	PatchUserReq          = user.PatchUserReq
	PatchUserResp         = user.PatchUserResp
	PatchUserRoleReq      = user.PatchUserRoleReq
	PatchUserRoleResp     = user.PatchUserRoleResp
	PermDetail            = user.PermDetail
	PermInfo              = user.PermInfo
	RefreshTokenReq       = user.RefreshTokenReq
	RefreshTokenResp      = user.RefreshTokenResp
	RegisterReq           = user.RegisterReq
	RegisterResp          = user.RegisterResp
	RemoveRoleForUserReq  = user.RemoveRoleForUserReq
	RemoveRoleForUserResp = user.RemoveRoleForUserResp
	RoleDetail            = user.RoleDetail
	RoleInfo              = user.RoleInfo
	RoleUserDetail        = user.RoleUserDetail
	RoleUserInfo          = user.RoleUserInfo
	StrategyDetail        = user.StrategyDetail
	StrategyInfo          = user.StrategyInfo
	SyncOpenLdapUsersReq  = user.SyncOpenLdapUsersReq
	SyncOpenLdapUsersResp = user.SyncOpenLdapUsersResp
	UpdatePermReq         = user.UpdatePermReq
	UpdatePermResp        = user.UpdatePermResp
	UpdateRoleForUserReq  = user.UpdateRoleForUserReq
	UpdateRoleForUserResp = user.UpdateRoleForUserResp
	UpdateRoleReq         = user.UpdateRoleReq
	UpdateRoleResp        = user.UpdateRoleResp
	UpdateStrategyReq     = user.UpdateStrategyReq
	UpdateStrategyResp    = user.UpdateStrategyResp
	UpdateUserReq         = user.UpdateUserReq
	UpdateUserResp        = user.UpdateUserResp
	UpdateUserRoleReq     = user.UpdateUserRoleReq
	UpdateUserRoleResp    = user.UpdateUserRoleResp
	UserDetail            = user.UserDetail
	UserInfo              = user.UserInfo
	UserRoleDetail        = user.UserRoleDetail
	UserRoleInfo          = user.UserRoleInfo

	AuthService interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		FreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*RefreshTokenResp, error)
		BlackListToken(ctx context.Context, in *BlackListTokenReq, opts ...grpc.CallOption) (*BlackListTokenResp, error)
	}

	defaultAuthService struct {
		cli zrpc.Client
	}
)

func NewAuthService(cli zrpc.Client) AuthService {
	return &defaultAuthService{
		cli: cli,
	}
}

func (m *defaultAuthService) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewAuthServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultAuthService) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error) {
	client := user.NewAuthServiceClient(m.cli.Conn())
	return client.Logout(ctx, in, opts...)
}

func (m *defaultAuthService) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := user.NewAuthServiceClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultAuthService) FreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*RefreshTokenResp, error) {
	client := user.NewAuthServiceClient(m.cli.Conn())
	return client.FreshToken(ctx, in, opts...)
}

func (m *defaultAuthService) BlackListToken(ctx context.Context, in *BlackListTokenReq, opts ...grpc.CallOption) (*BlackListTokenResp, error) {
	client := user.NewAuthServiceClient(m.cli.Conn())
	return client.BlackListToken(ctx, in, opts...)
}