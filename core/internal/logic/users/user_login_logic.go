package users

import (
	"context"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/users/rpc/users"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	resp = new(types.LoginReply)
	loginRet, err := l.svcCtx.UsersRPC.UserLogin(l.ctx, &users.LoginRequest{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		resp.Msg = loginRet.Msg
		return
	}
	resp.Msg = loginRet.Msg
	resp.Code = int(loginRet.Code)
	resp.Token = loginRet.Token
	resp.RefreshToken = loginRet.RefreshToken
	return
}
