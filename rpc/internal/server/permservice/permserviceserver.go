// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/super667/user/rpc/internal/logic/permservice"
	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"
)

type PermServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedPermServiceServer
}

func NewPermServiceServer(svcCtx *svc.ServiceContext) *PermServiceServer {
	return &PermServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *PermServiceServer) GetPermById(ctx context.Context, in *user.GetPermByIdReq) (*user.GetPermByIdResp, error) {
	l := permservicelogic.NewGetPermByIdLogic(ctx, s.svcCtx)
	return l.GetPermById(in)
}

func (s *PermServiceServer) CreatePerm(ctx context.Context, in *user.CreatePermReq) (*user.CreatePermResp, error) {
	l := permservicelogic.NewCreatePermLogic(ctx, s.svcCtx)
	return l.CreatePerm(in)
}

func (s *PermServiceServer) DeletePerm(ctx context.Context, in *user.DeletePermReq) (*user.DeletePermResp, error) {
	l := permservicelogic.NewDeletePermLogic(ctx, s.svcCtx)
	return l.DeletePerm(in)
}

func (s *PermServiceServer) UpdatePerm(ctx context.Context, in *user.UpdatePermReq) (*user.UpdatePermResp, error) {
	l := permservicelogic.NewUpdatePermLogic(ctx, s.svcCtx)
	return l.UpdatePerm(in)
}

func (s *PermServiceServer) PatchPerm(ctx context.Context, in *user.PatchPermReq) (*user.PatchPermResp, error) {
	l := permservicelogic.NewPatchPermLogic(ctx, s.svcCtx)
	return l.PatchPerm(in)
}

func (s *PermServiceServer) ListPerm(ctx context.Context, in *user.ListPermReq) (*user.ListPermResp, error) {
	l := permservicelogic.NewListPermLogic(ctx, s.svcCtx)
	return l.ListPerm(in)
}
