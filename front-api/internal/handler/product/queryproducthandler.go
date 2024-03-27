package product

import (
	"net/http"

	"github.com/feihua/zero-admin/front-api/internal/logic/product"
	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewQueryProductLogic(r.Context(), svcCtx)
		resp, err := l.QueryProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
