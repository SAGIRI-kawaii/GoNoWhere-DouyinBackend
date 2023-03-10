package logic

import (
	"context"
	"database/sql"
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/core/logx"
	"mini-douyin/common/errno"
	"mini-douyin/common/jwtx"
	"mini-douyin/common/upload"
	"mini-douyin/service/video/model"
	"mini-douyin/service/video/rpc/internal/svc"
	"mini-douyin/service/video/rpc/video"
)

type PublishActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext

	logx.Logger
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	var options = idgen.NewIdGeneratorOptions(1)
	idgen.SetIdGenerator(options)
	return &PublishActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishActionLogic) PublishAction(in *video.DouyinPublishActionRequest) (*video.DouyinPublishActionResponse, error) {
	claims, err := jwtx.ParseToken(in.Token)
	userid := claims.UserID
	if err != nil {
		return nil, err
	}
	videoID := idgen.NextId()
	videoUrl, err := upload.UploadVideo(&in.Data, videoID)
	if err != nil {
		return nil, err
	}
	video_t := model.Videos{
		DeletedAt:     sql.NullTime{},
		VideoId:       videoID,
		AuthorId:      int64(userid),
		Title:         in.Title,
		FavoriteCount: 0,
		CommentCount:  0,
		PlayUrl:       videoUrl,
		CoverUrl:      "http://rqbxff5oo.hn-bkt.clouddn.com/cover/cover.jpeg",
	}
	user, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, int64(userid))
	if err != nil {
		return nil, err
	}
	err1 := l.svcCtx.UserModel.Update(l.ctx, &model.Users{
		Id:              user.Id,
		CreateAt:        user.CreateAt,
		DeletedAt:       user.DeletedAt,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		UserId:          user.UserId,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       sql.NullInt64{Valid: true, Int64: int64(user.WorkCount.Int64 + 1)},
		FavoriteCount:   user.FavoriteCount,
	})
	if err1 != nil {
		return nil, err1
	}
	_, err = l.svcCtx.VideoModel.Insert(l.ctx, &video_t)
	if err != nil {
		upload.DeleteVideo(videoID)
		return nil, err
	}
	return &video.DouyinPublishActionResponse{
		StatusCode: int32(errno.OK.Code),
		StatusMsg:  &errno.OK.Message,
	}, nil
}
