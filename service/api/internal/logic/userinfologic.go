package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/user/rpc/userclient"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.Douyin_user_request) (resp *types.Douyin_user_response, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.DouyinUserRequest{
		Token: req.Token,
	})
	if err != nil {
		return nil, err
	}
	return &types.Douyin_user_response{
		StatusCode: 0,
		StatusMsg:  "success",
		User: types.Douyin_user_info{
			ID:              res.User.Id,
			Name:            res.User.Name,
			FollowCount:     *res.User.FollowCount,
			FollowerCount:   *res.User.FollowerCount,
			IsFollow:        res.User.IsFollow,
			Avatar:          *res.User.Avatar,
			BackgroundImage: *res.User.BackgroundImage,
			Signature:       *res.User.Signature,
			TotalFavorited:  *res.User.TotalFavorited,
			WorkCount:       *res.User.WorkCount,
			FavoriteCount:   *res.User.FavoriteCount,
		},
	}, nil
}
