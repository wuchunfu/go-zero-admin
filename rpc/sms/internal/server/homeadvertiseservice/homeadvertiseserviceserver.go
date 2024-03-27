// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/logic/homeadvertiseservice"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
)

type HomeAdvertiseServiceServer struct {
	svcCtx *svc.ServiceContext
	smsclient.UnimplementedHomeAdvertiseServiceServer
}

func NewHomeAdvertiseServiceServer(svcCtx *svc.ServiceContext) *HomeAdvertiseServiceServer {
	return &HomeAdvertiseServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *HomeAdvertiseServiceServer) HomeAdvertiseAdd(ctx context.Context, in *smsclient.HomeAdvertiseAddReq) (*smsclient.HomeAdvertiseAddResp, error) {
	l := homeadvertiseservicelogic.NewHomeAdvertiseAddLogic(ctx, s.svcCtx)
	return l.HomeAdvertiseAdd(in)
}

func (s *HomeAdvertiseServiceServer) HomeAdvertiseList(ctx context.Context, in *smsclient.HomeAdvertiseListReq) (*smsclient.HomeAdvertiseListResp, error) {
	l := homeadvertiseservicelogic.NewHomeAdvertiseListLogic(ctx, s.svcCtx)
	return l.HomeAdvertiseList(in)
}

func (s *HomeAdvertiseServiceServer) HomeAdvertiseUpdate(ctx context.Context, in *smsclient.HomeAdvertiseUpdateReq) (*smsclient.HomeAdvertiseUpdateResp, error) {
	l := homeadvertiseservicelogic.NewHomeAdvertiseUpdateLogic(ctx, s.svcCtx)
	return l.HomeAdvertiseUpdate(in)
}

func (s *HomeAdvertiseServiceServer) HomeAdvertiseDelete(ctx context.Context, in *smsclient.HomeAdvertiseDeleteReq) (*smsclient.HomeAdvertiseDeleteResp, error) {
	l := homeadvertiseservicelogic.NewHomeAdvertiseDeleteLogic(ctx, s.svcCtx)
	return l.HomeAdvertiseDelete(in)
}
