package share

import (
	"context"
	"gcloud/share/rpc/share"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareStatisticsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareStatisticsLogic {
	return &ShareStatisticsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareStatisticsLogic) ShareStatistics(req *types.ShareStatisticsRequest) (resp *types.ShareStatisticsReply, err error) {
	resp = &types.ShareStatisticsReply{}
	ssRet, err := l.svcCtx.ShareRPC.ShareStatistics(l.ctx, &share.ShareStatisticsRequest{})
	// Where("share_basic.deleted_at IS NULL").
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.ShareCount = int(ssRet.ShareCount)
	resp.ClickNum = int(ssRet.ClickNum)
	resp.Msg = "success"
	resp.Code = 0
	return
}
