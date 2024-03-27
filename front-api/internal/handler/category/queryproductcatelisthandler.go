package category

import (
	"net/http"

	"github.com/feihua/zero-admin/front-api/internal/logic/category"
	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryProductCateListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewQueryProductCateListLogic(r.Context(), svcCtx)
		resp, err := l.QueryProductCateList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
