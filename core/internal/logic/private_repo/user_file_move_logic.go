package private_repo

import (
	"context"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/private_repo/rpc/private_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	resp = new(types.UserFileMoveReply)
	ufmRet, err := l.svcCtx.PrivateRepoRPC.UserFileMove(l.ctx, &private_repo.UserFileMoveRequest{
		Identity:       req.Identity,
		ParentIdentity: req.ParentIdentity,
		UserIdentity:   userIdentity,
	})
	_ = ufmRet
	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Msg = "success"
	return
}
