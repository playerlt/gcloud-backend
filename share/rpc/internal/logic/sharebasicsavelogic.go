package logic

import (
	"context"
	"gcloud/core/define"
	"gcloud/core/helper"
	"gcloud/core/models"

	"gcloud/share/rpc/internal/svc"
	"gcloud/share/rpc/share"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(in *share.ShareBasicSaveRequest) (*share.ShareBasicSaveReply, error) {
	// todo: add your logic here and delete this line
	resp := new(share.ShareBasicSaveReply)
	// 获取资源详情 from repository_pool
	rp := new(models.RepositoryPool)
	err := l.svcCtx.Engine.
		Table("repository_pool").
		Where("identity = ?", in.RepositoryIdentity).
		First(rp).Error
	if err != nil {
		resp.Msg = "error"
		return resp, err
	}
	if rp.Id == 0 {
		resp.Msg = "资源不存在"
		return resp, *new(error)
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
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("name = ? AND parent_id = ? AND user_identity = ? AND deleted_at IS NULL", rp.Name, in.ParentId, in.UserIdentity).
		Count(&count).Error
	if count > 0 {
		resp.Msg = "exist"
		resp.Code = 405
		return resp, *new(error)
	}

	// 资源保存 to user_repository
	usr := &models.UserRepository{
		Identity:           helper.UUID(),
		UserIdentity:       in.UserIdentity,
		ParentId:           in.ParentId,
		RepositoryIdentity: in.RepositoryIdentity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}

	err = l.svcCtx.Engine.
		Select("identity", "parent_id", "user_identity", "repository_identity", "name", "ext", "created_at", "updated_at").
		Create(usr).Error
	if err != nil {
		resp.Msg = "save error"
		return resp, err
	}

	resp.Identity = usr.Identity
	resp.Msg = "success"
	return resp, nil
}
