package logic

import (
	"context"
	"gcloud/core/define"
	"gcloud/core/helper"
	"gcloud/core/models"

	"gcloud/users/rpc/internal/svc"
	"gcloud/users/rpc/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *users.UserRegisterRequest) (*users.UserRegisterReply, error) {
	// todo: add your logic here and delete this line
	resp := new(users.UserRegisterReply)
	if len(in.Name) < 6 {
		resp.Msg = "用户名长度不能小于6位"
		return resp, *new(error)
	}
	if len(in.Password) < 6 {
		resp.Msg = "密码长度不能小于6位"
		return resp, *new(error)
	}

	// 判断code是否一致
	code, err := l.svcCtx.RDB.Get(l.ctx, in.Email).Result()
	if err != nil {
		resp.Msg = "无效验证码"
		return resp, err
	}
	if code != in.Code {
		resp.Msg = "验证码错误"
		return resp, *new(error)
	}

	// 判断用户名是否已存在
	var count int64
	err = l.svcCtx.Engine.
		Table("user_basic").
		Where("name = ?", in.Name).
		Count(&count).Error
	if err != nil {
		resp.Msg = "出错了"
		return resp, err
	}
	if count > 0 {
		resp.Msg = "用户名已存在"
		return resp, *new(error)
	}

	// 入库
	user := &models.UserBasic{
		Identity: helper.UUID(),
		Name:     in.Name,
		Email:    in.Email,
		Avatar:   define.AvatarBaseUrl + helper.Random() + ".png",
		Password: helper.Md5(in.Password),
		Capacity: define.UserRepositoryMaxSize,
	}
	// fix: 需指定添加字段 (Select())，不推荐使用 Omit()
	err = l.svcCtx.Engine.
		Select("identity", "name", "email", "password", "capacity", "created_at", "updated_at").
		Create(user).Error
	if err != nil {
		resp.Msg = "出错了"
		return resp, err
	}
	return resp, nil
}
