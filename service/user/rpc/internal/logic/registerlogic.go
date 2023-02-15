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

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	// todo: add your logic here and delete this line

	_, err := l.svcCtx.LoginModel.FindOneByName(l.ctx, in.Username)

	if err == nil {
		return nil, status.Error(100, "用户名已存在")
	}

	if err == model.ErrNotFound {
		newUser := model.Logins{
			Name:     in.Username,
			PassWord: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		res, err := l.svcCtx.LoginModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(100, "注册失败")
		}
		newUser.Id, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(100, "error")
		}
		_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.Users{
			Id:            newUser.Id,
			Name:          newUser.Name,
			UserId:        newUser.Id,
			FollowCount:   0,
			FollowerCount: 0,
		})
		if err != nil {
			return nil, status.Error(100, "注册失败")
		}
		return &user.DouyinUserRegisterResponse{
			UserId: newUser.Id,
		}, nil
	}

	return nil, status.Error(100, "注册失败")
}
