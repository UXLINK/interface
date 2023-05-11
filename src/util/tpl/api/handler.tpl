package {{.PkgName}}

import (
	"net/http"

	{{.ImportPackages}}
	bizresponse "uxuy/src/util/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, bizresponse.ErrInvalidArgs.WithMessage(err.Error()))
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, bizresponse.NewSuccessResp({{if .HasResp}}resp{{else}}nil{{end}}))
		}
	}
}
