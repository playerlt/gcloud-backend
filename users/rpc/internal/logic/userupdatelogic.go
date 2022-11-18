package logic

import (
	"context"
	"gcloud/core/helper"
	"gcloud/core/models"

	"gcloud/users/rpc/internal/svc"
	"gcloud/users/rpc/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *users.UserUpdateRequest) (*users.UserUpdateReply, error) {
	// todo: add your logic here and delete this line
	resp := new(users.UserUpdateReply)
	user := &models.UserBasic{
		Name:   in.Name,
		Email:  in.Email,
		Avatar: in.Avatar,
	}
	if in.Password != "" {
		user.Password = helper.Md5(in.Password)
	}

	err := l.svcCtx.Engine.
		Table("user_basic").
		Where("identity = ?", in.UserIdentity).
		Updates(user).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}
