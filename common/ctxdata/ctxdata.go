package ctxdata

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyUserName get username from ctx
var CtxKeyUserName = "name"

func GetNameFromCtx(ctx context.Context) string {
	var userName string
	var err bool
	userNameCtx := ctx.Value(CtxKeyUserName)
	userName, err = userNameCtx.(string)
	if err == false {
		logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
	}
	return userName
}
