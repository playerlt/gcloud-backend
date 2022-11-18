package public_repo

import (
	"gcloud/core/internal/logic/public_repo"
	"net/http"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublicFileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublicFileListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := public_repo.NewPublicFileListLogic(r.Context(), svcCtx)
		resp, err := l.PublicFileList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
