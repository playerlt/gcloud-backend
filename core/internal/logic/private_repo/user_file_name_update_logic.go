package private_repo

import (
	"context"
	"gcloud/private_repo/rpc/private_repo"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	resp = new(types.UserFileNameUpdateReply)
	ufnuRet, err := l.svcCtx.PrivateRepoRPC.UserFileNameUpdate(l.ctx, &private_repo.UserFileNameUpdateRequest{
		Identity:     req.Identity,
		UserIdentity: userIdentity,
		Name:         req.Name,
	})
	resp.Msg = ufnuRet.Msg
	return
}
