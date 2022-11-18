package logic

import (
	"context"

	"gcloud/users/rpc/internal/svc"
	"gcloud/users/rpc/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(in *users.RefreshAuthorizationRequest) (*users.RefreshAuthorizationReply, error) {
	// todo: add your logic here and delete this line

	return &users.RefreshAuthorizationReply{}, nil
}
