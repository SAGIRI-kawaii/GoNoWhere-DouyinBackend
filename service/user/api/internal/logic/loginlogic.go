package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mini-douyin/service/user/api/internal/svc"
	"mini-douyin/service/user/api/internal/types"
	"mini-douyin/service/user/rpc/userclient"
	"strconv"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.Douyin_user_login_request) (resp *types.Douyin_user_login_response, err error) {
	// todo: add your logic here and delete this line

	//return &types.Douyin_user_login_response{
	//	User_id: int(1),
	//	Token:   req.Username,
	//}, nil
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.DouyinUserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	//
	//now := time.Now().Unix()
	//accessExpire := l.svcCtx.Config.Auth.AccessExpire
	//
	//accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.UserId)
	//println(res.UserId)
	//println(accessToken)
	if err != nil {
		return nil, err
	}
	//println("token@:" + accessToken)
	return &types.Douyin_user_login_response{
		StatusCode: 0,
		StatusMsg:  "success",
		UserID:     int(res.UserId),
		Token:      strconv.FormatInt(res.UserId, 10),
	}, nil
	return
}
