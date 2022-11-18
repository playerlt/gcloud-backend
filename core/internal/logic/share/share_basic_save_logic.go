package share

import (
	"context"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/share/rpc/share"
	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveReply, err error) {
	// logic：其他用户保存分享文件
	resp = new(types.ShareBasicSaveReply)
	sbsRet, err := l.svcCtx.ShareRPC.ShareBasicSave(l.ctx, &share.ShareBasicSaveRequest{
		UserIdentity:       userIdentity,
		RepositoryIdentity: req.RepositoryIdentity,
		ParentId:           req.ParentId,
	})
	resp.Msg = sbsRet.Msg
	resp.Code = int(sbsRet.Code)
	resp.Identity = sbsRet.Identity
	return
}
