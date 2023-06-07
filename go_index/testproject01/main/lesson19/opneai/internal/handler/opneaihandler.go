package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"openai-project/opneai/internal/logic"
	"openai-project/opneai/internal/svc"
	"openai-project/opneai/internal/types"
)

func OpneaiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOpneaiLogic(r.Context(), svcCtx)
		resp, err := l.Opneai(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
