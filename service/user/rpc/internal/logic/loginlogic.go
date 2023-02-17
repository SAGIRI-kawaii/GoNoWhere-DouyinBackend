package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mini-douyin/common/cryptx"
	"mini-douyin/service/user/model"

	"mini-douyin/service/user/rpc/internal/svc"
	"mini-douyin/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.DouyinUserLoginRequest) (*user.DouyinUserLoginResponse, error) {
	// todo: add your logic here and delete this line
	//查询用户是否存在
	res, err := l.svcCtx.LoginModel.FindOneByName(l.ctx, in.Username)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(100, "查询用户失败")
	}

	//判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if res.PassWord != password {
		return nil, status.Error(100, "密码错误")
	}
	return &user.DouyinUserLoginResponse{
		UserId: res.Id,
	}, nil

}
