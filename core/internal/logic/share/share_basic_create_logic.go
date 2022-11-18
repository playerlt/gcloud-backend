package share

import (
	"context"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/share/rpc/share"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	resp = new(types.ShareBasicCreateReply)
	sbcRet, err := l.svcCtx.ShareRPC.ShareBasicCreate(l.ctx, &share.ShareBasicCreateRequest{
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		ExpiredTime:            int32(req.ExpiredTime),
		Desc:                   req.Desc,
	})
	resp.Msg = sbcRet.Msg
	if err != nil {
		return
	}
	resp.Identity = sbcRet.Identity
	return
}
