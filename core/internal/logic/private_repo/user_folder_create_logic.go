package private_repo

import (
	"context"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/private_repo/rpc/private_repo"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateReply, err error) {
	resp = new(types.UserFolderCreateReply)
	ufcRet, err := l.svcCtx.PrivateRepoRPC.UserFolderCreate(l.ctx, &private_repo.UserFolderCreateRequest{
		ParentId:     req.ParentId,
		UserIdentity: userIdentity,
		Name:         req.Name,
	})
	resp.Msg = ufcRet.Msg
	if err != nil {
		return
	}
	resp.Identity = ufcRet.Identity
	return
}
