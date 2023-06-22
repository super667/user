package authservicelogic

import (
	"context"
	"github.com/super667/user/user/model"
	"google.golang.org/grpc/status"

	"github.com/super667/user/common/cryptx"
	"github.com/super667/user/user/rpc/internal/svc"
	"github.com/super667/user/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// 判断手机号是否已经注册
	_, err := l.svcCtx.UserModel.GetOne(l.ctx, in.Phone)
	if err == nil {
		return nil, status.Error(100, "该用户已存在")
	}

	if err == model.ErrNotFound {
		newUser := model.User{
			UserName: in.UserName,
			Phone:    in.Phone,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		newUser.Id, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &user.RegisterResp{
			Id:       newUser.Id,
			UserName: newUser.UserName,
			Phone:    newUser.Phone,
		}, nil

	}

	return nil, status.Error(500, err.Error())
}
