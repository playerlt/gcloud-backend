package share

import (
	"context"
	"gcloud/share/rpc/share"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailReply, err error) {
	// logic：其他用户获取分享文件详情
	resp = new(types.ShareBasicDetailReply)
	sbdRet, err := l.svcCtx.ShareRPC.ShareBasicDetail(l.ctx, &share.ShareBasicDetailRequest{
		Identity: req.Identity,
	})
	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Detail = sbdRet
	resp.Msg = "success"
	return
}
