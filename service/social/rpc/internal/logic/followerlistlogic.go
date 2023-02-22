package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mini-douyin/common/jwtx"
	"mini-douyin/service/social/rpc/follow"
	"mini-douyin/service/social/rpc/internal/svc"
	"mini-douyin/service/user/model/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowerListLogic) FollowerList(in *follow.DouyinRelationFollowerListRequest) (*follow.DouyinRelationFollowerListResponse, error) {
	// todo: add your logic here and delete this line
	_, err := jwtx.ParseToken(in.Token)
	if err != nil {
		return nil, err
	}
	/*if in.UserId != res.UserID {
		return nil, status.Error(100, "非法token")
	}*/

	var u []*follow.User
	follows, err := l.svcCtx.FollowModel.FindAllByToUserId(l.ctx, in.UserId)
	if err != nil {
		if err == users.ErrNotFound {
			return nil, status.Error(100, "user不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	for _, f := range follows {
		var isf bool
		err := l.svcCtx.FollowModel.FindOneById(l.ctx, in.UserId, f.UserId)
		if err == nil {
			isf = true
		} else {
			isf = false
		}

		res, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, f.UserId)
		if err != nil {
			if err == users.ErrNotFound {
				return nil, status.Error(100, "用户不存在")
			}
			return nil, status.Error(100, "查询用户失败")
		}

		us := &follow.User{
			Id:              res.UserId,
			Name:            res.Name,
			FollowCount:     &res.FollowerCount,
			FollowerCount:   &res.FollowerCount,
			IsFollow:        isf,
			Avatar:          &res.Avatar.String,
			BackgroundImage: &res.BackgroundImage.String,
			Signature:       &res.Signature.String,
			TotalFavorited:  &res.TotalFavorited.Int64,
			WorkCount:       &res.WorkCount.Int64,
			FavoriteCount:   &res.FavoriteCount.Int64,
		}
		u = append(u, us)

	}

	return &follow.DouyinRelationFollowerListResponse{
		UserList: u,
	}, nil
}
