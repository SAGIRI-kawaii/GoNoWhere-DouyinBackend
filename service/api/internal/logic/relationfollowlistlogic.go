package logic

import (
	"context"

	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/social/rpc/followclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowListLogic) RelationFollowList(req *types.Douyin_relation_follow_list_request) (resp *types.Douyin_relation_follow_list_response, err error) {

	res, err := l.svcCtx.SocialRpc.FollowList(l.ctx, &followclient.DouyinRelationFollowListRequest{
		UserId: req.UserId,
		Token:  req.Token,
	})

	if err != nil {
		return nil, err
	}

	var UserList []types.Douyin_user_info = make([]types.Douyin_user_info, 0)
	for _, item := range res.UserList {
		UserList = append(UserList, types.Douyin_user_info{
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
		})
	}

	return &types.Douyin_relation_follow_list_response{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
		UserList:   UserList,
	}, nil
}
