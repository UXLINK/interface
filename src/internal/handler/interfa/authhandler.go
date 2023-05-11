package interfa

import (
	"net/http"

	"uxuy/src/internal/logic/interfa"
	"uxuy/src/internal/svc"
	bizresponse "uxuy/src/util/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := interfa.NewAuthLogic(r.Context(), svcCtx)
		resp, err := l.Auth()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, bizresponse.NewSuccessResp(resp))
		}
	}
}
