package logic

import (
	"context"
	"gcloud/core/helper"
	"gcloud/core/models"

	"gcloud/share/rpc/internal/svc"
	"gcloud/share/rpc/share"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(in *share.ShareBasicCreateRequest) (*share.ShareBasicCreateReply, error) {
	// todo: add your logic here and delete this line
	resp := new(share.ShareBasicCreateReply)
	idna := helper.UUID()
	usr := new(models.UserRepository)
	err := l.svcCtx.Engine.
		Where("identity = ?", in.UserRepositoryIdentity).
		First(usr).Error
	if err != nil {
		return resp, err
	}
	if usr.Id == 0 {
		resp.Msg = "user resource not found"
		return resp, *new(error)
	}

	data := &models.ShareBasic{
		Identity:               idna,
		UserIdentity:           in.UserIdentity,
		UserRepositoryIdentity: in.UserRepositoryIdentity,
		RepositoryIdentity:     usr.RepositoryIdentity,
		ExpiredTime:            int(in.ExpiredTime),
		Desc:                   in.Desc,
	}
	err = l.svcCtx.Engine.
		Select("identity", "user_identity", "repository_identity", "user_repository_identity", "expired_time", "desc", "created_at", "updated_at").
		Create(data).Error
	if err != nil {
		resp.Msg = "error"
		return resp, err
	}

	resp.Identity = idna
	resp.Msg = "success"
	return resp, nil
}
