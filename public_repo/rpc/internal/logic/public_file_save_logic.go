package logic

import (
	"context"
	"gcloud/core/define"
	"gcloud/core/helper"
	"gcloud/core/models"

	"gcloud/public_repo/rpc/internal/svc"
	"gcloud/public_repo/rpc/public_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublicFileSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileSaveLogic {
	return &PublicFileSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublicFileSaveLogic) PublicFileSave(in *public_repo.PublicRepositorySaveRequest) (*public_repo.PublicRepositorySaveReply, error) {
	// todo: add your logic here and delete this line
	resp := new(public_repo.PublicRepositorySaveReply)

	usr := &models.PublicRepository{
		Identity:           helper.UUID(),
		UserIdentity:       in.UserIdentity,
		ParentId:           in.ParentId,
		RepositoryIdentity: in.RepositoryIdentity,
		Name:               in.Name,
		Ext:                in.Ext,
	}

	var Size struct {
		TotalSize int `json:"total_size"`
	}
	l.svcCtx.Engine.
		Table("public_repository").
		Select("sum(repository_pool.size) as total_size").
		Where("public_repository.user_identity = ? AND public_repository.deleted_at IS NULL", in.UserIdentity).
		Joins("left join repository_pool on public_repository.repository_identity = repository_pool.identity").
		Take(&Size)
	if in.UserIdentity != "USER_1" && Size.TotalSize >= define.PublicRepositoryMaxSize {
		resp.Msg = "容量不足"
		return resp, *new(error)
	}

	var count int64
	err := l.svcCtx.Engine.
		Table("public_repository").
		Where("name = ? AND parent_id = ? AND user_identity = ? AND deleted_at IS NULL", in.Name, in.ParentId, in.UserIdentity).
		Count(&count).Error
	if count > 0 {
		resp.Msg = "exist"
		resp.Code = 405
		return resp, *new(error)
	}

	err = l.svcCtx.Engine.
		Select("identity", "parent_id", "user_identity", "repository_identity", "name", "ext", "created_at", "updated_at").
		Create(usr).Error
	if err != nil {
		resp.Msg = "error"
		return resp, err
	}

	resp.Msg = "success"
	resp.Code = 200
	return resp, nil
}
