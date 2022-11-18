package logic

import (
	"context"

	"gcloud/private_repo/rpc/internal/svc"
	"gcloud/private_repo/rpc/private_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(in *private_repo.UserFileNameUpdateRequest) (*private_repo.UserFileNameUpdateReply, error) {
	// todo: add your logic here and delete this line
	resp := new(private_repo.UserFileNameUpdateReply)
	if in.Name == "" {
		resp.Msg = "文件名为空"
		return resp, *new(error)
	}

	// 判断当前文件名在该层级下是否已存在
	var cnt int64
	err := l.svcCtx.Engine.
		Table("user_repository").
		Where("name = ?", in.Name).
		Where("parent_id = (select parent_id from user_repository ur where ur.identity = ?)", in.Identity).
		Where("deleted_at IS NULL").
		Count(&cnt).Error

	if err != nil {
		resp.Msg = "error"
		return resp, err
	}
	if cnt > 0 {
		resp.Msg = "文件名已存在"
		return resp, *new(error)
	}

	// 更新文件名
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND user_identity = ?", in.Identity, in.UserIdentity).
		Update("name", in.Name).Error

	if err != nil {
		resp.Msg = "error"
		return resp, err
	}
	resp.Msg = "success"
	return resp, nil
}
