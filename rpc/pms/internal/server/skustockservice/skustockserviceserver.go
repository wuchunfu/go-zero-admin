// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package server

import (
	"context"

	"zero-admin/rpc/pms/internal/logic/skustockservice"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"
)

type SkuStockServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedSkuStockServiceServer
}

func NewSkuStockServiceServer(svcCtx *svc.ServiceContext) *SkuStockServiceServer {
	return &SkuStockServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *SkuStockServiceServer) SkuStockAdd(ctx context.Context, in *pmsclient.SkuStockAddReq) (*pmsclient.SkuStockAddResp, error) {
	l := skustockservicelogic.NewSkuStockAddLogic(ctx, s.svcCtx)
	return l.SkuStockAdd(in)
}

func (s *SkuStockServiceServer) SkuStockList(ctx context.Context, in *pmsclient.SkuStockListReq) (*pmsclient.SkuStockListResp, error) {
	l := skustockservicelogic.NewSkuStockListLogic(ctx, s.svcCtx)
	return l.SkuStockList(in)
}

func (s *SkuStockServiceServer) SkuStockUpdate(ctx context.Context, in *pmsclient.SkuStockUpdateReq) (*pmsclient.SkuStockUpdateResp, error) {
	l := skustockservicelogic.NewSkuStockUpdateLogic(ctx, s.svcCtx)
	return l.SkuStockUpdate(in)
}

func (s *SkuStockServiceServer) SkuStockDelete(ctx context.Context, in *pmsclient.SkuStockDeleteReq) (*pmsclient.SkuStockDeleteResp, error) {
	l := skustockservicelogic.NewSkuStockDeleteLogic(ctx, s.svcCtx)
	return l.SkuStockDelete(in)
}

// 取消订单的时候,释放库存
func (s *SkuStockServiceServer) ReleaseSkuStockLock(ctx context.Context, in *pmsclient.ReleaseSkuStockLockReq) (*pmsclient.ReleaseSkuStockLockResp, error) {
	l := skustockservicelogic.NewReleaseSkuStockLockLogic(ctx, s.svcCtx)
	return l.ReleaseSkuStockLock(in)
}

// 下单的时候,锁定库存
func (s *SkuStockServiceServer) LockSkuStockLock(ctx context.Context, in *pmsclient.LockSkuStockLockReq) (*pmsclient.LockSkuStockLockResp, error) {
	l := skustockservicelogic.NewLockSkuStockLockLogic(ctx, s.svcCtx)
	return l.LockSkuStockLock(in)
}

// 根据ProductSkuId查询sku
func (s *SkuStockServiceServer) QuerySkuStockByProductSkuId(ctx context.Context, in *pmsclient.QuerySkuStockByProductSkuIdReq) (*pmsclient.SkuStockListData, error) {
	l := skustockservicelogic.NewQuerySkuStockByProductSkuIdLogic(ctx, s.svcCtx)
	return l.QuerySkuStockByProductSkuId(in)
}
