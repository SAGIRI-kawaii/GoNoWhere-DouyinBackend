package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mini-douyin/service/interact/api/internal/logic"
	"mini-douyin/service/interact/api/internal/svc"
	"mini-douyin/service/interact/api/internal/types"
)

func CommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Douyin_comment_list_request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCommentListLogic(r.Context(), svcCtx)
		resp, err := l.CommentList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
