package logic

import (
	"context"
	"gcloud/core/define"
	"time"

	"gcloud/public_repo/rpc/internal/svc"
	"gcloud/public_repo/rpc/public_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublicFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileListLogic {
	return &PublicFileListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublicFileListLogic) PublicFileList(in *public_repo.PublicFileListRequest) (*public_repo.PublicFileListReply, error) {
	// todo: add your logic here and delete this line
	resp := new(public_repo.PublicFileListReply)

	publicFile := make([]*public_repo.PublicFile, 0)
	var cnt int64

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

	err := l.svcCtx.Engine.
		Table("public_repository").
		Select("public_repository.id, public_repository.parent_id, public_repository.identity, "+
			"public_repository.repository_identity, public_repository.ext, public_repository.updated_at,"+
			"public_repository.name, repository_pool.path, repository_pool.size, user_basic.name as owner").
		Where("public_repository.deleted_at = ? OR public_repository.deleted_at IS NULL", time.Time{}.Format(define.Datetime)).
		Joins("left join repository_pool on public_repository.repository_identity = repository_pool.identity").
		Joins("left join user_basic on public_repository.user_identity = user_basic.identity").
		Find(&publicFile).Error
	// Limit(size).
	// Offset(offset).

	if err != nil {
		return resp, err
	}

	// 查询总数
	err = l.svcCtx.Engine.
		Table("public_repository").
		// TODO parent_id = ? AND
		Where("deleted_at IS NULL").
		Count(&cnt).Error
	if err != nil {
		return resp, err
	}

	resp.List = publicFile
	resp.Count = cnt
	return resp, nil
}
