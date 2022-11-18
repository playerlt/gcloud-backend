package logic

import (
	"context"

	"gcloud/share/rpc/internal/svc"
	"gcloud/share/rpc/share"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(in *share.ShareBasicDetailRequest) (*share.ShareBasicDetailReply, error) {
	// todo: add your logic here and delete this line
	// logic：其他用户获取分享文件详情
	resp := new(share.ShareBasicDetailReply)
	// 1 更新分享记录的点击次数
	err := l.svcCtx.Engine.
		Table("share_basic").
		Where("identity = ?", in.Identity).
		Exec("UPDATE share_basic SET click_num = click_num + 1 where identity = ?", in.Identity).Error
	if err != nil {
		return resp, nil
	}

	// 2 获取资源详细信息
	err = l.svcCtx.Engine.
		Table("share_basic").
		Select("share_basic.identity, share_basic.repository_identity, user_repository.name, repository_pool.ext, "+
			"repository_pool.path, repository_pool.size, share_basic.click_num, share_basic.desc, "+
			"user_basic.name as owner, user_basic.avatar, share_basic.expired_time, share_basic.updated_at").
		Joins("LEFT JOIN repository_pool ON repository_pool.identity = share_basic.repository_identity").
		Joins("LEFT JOIN user_repository ON user_repository.identity = share_basic.user_repository_identity").
		Joins("left join user_basic on share_basic.user_identity = user_basic.identity").
		Where("share_basic.identity = ?", in.Identity).
		First(resp).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}
