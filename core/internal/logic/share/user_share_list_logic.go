package share

import (
	"context"
	"gcloud/share/rpc/share"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserShareListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserShareListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserShareListLogic {
	return &UserShareListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserShareListLogic) UserShareList(req *types.UserShareListRequest, UserIdentity string) (resp *types.UserShareListReply, err error) {
	resp = new(types.UserShareListReply)

	usRet, err := l.svcCtx.ShareRPC.UserShareList(l.ctx, &share.UserShareListRequest{
		UserIdentity: UserIdentity,
	})
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.List = usRet.List
	resp.Msg = "success"
	resp.Code = 0

	return
}
