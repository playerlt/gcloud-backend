package logic

import (
	"context"

	"gcloud/share/rpc/internal/svc"
	"gcloud/share/rpc/share"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserShareListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserShareListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserShareListLogic {
	return &UserShareListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserShareListLogic) UserShareList(in *share.UserShareListRequest) (*share.UserShareListReply, error) {
	// todo: add your logic here and delete this line
	shareFile := make([]*share.ShareBasicDetailReply, 0)
	resp := new(share.UserShareListReply)

	err := l.svcCtx.Engine.
		Table("share_basic").
		Select("share_basic.identity, share_basic.repository_identity, user_repository.name, repository_pool.ext, "+
			"repository_pool.path, repository_pool.size, share_basic.click_num, share_basic.desc, "+
			"user_basic.name as owner, user_basic.avatar, share_basic.expired_time, share_basic.updated_at").
		Joins("LEFT JOIN repository_pool ON repository_pool.identity = share_basic.repository_identity").
		Joins("LEFT JOIN user_repository ON user_repository.identity = share_basic.user_repository_identity").
		Joins("left join user_basic on share_basic.user_identity = user_basic.identity").
		Where("share_basic.user_identity = ? ", in.UserIdentity).
		Where("share_basic.deleted_at IS NULL").
		Order("share_basic.updated_at desc").
		Find(&shareFile).Error

	if err != nil {
		return resp, err
	}

	resp.List = shareFile
	return resp, nil
}
