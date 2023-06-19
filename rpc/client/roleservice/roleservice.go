// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package roleservice

import (
	"context"

	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
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
	ListPermReq           = user.ListPermReq
	ListPermResp          = user.ListPermResp
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
	RoleDetail            = user.RoleDetail
	RoleInfo              = user.RoleInfo
	StrategyDetail        = user.StrategyDetail
	StrategyInfo          = user.StrategyInfo
	SyncOpenLdapUsersReq  = user.SyncOpenLdapUsersReq
	SyncOpenLdapUsersResp = user.SyncOpenLdapUsersResp
	UpdatePermReq         = user.UpdatePermReq
	UpdatePermResp        = user.UpdatePermResp
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

	RoleService interface {
		GetRoleById(ctx context.Context, in *GetRoleByIdReq, opts ...grpc.CallOption) (*GetRoleByIdResp, error)
		CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleResp, error)
		DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleResp, error)
		UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleResp, error)
		PatchRole(ctx context.Context, in *PatchRoleReq, opts ...grpc.CallOption) (*PatchRoleResp, error)
		ListRole(ctx context.Context, in *ListRoleReq, opts ...grpc.CallOption) (*ListRoleResp, error)
	}

	defaultRoleService struct {
		cli zrpc.Client
	}
)

func NewRoleService(cli zrpc.Client) RoleService {
	return &defaultRoleService{
		cli: cli,
	}
}

func (m *defaultRoleService) GetRoleById(ctx context.Context, in *GetRoleByIdReq, opts ...grpc.CallOption) (*GetRoleByIdResp, error) {
	client := user.NewRoleServiceClient(m.cli.Conn())
	return client.GetRoleById(ctx, in, opts...)
}

func (m *defaultRoleService) CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleResp, error) {
	client := user.NewRoleServiceClient(m.cli.Conn())
	return client.CreateRole(ctx, in, opts...)
}

func (m *defaultRoleService) DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleResp, error) {
	client := user.NewRoleServiceClient(m.cli.Conn())
	return client.DeleteRole(ctx, in, opts...)
}

func (m *defaultRoleService) UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleResp, error) {
	client := user.NewRoleServiceClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

func (m *defaultRoleService) PatchRole(ctx context.Context, in *PatchRoleReq, opts ...grpc.CallOption) (*PatchRoleResp, error) {
	client := user.NewRoleServiceClient(m.cli.Conn())
	return client.PatchRole(ctx, in, opts...)
}

func (m *defaultRoleService) ListRole(ctx context.Context, in *ListRoleReq, opts ...grpc.CallOption) (*ListRoleResp, error) {
	client := user.NewRoleServiceClient(m.cli.Conn())
	return client.ListRole(ctx, in, opts...)
}
