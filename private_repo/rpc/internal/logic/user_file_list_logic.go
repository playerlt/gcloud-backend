package logic

import (
	"context"
	"gcloud/core/define"
	"time"

	"gcloud/private_repo/rpc/internal/svc"
	"gcloud/private_repo/rpc/private_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserFileListLogic) UserFileList(in *private_repo.UserFileListRequest) (*private_repo.UserFileListReply, error) {
	// todo: add your logic here and delete this line
	usrFile := make([]*private_repo.UserFile, 0)
	deletedFile := make([]*private_repo.DeletedUserFile, 0)
	var cnt int64
	resp := new(private_repo.UserFileListReply)

	// 分页参数
	size := in.Size
	if size == 0 {
		size = int32(define.PageSize)
	}
	page := in.Page
	if page == 0 {
		page = 1
	}
	// offset := (page - 1) * size

	// TODO 按文件名查询
	// id := req.Id =》 parent_id
	// if id == 0 {
	// 	id = -1
	// }

	err := l.svcCtx.Engine.
		Table("user_repository").
		Select("user_repository.id, user_repository.parent_id, user_repository.identity, "+
			"user_repository.repository_identity, user_repository.ext, user_repository.updated_at,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Where("user_identity = ? ", in.UserIdentity).
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at IS NULL", time.Time{}.Format(define.Datetime)).
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Find(&usrFile).Error
	// Limit(size).
	// Offset(offset).
	if err != nil {
		return resp, err
	}

	err = l.svcCtx.Engine.
		Table("user_repository").
		Select("user_repository.id, user_repository.parent_id, user_repository.identity, "+
			"user_repository.repository_identity, user_repository.ext, user_repository.deleted_at,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Where("user_identity = ? ", in.UserIdentity).
		Where("user_repository.deleted_at IS NOT NULL").
		// Order("user_repository.deleted_at desc").
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Find(&deletedFile).Error

	if err != nil {
		return resp, err
	}

	// 查询总数
	err = l.svcCtx.Engine.
		Table("user_repository").
		// TODO parent_id = ? AND
		Where("user_identity = ? AND deleted_at IS NULL", in.UserIdentity).
		Count(&cnt).Error
	if err != nil {
		return resp, err
	}

	resp.List = usrFile
	resp.DeletedList = deletedFile
	resp.Count = cnt
	return resp, nil
}
