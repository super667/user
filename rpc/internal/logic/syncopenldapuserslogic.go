package logic

import (
	"context"
	"github.com/super667/user/model"

	"github.com/super667/user/rpc/internal/svc"
	"github.com/super667/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncOpenLdapUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncOpenLdapUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncOpenLdapUsersLogic {
	return &SyncOpenLdapUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SyncOpenLdapUsersLogic) SyncOpenLdapUsers(in *user.SyncOpenLdapUsersReq) (*user.SyncOpenLdapUsersResp, error) {
	resp := &user.SyncOpenLdapUsersResp{}
	ldapUsers, err := l.svcCtx.LdapPool.GetAllUsers(l.ctx)
	if err != nil {
		return resp, err
	}

	for _, user := range ldapUsers {
		u := &model.User{
			Number:   user.Number,
			UserName: user.Name,
			Phone:    user.Mobile,
			Email:    user.Mail,
		}
		err = l.addUser(u)
		if err != nil {
			return resp, err
		}
	}

	return &user.SyncOpenLdapUsersResp{}, nil
}

func (l *SyncOpenLdapUsersLogic) addUser(u *model.User) error {
	exist, err := l.svcCtx.UserModel.Exist(l.ctx, u.UserName)
	if err != nil {
		return err
	}
	if !exist {
		insertRes, err := l.svcCtx.UserModel.Insert(l.ctx, u)
		if err != nil {
			return err
		}
		_, err = insertRes.LastInsertId()
		if err != nil {
			return err
		}
	}
	return nil
}