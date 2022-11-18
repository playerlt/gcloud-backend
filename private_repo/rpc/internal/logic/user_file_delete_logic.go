package logic

import (
	"context"
	"gcloud/core/models"

	"gcloud/private_repo/rpc/internal/svc"
	"gcloud/private_repo/rpc/private_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(in *private_repo.UserFileDeleteRequest) (*private_repo.UserFileDeleteReply, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.Engine.
		Where("user_identity = ? AND identity = ?", in.UserIdentity, in.Identity).
		Delete(new(models.UserRepository)).Error
	resp := new(private_repo.UserFileDeleteReply)

	if err != nil {
		return resp, err
	}
	return &private_repo.UserFileDeleteReply{}, nil
}
