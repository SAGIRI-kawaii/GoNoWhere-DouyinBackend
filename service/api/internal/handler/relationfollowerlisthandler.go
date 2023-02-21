package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mini-douyin/service/api/internal/logic"
	"mini-douyin/service/api/internal/svc"
	"mini-douyin/service/api/internal/types"
)

func RelationFollowerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Douyin_relation_follower_list_request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRelationFollowerListLogic(r.Context(), svcCtx)
		resp, err := l.RelationFollowerList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
