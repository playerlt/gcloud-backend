package users

import (
	"context"
	"gcloud/users/rpc/users"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterCountLogic {
	return &RegisterCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterCountLogic) RegisterCount(req *types.RegisterCountRequest) (resp *types.RegisterCountReply, err error) {
	resp = &types.RegisterCountReply{}
	rcRet, err := l.svcCtx.UsersRPC.RegisterCount(l.ctx, &users.RegisterCountRequest{})
	if err != nil {
		resp.Msg = "出错了"
		return
	}

	resp.Msg = "success"
	resp.Count = rcRet.Count
	return
}
