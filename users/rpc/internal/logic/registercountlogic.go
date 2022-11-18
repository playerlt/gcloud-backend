package logic

import (
	"context"

	"gcloud/users/rpc/internal/svc"
	"gcloud/users/rpc/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterCountLogic {
	return &RegisterCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterCountLogic) RegisterCount(in *users.RegisterCountRequest) (*users.RegisterCountReply, error) {
	// todo: add your logic here and delete this line
	resp := &users.RegisterCountReply{}
	var count int64
	err := l.svcCtx.Engine.
		Table("user_basic").
		Count(&count).Error
	if err != nil {
		return resp, *new(error)
	}
	resp.Count = count
	return resp, nil
}
