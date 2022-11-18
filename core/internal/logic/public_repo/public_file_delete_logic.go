package public_repo

import (
	"context"
	"gcloud/public_repo/rpc/public_repo"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileDeleteLogic {
	return &PublicFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFileDeleteLogic) PublicFileDelete(req *types.UserFileDeleteRequest, userIdentity string) (resp *types.UserFileDeleteReply, err error) {
	resp = new(types.UserFileDeleteReply)
	pfdRet, err := l.svcCtx.PublicRepoRPC.PublicFileDelete(l.ctx, &public_repo.UserFileDeleteRequest{
		Identity:     req.Identity,
		UserIdentity: userIdentity,
	})
	_ = pfdRet
	return
}
