// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/logic/homerecommendsubjectservice"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
)

type HomeRecommendSubjectServiceServer struct {
	svcCtx *svc.ServiceContext
	smsclient.UnimplementedHomeRecommendSubjectServiceServer
}

func NewHomeRecommendSubjectServiceServer(svcCtx *svc.ServiceContext) *HomeRecommendSubjectServiceServer {
	return &HomeRecommendSubjectServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *HomeRecommendSubjectServiceServer) HomeRecommendSubjectAdd(ctx context.Context, in *smsclient.HomeRecommendSubjectAddReq) (*smsclient.HomeRecommendSubjectAddResp, error) {
	l := homerecommendsubjectservicelogic.NewHomeRecommendSubjectAddLogic(ctx, s.svcCtx)
	return l.HomeRecommendSubjectAdd(in)
}

func (s *HomeRecommendSubjectServiceServer) HomeRecommendSubjectList(ctx context.Context, in *smsclient.HomeRecommendSubjectListReq) (*smsclient.HomeRecommendSubjectListResp, error) {
	l := homerecommendsubjectservicelogic.NewHomeRecommendSubjectListLogic(ctx, s.svcCtx)
	return l.HomeRecommendSubjectList(in)
}

func (s *HomeRecommendSubjectServiceServer) HomeRecommendSubjectUpdate(ctx context.Context, in *smsclient.HomeRecommendSubjectUpdateReq) (*smsclient.HomeRecommendSubjectUpdateResp, error) {
	l := homerecommendsubjectservicelogic.NewHomeRecommendSubjectUpdateLogic(ctx, s.svcCtx)
	return l.HomeRecommendSubjectUpdate(in)
}

func (s *HomeRecommendSubjectServiceServer) HomeRecommendSubjectDelete(ctx context.Context, in *smsclient.HomeRecommendSubjectDeleteReq) (*smsclient.HomeRecommendSubjectDeleteResp, error) {
	l := homerecommendsubjectservicelogic.NewHomeRecommendSubjectDeleteLogic(ctx, s.svcCtx)
	return l.HomeRecommendSubjectDelete(in)
}
