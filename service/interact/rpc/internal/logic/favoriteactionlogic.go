package logic

import (
	"context"

	"mini-douyin/common/jwtx"
	model "mini-douyin/service/interact/model/favorites"
	"mini-douyin/service/interact/model/videos"
	"mini-douyin/service/interact/rpc/interact"
	"mini-douyin/service/interact/rpc/internal/svc"
	"mini-douyin/service/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *interact.DouyinFavoriteActionRequest) (*interact.DouyinFavoriteActionResponse, error) {
	// todo: add your logic here and delete this line

	// UserId := in.Token
	// String转int64
	// UserId, err := strconv.ParseInt(in.Token, 10, 0)
	claims, err := jwtx.ParseToken(in.Token)
	UserId := claims.UserID
	if err != nil {
		return nil, err
	}

	//	判断   VideoId 是否存在

	var existvideo *videos.Videos
	existvideo, err = l.svcCtx.VideoModel.FindOneByVideoId(l.ctx, in.VideoId)
	if err != nil {
		return nil, err
	}
	if existvideo == nil {
		var videonul string = "要点赞的video不存在"
		return &interact.DouyinFavoriteActionResponse{
			StatusCode: int32(0),
			StatusMsg:  &videonul,
		}, nil
	}

	// action_type：1-点赞，2-取消点赞
	ActionType := in.ActionType
	if ActionType == 1 {
		newFavorite := model.Favorites{
			UserId:  UserId,
			VideoId: in.VideoId,
		}
		l.svcCtx.FavoriteModel.Insert(l.ctx, &newFavorite)

		// 修改被点赞video信息的点赞数
		err = l.svcCtx.VideoModel.AddFavoriteByVideoId(l.ctx, in.VideoId)
		if err != nil {
			return nil, status.Error(100, "数据库操作出错")
		}

		// 修改该用户的点赞数
		err = l.svcCtx.UserModel.AddFavoriteByUserId(l.ctx, UserId)
		if err != nil {
			return nil, status.Error(100, "数据库操作出错")
		}

		// 修改被点赞video 作者的 被点赞数
		c, err := l.svcCtx.VideoRpc.SearchVideo(l.ctx, &video.DouyinSearchRequest{
			VideoId: in.VideoId,
		})
		if err != nil {
			return nil, err
		}
		err = l.svcCtx.UserModel.AddFavoritedByUId(l.ctx, c.Video.Author.Id) //不是uid
		if err != nil {
			return nil, status.Error(100, "数据库操作出错")
		}

		// return
		var a string = "点赞成功"
		return &interact.DouyinFavoriteActionResponse{
			StatusCode: int32(0),
			StatusMsg:  &a,
		}, nil

	} else if ActionType == 2 {
		err := l.svcCtx.FavoriteModel.DeleteByUserId2VideoId(l.ctx, UserId, in.VideoId)
		if err != nil {
			return nil, status.Error(100, "取消点赞失败")
		}
		// 修改被点赞video信息的点赞数
		err = l.svcCtx.VideoModel.ReduceFavoriteByVideoId(l.ctx, in.VideoId)
		if err != nil {
			return nil, status.Error(100, "数据库操作出错")
		}
		// 修改该用户的点赞数
		err = l.svcCtx.UserModel.ReduceFavoriteByUserId(l.ctx, UserId)
		if err != nil {
			return nil, status.Error(100, "数据库操作出错")
		}
		// 修改被点赞video 作者的 被点赞数
		c, err := l.svcCtx.VideoRpc.SearchVideo(l.ctx, &video.DouyinSearchRequest{
			VideoId: in.VideoId,
		})
		if err != nil {
			return nil, err
		}
		err = l.svcCtx.UserModel.ReduceFavoritedByUId(l.ctx, c.Video.Author.Id) //是uid
		if err != nil {
			return nil, status.Error(100, "数据库操作出错")
		}

		// return
		var a string = "取消点赞成功"
		return &interact.DouyinFavoriteActionResponse{
			StatusCode: int32(0),
			StatusMsg:  &a,
		}, nil
	}
	var a string = "ActionType输入异常"
	return &interact.DouyinFavoriteActionResponse{
		StatusCode: int32(0),
		StatusMsg:  &a,
	}, nil

}
