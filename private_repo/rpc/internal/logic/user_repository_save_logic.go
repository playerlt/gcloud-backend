package logic

import (
	"context"
	"gcloud/core/define"
	"gcloud/core/helper"
	"gcloud/core/models"

	"gcloud/private_repo/rpc/internal/svc"
	"gcloud/private_repo/rpc/private_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(in *private_repo.UserRepositorySaveRequest) (*private_repo.UserRepositorySaveReply, error) {
	// todo: add your logic here and delete this line
	resp := new(private_repo.UserRepositorySaveReply)

	usr := &models.UserRepository{
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
		Table("user_repository").
		Select("sum(repository_pool.size) as total_size").
		Where("user_repository.user_identity = ? AND user_repository.deleted_at IS NULL", in.UserIdentity).
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Take(&Size)
	if Size.TotalSize >= define.UserRepositoryMaxSize {
		resp.Msg = "容量不足"
		return resp, *new(error)
	}

	var count int64
	err := l.svcCtx.Engine.
		Table("user_repository").
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
		return resp, *new(error)
	}

	resp.Msg = "success"
	resp.Code = 200
	return resp, nil
}
