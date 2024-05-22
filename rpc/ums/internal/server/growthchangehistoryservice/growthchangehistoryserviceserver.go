// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/logic/growthchangehistoryservice"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
)

type GrowthChangeHistoryServiceServer struct {
	svcCtx *svc.ServiceContext
	umsclient.UnimplementedGrowthChangeHistoryServiceServer
}

func NewGrowthChangeHistoryServiceServer(svcCtx *svc.ServiceContext) *GrowthChangeHistoryServiceServer {
	return &GrowthChangeHistoryServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加成长值变化历史记录
func (s *GrowthChangeHistoryServiceServer) GrowthChangeHistoryAdd(ctx context.Context, in *umsclient.GrowthChangeHistoryAddReq) (*umsclient.GrowthChangeHistoryAddResp, error) {
	l := growthchangehistoryservicelogic.NewGrowthChangeHistoryAddLogic(ctx, s.svcCtx)
	return l.GrowthChangeHistoryAdd(in)
}

// 查询成长值变化历史记录列表
func (s *GrowthChangeHistoryServiceServer) GrowthChangeHistoryList(ctx context.Context, in *umsclient.GrowthChangeHistoryListReq) (*umsclient.GrowthChangeHistoryListResp, error) {
	l := growthchangehistoryservicelogic.NewGrowthChangeHistoryListLogic(ctx, s.svcCtx)
	return l.GrowthChangeHistoryList(in)
}
