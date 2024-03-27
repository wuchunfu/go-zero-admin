package handler

import (
	"net/http"

	"github.com/feihua/zero-admin/api/internal/logic/sms/homerecommendproduct"
	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomeRecommendProductUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateHomeRecommendProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHomeRecommendProductUpdateLogic(r.Context(), ctx)
		resp, err := l.HomeRecommendProductUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
