package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mini-douyin/service/social/rpc/follow"
	"mini-douyin/service/social/rpc/internal/svc"
	"mini-douyin/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowListLogic) FollowList(in *follow.DouyinRelationFollowListRequest) (*follow.DouyinRelationFollowListResponse, error) {
	// todo: add your logic here and delete this line
	var u []*follow.User
	follows, err := l.svcCtx.FollowModel.FindAllByUserId(l.ctx, in.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "user不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	for _, f := range follows {
		res, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, f.ToUserId)
		if err != nil {
			if err == model.ErrNotFound {
				return nil, status.Error(100, "用户不存在")
			}
			return nil, status.Error(100, "查询用户失败")
		}
		us := &follow.User{
			Id:              res.Id,
			Name:            res.Name,
			FollowCount:     &res.FollowerCount,
			FollowerCount:   &res.FollowerCount,
			IsFollow:        true,
			Avatar:          &res.Avatar.String,
			BackgroundImage: &res.BackgroundImage.String,
			Signature:       &res.Signature.String,
			TotalFavorited:  &res.TotalFavorited.Int64,
			WorkCount:       &res.WorkCount.Int64,
			FavoriteCount:   &res.FavoriteCount.Int64,
		}
		u = append(u, us)

	}

	return &follow.DouyinRelationFollowListResponse{
		UserList: u,
	}, nil
}
