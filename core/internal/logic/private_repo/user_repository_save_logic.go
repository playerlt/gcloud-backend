package private_repo

import (
	"context"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/private_repo/rpc/private_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, UserIdentity string) (resp *types.UserRepositorySaveReply, err error) {
	// 用户新增文件
	resp = new(types.UserRepositorySaveReply)
	ursRet, err := l.svcCtx.PrivateRepoRPC.UserRepositorySave(l.ctx, &private_repo.UserRepositorySaveRequest{
		ParentId:     req.ParentId,
		UserIdentity: UserIdentity,
		Name:         req.Name,
	})
	resp.Msg = ursRet.Msg
	resp.Code = int(ursRet.Code)
	return
}
