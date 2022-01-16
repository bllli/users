package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"users/users/internal/logic"
	"users/users/internal/svc"
	"users/users/internal/types"
)

func WechatAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WechatAuthRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewWechatAuthLogic(r.Context(), svcCtx)
		resp, err := l.WechatAuth(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
