package handler

import (
	"net/http"

	"e5Code-Service/service/project/api/internal/logic"
	"e5Code-Service/service/project/api/internal/svc"
	"e5Code-Service/service/project/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func listProjectFilesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListProjectFilesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewListProjectFilesLogic(r.Context(), svcCtx)
		resp, err := l.ListProjectFiles(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
