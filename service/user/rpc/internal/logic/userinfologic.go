package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mini-douyin/common/jwtx"
	"mini-douyin/service/user/model/users"

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

	claims, _ := jwtx.ParseToken(in.Token)
	res, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, claims.UserID)
	if err != nil {
		if err == users.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(100, "查询用户失败")
	}

	return &user.DouyinUserResponse{
		User: &user.DouyinUser{
			Id:              claims.UserID,
			Name:            res.Name,
			FollowCount:     &res.FollowerCount,
			FollowerCount:   &res.FollowerCount,
			IsFollow:        false,
			Avatar:          &res.Avatar.String,
			BackgroundImage: &res.BackgroundImage.String,
			Signature:       &res.Signature.String,
			TotalFavorited:  &res.TotalFavorited.Int64,
			WorkCount:       &res.WorkCount.Int64,
			FavoriteCount:   &res.FavoriteCount.Int64,
		},
	}, nil
}
