package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mini-douyin/common/jwtx"
	"mini-douyin/service/user/api/internal/svc"
	"mini-douyin/service/user/api/internal/types"
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

	res1, err := jwtx.ParseToken2Uid(l.svcCtx.Config.Auth.AccessSecret, req.Token)
	if err != nil {
		return nil, err
	}
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.DouyinUserRequest{
		UserId: int64(res1),
	})
	if err != nil {
		return nil, err
	}
	return &types.Douyin_user_response{
		StatusCode: int(0),
		StatusMsg:  "success",
		User: types.Douyin_user_info{ID: int(res.User.Id), Name: res.User.Name, FollowCount: 0,
			FollowerCount: 0, IsFollow: false},
	}, nil
}
