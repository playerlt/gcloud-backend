package logic

import (
	"context"
	"gcloud/core/helper"
	"gcloud/core/models"

	"gcloud/public_repo/rpc/internal/svc"
	"gcloud/public_repo/rpc/public_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFolderCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublicFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFolderCreateLogic {
	return &PublicFolderCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublicFolderCreateLogic) PublicFolderCreate(in *public_repo.UserFolderCreateRequest) (*public_repo.UserFolderCreateReply, error) {
	// todo: add your logic here and delete this line
	resp := new(public_repo.UserFolderCreateReply)

	if in.Name == "" {
		resp.Msg = "name is empty"
		return resp, *new(error)
	}

	// 判断当前文件名在该层级下是否已存在
	var cnt int64
	err := l.svcCtx.Engine.
		Table("public_repository").
		Where("name = ? AND parent_id = ? AND user_identity = ?", in.Name, in.ParentId, in.UserIdentity).
		Count(&cnt).Error

	if err != nil {
		resp.Msg = "error"
		return resp, err
	}
	if cnt > 0 {
		resp.Msg = "exits"
		return resp, *new(error)
	}

	// 创建文件夹
	data := &models.PublicRepository{
		Identity:     helper.UUID(),
		UserIdentity: in.UserIdentity,
		ParentId:     in.ParentId,
		Name:         in.Name,
	}
	err = l.svcCtx.Engine.
		Table("public_repository").
		Select("identity", "name", "user_identity", "parent_id", "created_at", "updated_at").
		Create(data).Error
	if err != nil {
		resp.Msg = "error"
		return resp, err
	}
	resp.Msg = "success"
	return resp, nil
}
