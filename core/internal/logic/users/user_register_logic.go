package users

import (
	"context"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/users/rpc/users"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	resp = new(types.UserRegisterReply)
	registerRet, err := l.svcCtx.UsersRPC.UserRegister(l.ctx, &users.UserRegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Code:     req.Code,
		Password: req.Password,
	})
	if err != nil {
		resp.Msg = registerRet.Msg
		resp.Code = int(registerRet.Code)
		return
	}
	resp.Msg = "注册成功"
	resp.Code = 200
	return
}

func Random() {
	panic("unimplemented")
}
