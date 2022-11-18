package users

import (
	"context"
	"gcloud/users/rpc/users"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateRequest, userIdentity string) (resp *types.UserUpdateReply, err error) {
	resp = new(types.UserUpdateReply)
	updateRet, err := l.svcCtx.UsersRPC.UserUpdate(l.ctx, &users.UserUpdateRequest{
		Name:         req.Name,
		Password:     req.Password,
		Email:        req.Email,
		Avatar:       req.Avatar,
		UserIdentity: userIdentity,
	})
	_ = updateRet
	if err != nil {
		resp.Msg = "出错了"
		return
	}
	resp.Msg = "success"
	return
}

// 根据字段查询是否存在
func findInfoIsExits(l *UserUpdateLogic, field string, value string) (exits bool) {
	var count int64
	l.svcCtx.Engine.
		Table("user_basic").
		Where(field+" = ?", value).
		Count(&count)
	return count > 0
}

// 根据字段更新
func updateInfo(l *UserUpdateLogic, field string, value string, userIdentity string) (err error) {
	err = l.svcCtx.Engine.
		Table("user_basic").
		Where("identity = ?", userIdentity).
		Update(field, value).Error
	return
}
