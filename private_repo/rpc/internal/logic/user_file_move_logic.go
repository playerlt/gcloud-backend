package logic

import (
	"context"
	"gcloud/core/models"

	"gcloud/private_repo/rpc/internal/svc"
	"gcloud/private_repo/rpc/private_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserFileMoveLogic) UserFileMove(in *private_repo.UserFileMoveRequest) (*private_repo.UserFileMoveReply, error) {
	// todo: add your logic here and delete this line
	resp := new(private_repo.UserFileMoveReply)
	// parentId
	parentData := new(models.UserRepository)
	err := l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND user_identity = ?", in.ParentIdentity, in.UserIdentity).
		First(parentData).Error
	if err != nil {
		resp.Msg = "error"
		return resp, err
	}
	if parentData.Id == 0 {
		resp.Msg = "文件夹不存在"
		return resp, err
	}

	// 更新记录的ParentId
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND deleted_at IS NULL", in.Identity).
		Update("parent_id", int64(parentData.Id)).Error
	if err != nil {
		resp.Msg = "error"
		return resp, err
	}
	resp.Msg = "success"
	return resp, nil
}
