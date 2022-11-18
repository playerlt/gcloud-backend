package logic

import (
	"context"

	"gcloud/users/rpc/internal/svc"
	"gcloud/users/rpc/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDetailLogic) UserDetail(in *users.UserDetailRequest) (*users.UserDetailReply, error) {
	// todo: add your logic here and delete this line

	return &users.UserDetailReply{}, nil
}
