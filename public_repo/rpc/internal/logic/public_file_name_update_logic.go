package logic

import (
	"context"

	"gcloud/public_repo/rpc/internal/svc"
	"gcloud/public_repo/rpc/public_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileNameUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublicFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileNameUpdateLogic {
	return &PublicFileNameUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublicFileNameUpdateLogic) PublicFileNameUpdate(in *public_repo.UserFileNameUpdateRequest) (*public_repo.UserFileNameUpdateReply, error) {
	// todo: add your logic here and delete this line
	resp := new(public_repo.UserFileNameUpdateReply)
	if in.Name == "" {
		resp.Msg = "文件名为空"
		return resp, *new(error)
	}

	// 判断当前文件名在该层级下是否已存在
	var cnt int64
	err := l.svcCtx.Engine.
		Table("public_repository").
		Where("name = ?", in.Name).
		Where("parent_id = (select parent_id from public_repository ur where ur.identity = ?)", in.Identity).
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
		Table("public_repository").
		Where("identity = ? AND user_identity = ?", in.Identity, in.UserIdentity).
		Update("name", in.Name).Error

	if err != nil {
		resp.Msg = "error"
		return resp, err
	}
	resp.Msg = "success"
	return resp, nil
}
