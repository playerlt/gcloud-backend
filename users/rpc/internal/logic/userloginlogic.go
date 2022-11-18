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

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *users.LoginRequest) (*users.LoginReply, error) {
	// todo: add your logic here and delete this line
	resp := new(users.LoginReply)
	if len(in.Name) < 6 {
		resp.Msg = "用户名长度不能小于6位"
		return resp, *new(error)
	}
	if len(in.Password) < 6 {
		resp.Msg = "密码长度不能小于6位"
		return resp, *new(error)
	}

	// 从数据库中查询当前用户
	user := new(models.UserBasic)
	l.svcCtx.Engine.
		Where("name = ? AND password = ?", in.Name, helper.Md5(in.Password)).
		First(&user)
	// has, err := models.Engine.Where("name = ? AND password = ?", req.Name, req.Password).Get(user)

	if user.Id == 0 {
		resp.Msg = "用户名或密码错误"
		return resp, *new(error)
	}

	// 生成普通 token1
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name, define.TokenExpire)
	if err != nil {
		resp.Msg = "生成token失败"
		return resp, err
	}
	// 生成用于刷新 token1 的 token2
	// 当 token1 失效后，使用 token2 生成新 token1
	refreshToken, err := helper.GenerateToken(user.Id, user.Identity, user.Name, define.RefreshTokenExpire)
	if err != nil {
		resp.Msg = "生成token失败"
		return resp, err
	}

	resp.Token = token
	resp.RefreshToken = refreshToken
	resp.Msg = "用户登录成功"
	return resp, nil
}
