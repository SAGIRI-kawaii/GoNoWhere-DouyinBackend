package logic

import (
	"context"

	"mini-douyin/service/interact/rpc/interact"
	"mini-douyin/service/interact/rpc/internal/svc"
	"mini-douyin/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *interact.DouyinCommentListRequest) (*interact.DouyinCommentListResponse, error) {
	// todo: add your logic here and delete this line

	// FindList方法查DB返回的list是 Comment  类型
	list, err := l.svcCtx.CommentModel.FindList(l.ctx, in.VideoId)
	if err != nil {
		return nil, err
	}

	// 定义最终返回的结果comment_list,
	var CommentsList []*interact.Comment = make([]*interact.Comment, 0)

	// todo：需要将FindList查到的Comment 填进去,再通过uid查user信息补充User字段
	// a =model.Comments
	for _, item := range list {
		// 通过uid查user信息补充User字段
		res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.DouyinUserRequest{
			Token:  in.Token,
			UserId: item.UserId,
		})
		if err != nil {
			return nil, err //用户查询失败时打印
		}
		// println(res.User.Name)
		var newComment = interact.User{
			Id:              res.User.Id,
			Name:            res.User.Name,
			FollowCount:     res.User.FollowCount,
			FollowerCount:   res.User.FollowerCount,
			IsFollow:        res.User.IsFollow, //待查
			Avatar:          res.User.Avatar,
			BackgroundImage: res.User.BackgroundImage,
			Signature:       res.User.Signature,
			TotalFavorited:  res.User.TotalFavorited,
			WorkCount:       res.User.WorkCount,
			FavoriteCount:   res.User.FavoriteCount,
		}
		var timetostring string = item.CreateAt.Format("01-02")
		// var CommentsList []*interact.DouyinComment =make([]*interact.DouyinComment, 0)
		CommentsList = append(CommentsList, &interact.Comment{
			Id:         item.Id,
			User:       &newComment,
			Content:    item.Content,
			CreateDate: timetostring,
		})
	}

	//   CommentList []*DouyinComment
	var a string = "获取评论列表success"
	return &interact.DouyinCommentListResponse{
		StatusCode:  int32(0),
		StatusMsg:   &a,
		CommentList: CommentsList,
	}, nil
}
