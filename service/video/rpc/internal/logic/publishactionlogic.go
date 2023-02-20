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
	"strconv"
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
	// todo: add your logic here and delete this line
	token, err := strconv.ParseInt(in.Token, 10, 64)
	if err != nil {
		return nil, err
	}
	userid, err := jwtx.ParseToken2Uid("a", uint64(token))
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
		CoverUrl:      "",
	}
	_, err = l.svcCtx.VideoModel.Insert(l.ctx, &video_t)
	if err != nil {
		return nil, err
	}
	return &video.DouyinPublishActionResponse{
		StatusCode: int32(errno.OK.Code),
		StatusMsg:  &errno.OK.Message,
	}, nil
}
