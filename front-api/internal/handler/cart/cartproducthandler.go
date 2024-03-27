package cart

import (
	"net/http"

	"github.com/feihua/zero-admin/front-api/internal/logic/cart"
	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CartProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cart.NewCartProductLogic(r.Context(), svcCtx)
		resp, err := l.CartProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
