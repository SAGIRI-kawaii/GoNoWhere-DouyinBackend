package logic

import (
	"context"

	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/social/rpc/followclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFriendListLogic {
	return &RelationFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFriendListLogic) RelationFriendList(req *types.Douyin_relation_friend_list_request) (resp *types.Douyin_relation_friend_list_response, err error) {

	res, err := l.svcCtx.SocialRpc.FollowerList(l.ctx, &followclient.DouyinRelationFollowerListRequest{
		UserId: req.UserId,
		Token:  req.Token,
	})

	if err != nil {
		return nil, err
	}

	var UserList []types.Douyin_friend_user_info = make([]types.Douyin_friend_user_info, 0)
	for _, item := range res.UserList {
		UserList = append(UserList, types.Douyin_friend_user_info{
			ID:              item.Id,
			Name:            item.Name,
			FollowCount:     *item.FollowCount,
			FollowerCount:   *item.FollowerCount,
			IsFollow:        item.IsFollow,
			Avatar:          *item.Avatar,
			BackgroundImage: *item.BackgroundImage,
			Signature:       *item.Signature,
			TotalFavorited:  *item.TotalFavorited,
			WorkCount:       *item.WorkCount,
			FavoriteCount:   *item.FavoriteCount,
			// todo: 添加friend参数
			// Message: *item.Message
			// MessageType: *item.MessageType
		})
	}

	return &types.Douyin_relation_friend_list_response{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
		UserList:   UserList,
	}, nil
}
