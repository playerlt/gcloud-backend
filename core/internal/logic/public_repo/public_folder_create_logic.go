package public_repo

import (
	"context"
	"gcloud/public_repo/rpc/public_repo"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFolderCreateLogic {
	return &PublicFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFolderCreateLogic) PublicFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateReply, err error) {
	resp = new(types.UserFolderCreateReply)
	pfcRet, err := l.svcCtx.PublicRepoRPC.PublicFolderCreate(l.ctx, &public_repo.UserFolderCreateRequest{
		ParentId:     req.ParentId,
		Name:         req.Name,
		UserIdentity: userIdentity,
	})
	resp.Msg = pfcRet.Msg
	if err == nil {
		resp.Identity = pfcRet.Identity
	}
	return
}
