package attention

import (
	"net/http"

	"github.com/feihua/zero-admin/front-api/internal/logic/member/attention"
	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AttentionAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddAttentionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := attention.NewAttentionAddLogic(r.Context(), svcCtx)
		resp, err := l.AttentionAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
