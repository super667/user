// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userroleservice

import (
	"context"

	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddRoleForUserReq     = user.AddRoleForUserReq
	AddRoleForUserResp    = user.AddRoleForUserResp
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

	UserRoleService interface {
		GetUserRoleById(ctx context.Context, in *GetUserRoleByIdReq, opts ...grpc.CallOption) (*GetUserRoleByIdResp, error)
		CreateUserRole(ctx context.Context, in *CreateUserRoleReq, opts ...grpc.CallOption) (*CreateUserRoleResp, error)
		DeleteUserRole(ctx context.Context, in *DeleteUserRoleReq, opts ...grpc.CallOption) (*DeleteUserRoleResp, error)
		UpdateUserRole(ctx context.Context, in *UpdateUserRoleReq, opts ...grpc.CallOption) (*UpdateUserRoleResp, error)
		PatchUserRole(ctx context.Context, in *PatchUserRoleReq, opts ...grpc.CallOption) (*PatchUserRoleResp, error)
		ListUserRole(ctx context.Context, in *ListUserRoleReq, opts ...grpc.CallOption) (*ListUserRoleResp, error)
	}

	defaultUserRoleService struct {
		cli zrpc.Client
	}
)

func NewUserRoleService(cli zrpc.Client) UserRoleService {
	return &defaultUserRoleService{
		cli: cli,
	}
}

func (m *defaultUserRoleService) GetUserRoleById(ctx context.Context, in *GetUserRoleByIdReq, opts ...grpc.CallOption) (*GetUserRoleByIdResp, error) {
	client := user.NewUserRoleServiceClient(m.cli.Conn())
	return client.GetUserRoleById(ctx, in, opts...)
}

func (m *defaultUserRoleService) CreateUserRole(ctx context.Context, in *CreateUserRoleReq, opts ...grpc.CallOption) (*CreateUserRoleResp, error) {
	client := user.NewUserRoleServiceClient(m.cli.Conn())
	return client.CreateUserRole(ctx, in, opts...)
}

func (m *defaultUserRoleService) DeleteUserRole(ctx context.Context, in *DeleteUserRoleReq, opts ...grpc.CallOption) (*DeleteUserRoleResp, error) {
	client := user.NewUserRoleServiceClient(m.cli.Conn())
	return client.DeleteUserRole(ctx, in, opts...)
}

func (m *defaultUserRoleService) UpdateUserRole(ctx context.Context, in *UpdateUserRoleReq, opts ...grpc.CallOption) (*UpdateUserRoleResp, error) {
	client := user.NewUserRoleServiceClient(m.cli.Conn())
	return client.UpdateUserRole(ctx, in, opts...)
}

func (m *defaultUserRoleService) PatchUserRole(ctx context.Context, in *PatchUserRoleReq, opts ...grpc.CallOption) (*PatchUserRoleResp, error) {
	client := user.NewUserRoleServiceClient(m.cli.Conn())
	return client.PatchUserRole(ctx, in, opts...)
}

func (m *defaultUserRoleService) ListUserRole(ctx context.Context, in *ListUserRoleReq, opts ...grpc.CallOption) (*ListUserRoleResp, error) {
	client := user.NewUserRoleServiceClient(m.cli.Conn())
	return client.ListUserRole(ctx, in, opts...)
}
