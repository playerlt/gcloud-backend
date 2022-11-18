package logic

import (
	"context"
	"gcloud/core/models"

	"gcloud/public_repo/rpc/internal/svc"
	"gcloud/public_repo/rpc/public_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublicFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileDeleteLogic {
	return &PublicFileDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublicFileDeleteLogic) PublicFileDelete(in *public_repo.UserFileDeleteRequest) (*public_repo.UserFileDeleteReply, error) {
	// todo: add your logic here and delete this line
	resp := new(public_repo.UserFileDeleteReply)

	err := l.svcCtx.Engine.
		Where("user_identity = ? AND identity = ?", in.UserIdentity, in.Identity).
		Delete(new(models.PublicRepository)).Error

	if err != nil {
		return resp, err
	}
	return resp, nil
}
