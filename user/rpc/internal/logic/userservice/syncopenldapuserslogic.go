package userservicelogic

import (
	"context"
	"github.com/super667/user/common/cryptx"
	"github.com/super667/user/user/model"

	"github.com/super667/user/user/rpc/internal/svc"
	"github.com/super667/user/user/rpc/user"

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
			UserName:     user.Name,
			NickName:     user.DisplayName,
			Email:        user.Mail,
			Number:       user.Number,
			Phone:        user.Mobile,
			Address:      user.PostalAddress,
			DeptCode:     user.DeptCode,
			Position:     user.Position,
			Introduction: user.CN,
			Source:       "openldap",
			DeptName:     user.DeptName,
			UserDn:       user.DN,
			Password:     cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, l.svcCtx.Config.Ldap.UserInitPassword),
		}
		err = l.addUser(u)
		if err != nil {
			return resp, err
		}
	}

	return &user.SyncOpenLdapUsersResp{Res: "完成"}, nil
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
