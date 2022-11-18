package public_repo

import (
	"context"
	"gcloud/public_repo/rpc/public_repo"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFileSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileSaveLogic {
	return &PublicFileSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFileSaveLogic) PublicFileSave(req *types.PublicRepositorySaveRequest, UserIdentity string) (resp *types.PublicRepositorySaveReply, err error) {
	resp = new(types.PublicRepositorySaveReply)
	pfsRet, err := l.svcCtx.PublicRepoRPC.PublicFileSave(l.ctx, &public_repo.PublicRepositorySaveRequest{
		RepositoryIdentity: req.RepositoryIdentity,
		Name:               req.Name,
		ParentId:           req.ParentId,
		Ext:                req.Ext,
		UserIdentity:       UserIdentity,
	})
	resp.Code = int(pfsRet.Code)
	resp.Msg = pfsRet.Msg
	return
}
