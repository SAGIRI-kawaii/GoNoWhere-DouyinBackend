package logic

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
	"mini-douyin/service/video/rpc/videoservice"
)

const (
	defaultMultipartMemory = 32 << 20 // 32 MB
)

type PublishActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func FromFile(r *http.Request, name string) ([]byte, error) {
	if r.MultipartForm == nil {
		if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, _, err := r.FormFile(name)
	if err != nil {
		return nil, err
	}
	byteContainer, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	err = f.Close()
	if err != nil {
		return nil, err
	}
	return byteContainer, err
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishActionLogic) PublishAction(req *http.Request) (resp *types.Douyin_publish_action_response, err error) {
	data, err := FromFile(req, "data")
	token := req.FormValue("token")
	title := req.FormValue("title")
	res, err := l.svcCtx.VideoRpc.PublishAction(l.ctx, &videoservice.DouyinPublishActionRequest{
		Token: token,
		Data:  data,
		Title: title,
	})
	if err != nil {
		return nil, err
	}
	return &types.Douyin_publish_action_response{
		StatusCode: res.StatusCode,
		StatusMsg:  *res.StatusMsg,
	}, nil
}
