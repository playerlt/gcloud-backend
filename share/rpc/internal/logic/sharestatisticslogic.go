package logic

import (
	"context"

	"gcloud/share/rpc/internal/svc"
	"gcloud/share/rpc/share"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareStatisticsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShareStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareStatisticsLogic {
	return &ShareStatisticsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShareStatisticsLogic) ShareStatistics(in *share.ShareStatisticsRequest) (*share.ShareStatisticsReply, error) {
	// todo: add your logic here and delete this line
	resp := &share.ShareStatisticsReply{}

	var shareCount int64
	err := l.svcCtx.Engine.
		Table("share_basic").
		Count(&shareCount).Error
	if err != nil {
		return resp, err
	}

	var Num struct {
		ClickNum int `json:"click_num"`
	}
	err = l.svcCtx.Engine.
		Table("share_basic").
		Select("sum(share_basic.click_num) as click_num").
		Take(&Num).Error
	// Where("share_basic.deleted_at IS NULL").
	if err != nil {
		return resp, err
	}

	resp.ShareCount = int32(int(shareCount))
	resp.ClickNum = int32(Num.ClickNum)
	return resp, nil
}
