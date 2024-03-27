// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	brand "github.com/feihua/zero-admin/front-api/internal/handler/brand"
	cart "github.com/feihua/zero-admin/front-api/internal/handler/cart"
	category "github.com/feihua/zero-admin/front-api/internal/handler/category"
	collection "github.com/feihua/zero-admin/front-api/internal/handler/collection"
	history "github.com/feihua/zero-admin/front-api/internal/handler/history"
	home "github.com/feihua/zero-admin/front-api/internal/handler/home"
	memberaddress "github.com/feihua/zero-admin/front-api/internal/handler/member/address"
	memberattention "github.com/feihua/zero-admin/front-api/internal/handler/member/attention"
	membercoupon "github.com/feihua/zero-admin/front-api/internal/handler/member/coupon"
	membermember "github.com/feihua/zero-admin/front-api/internal/handler/member/member"
	order "github.com/feihua/zero-admin/front-api/internal/handler/order"
	product "github.com/feihua/zero-admin/front-api/internal/handler/product"
	"github.com/feihua/zero-admin/front-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/detail/:brandId",
				Handler: brand.BrandDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: brand.BrandListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/brand"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: cart.CartItemAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/clear",
				Handler: cart.CarItemClearHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: cart.CartItemDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: cart.CarItemListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/product/:productId",
				Handler: cart.CartProductHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/promotionList",
				Handler: cart.CarItemtListPromotionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateAttr",
				Handler: cart.CartUpdateAttrHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateQuantity",
				Handler: cart.CartUpdateQuantityHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/cart"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/productCateList/:parentId",
				Handler: category.ProductCateListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/queryProductCateList",
				Handler: category.QueryProductCateListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/category"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: collection.AddProductCollectionAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/clear",
				Handler: collection.ProductCollectionClearHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/delete/:id",
				Handler: collection.ProductCollectionDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: collection.ProductCollectionListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/collection"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: history.AddReadHistoryAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/clear",
				Handler: history.ReadHistoryClearHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/delete/:id",
				Handler: history.ReadHistoryDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: history.ReadHistoryListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/history"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/index",
				Handler: home.HomeIndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/recommendBrandList",
				Handler: home.RecommendBrandListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/recommendHotProductList",
				Handler: home.RecommendHotProductListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/recommendNewProductList",
				Handler: home.RecommendNewProductListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/recommendProductList",
				Handler: home.RecommendProductListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/home"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: memberaddress.MemberAddressAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/delete/:id",
				Handler: memberaddress.MemberAddressDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: memberaddress.MemberAddressListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: memberaddress.MemberAddressUpdateHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/member/address"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: memberattention.AttentionAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/clear",
				Handler: memberattention.AttentionClearHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/delete/:id",
				Handler: memberattention.AttentionDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: memberattention.AttentionListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/member/attention"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: membercoupon.CouponAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list/:useStatus",
				Handler: membercoupon.CouponHistoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listByCart/:useStatus",
				Handler: membercoupon.CouponListByCartHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listByProductId/:productId",
				Handler: membercoupon.CouponListByProductIdHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/member/coupon"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: membermember.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: membermember.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/member"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: membermember.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateMember",
				Handler: membermember.UpdateMemberHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updatePassword",
				Handler: membermember.UpdatePasswordHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/member"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/orderPay",
				Handler: order.OrderPayHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/orderPayQuery/:orderId",
				Handler: order.OrderPayQueryHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/pay"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/notify",
				Handler: order.NotifyHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/pay"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/cancelUserOrder/:orderId",
				Handler: order.CancelUserOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/confirmReceiveOrder/:orderId",
				Handler: order.ConfirmReceiveOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/deleteOrder/:orderId",
				Handler: order.DeleteOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/generateConfirmOrder",
				Handler: order.GenerateConfirmOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/generateOrder",
				Handler: order.GenerateOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/orderDetail/:orderId",
				Handler: order.OrderDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/orderList/:status/:current",
				Handler: order.OrderListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/payCallback",
				Handler: order.PayCallbackHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/returnApply",
				Handler: order.ReturnApplyHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/order"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/queryProduct/:id",
				Handler: product.QueryProductHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/queryProductList/:productCategoryId",
				Handler: product.QueryProductListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/product"),
	)
}
