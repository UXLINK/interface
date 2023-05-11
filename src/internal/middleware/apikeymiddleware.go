package middleware

import (
	"net/http"
	"uxuy/src/internal/model"
	"uxuy/src/util/ctxdata"
	bizresponse "uxuy/src/util/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type ApiKeyMiddleware struct {
	DappModel *model.DappModel
}

func NewApiKeyMiddleware(dappModel *model.DappModel) *ApiKeyMiddleware {

	return &ApiKeyMiddleware{
		DappModel: dappModel,
	}
}

func (m *ApiKeyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey := r.Header.Get("ApiKey")

		dappEntity, err := m.DappModel.FindDappByApiKey(apiKey)
		if err != nil {
			httpx.OkJson(w, bizresponse.ErrClientCancel)
			return
		}

		ctx := ctxdata.SetDappIToCtx(r.Context(), dappEntity.DappId)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
