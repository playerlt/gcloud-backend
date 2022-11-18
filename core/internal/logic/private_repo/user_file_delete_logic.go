package private_repo

import (
	"context"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/private_repo/rpc/private_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest, userIdentity string) (resp *types.UserFileDeleteReply, err error) {
	resp = new(types.UserFileDeleteReply)
	ufdRet, err := l.svcCtx.PrivateRepoRPC.UserFileDelete(l.ctx, &private_repo.UserFileDeleteRequest{
		Identity:     req.Identity,
		UserIdentity: userIdentity,
	})
	if err != nil {
		resp.Msg = "error"
		return
	}
	_ = ufdRet
	resp.Msg = "success"
	return
}
