package share

import (
	"context"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/share/rpc/share"

	"github.com/zeromicro/go-zero/core/logx"
)

type PopularShareListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPopularShareListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PopularShareListLogic {
	return &PopularShareListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PopularShareListLogic) PopularShareList(req *types.PopularShareListRequest) (resp *types.PopularShareListReply, err error) {
	resp = new(types.PopularShareListReply)
	pslRet, err := l.svcCtx.ShareRPC.PopularShareList(l.ctx, &share.PopularShareListRequest{
		ClickNum: int32(req.ClickNum),
	})

	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.List = pslRet.List
	resp.Msg = "success"
	resp.Code = 0
	return
}
