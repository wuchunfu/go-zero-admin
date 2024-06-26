// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/logic/flashpromotionservice"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
)

type FlashPromotionServiceServer struct {
	svcCtx *svc.ServiceContext
	smsclient.UnimplementedFlashPromotionServiceServer
}

func NewFlashPromotionServiceServer(svcCtx *svc.ServiceContext) *FlashPromotionServiceServer {
	return &FlashPromotionServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加限时购表
func (s *FlashPromotionServiceServer) AddFlashPromotion(ctx context.Context, in *smsclient.AddFlashPromotionReq) (*smsclient.AddFlashPromotionResp, error) {
	l := flashpromotionservicelogic.NewAddFlashPromotionLogic(ctx, s.svcCtx)
	return l.AddFlashPromotion(in)
}

// 删除限时购表
func (s *FlashPromotionServiceServer) DeleteFlashPromotion(ctx context.Context, in *smsclient.DeleteFlashPromotionReq) (*smsclient.DeleteFlashPromotionResp, error) {
	l := flashpromotionservicelogic.NewDeleteFlashPromotionLogic(ctx, s.svcCtx)
	return l.DeleteFlashPromotion(in)
}

// 更新限时购表
func (s *FlashPromotionServiceServer) UpdateFlashPromotion(ctx context.Context, in *smsclient.UpdateFlashPromotionReq) (*smsclient.UpdateFlashPromotionResp, error) {
	l := flashpromotionservicelogic.NewUpdateFlashPromotionLogic(ctx, s.svcCtx)
	return l.UpdateFlashPromotion(in)
}

// 更新限时购表状态
func (s *FlashPromotionServiceServer) UpdateFlashPromotionStatus(ctx context.Context, in *smsclient.UpdateFlashPromotionStatusReq) (*smsclient.UpdateFlashPromotionStatusResp, error) {
	l := flashpromotionservicelogic.NewUpdateFlashPromotionStatusLogic(ctx, s.svcCtx)
	return l.UpdateFlashPromotionStatus(in)
}

// 查询限时购表详情
func (s *FlashPromotionServiceServer) QueryFlashPromotionDetail(ctx context.Context, in *smsclient.QueryFlashPromotionDetailReq) (*smsclient.QueryFlashPromotionDetailResp, error) {
	l := flashpromotionservicelogic.NewQueryFlashPromotionDetailLogic(ctx, s.svcCtx)
	return l.QueryFlashPromotionDetail(in)
}

// 查询限时购表列表
func (s *FlashPromotionServiceServer) QueryFlashPromotionList(ctx context.Context, in *smsclient.QueryFlashPromotionListReq) (*smsclient.QueryFlashPromotionListResp, error) {
	l := flashpromotionservicelogic.NewQueryFlashPromotionListLogic(ctx, s.svcCtx)
	return l.QueryFlashPromotionList(in)
}

// 查询当前时间是否有秒杀活动
func (s *FlashPromotionServiceServer) QueryFlashPromotionListByDate(ctx context.Context, in *smsclient.QueryFlashPromotionListByDateReq) (*smsclient.QueryFlashPromotionListByDateResp, error) {
	l := flashpromotionservicelogic.NewQueryFlashPromotionListByDateLogic(ctx, s.svcCtx)
	return l.QueryFlashPromotionListByDate(in)
}
