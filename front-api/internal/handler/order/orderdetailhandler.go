package order

import (
	"net/http"

	"github.com/feihua/zero-admin/front-api/internal/logic/order"
	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewOrderDetailLogic(r.Context(), svcCtx)
		resp, err := l.OrderDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
