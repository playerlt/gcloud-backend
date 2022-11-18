package public_repo

import (
	"context"
	"gcloud/public_repo/rpc/public_repo"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileNameUpdateLogic {
	return &PublicFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFileNameUpdateLogic) PublicFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	resp = new(types.UserFileNameUpdateReply)
	pfnuRet, err := l.svcCtx.PublicRepoRPC.PublicFileNameUpdate(l.ctx, &public_repo.UserFileNameUpdateRequest{
		Identity:     req.Identity,
		Name:         req.Name,
		UserIdentity: userIdentity,
	})
	resp.Msg = pfnuRet.Msg
	return
}
