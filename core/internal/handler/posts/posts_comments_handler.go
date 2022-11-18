package posts

import (
	"gcloud/core/internal/logic/posts"
	"net/http"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PostsCommentsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostsCommentsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := posts.NewPostsCommentsLogic(r.Context(), svcCtx)
		resp, err := l.PostsComments(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
