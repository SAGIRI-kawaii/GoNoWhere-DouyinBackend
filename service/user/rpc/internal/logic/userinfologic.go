package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mini-douyin/service/user/model"

	"mini-douyin/service/user/rpc/internal/svc"
	"mini-douyin/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.DouyinUserRequest) (*user.DouyinUserResponse, error) {
	// todo: add your logic here and delete this line
	//查询用户是否存在
	res, err := l.svcCtx.LoginModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(100, "查询用户失败")
	}

	res1, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(100, "查询用户失败")
	}
	return &user.DouyinUserResponse{
		User: &user.DouyinUser{
			Id:            res.Id,
			Name:          res.Name,
			FollowCount:   &res1.FollowerCount,
			FollowerCount: &res1.FollowerCount,
		},
	}, nil
}
