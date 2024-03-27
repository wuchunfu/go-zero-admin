// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/internal/logic/syslogservice"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
)

type SysLogServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedSysLogServiceServer
}

func NewSysLogServiceServer(svcCtx *svc.ServiceContext) *SysLogServiceServer {
	return &SysLogServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *SysLogServiceServer) SysLogAdd(ctx context.Context, in *sysclient.SysLogAddReq) (*sysclient.SysLogAddResp, error) {
	l := syslogservicelogic.NewSysLogAddLogic(ctx, s.svcCtx)
	return l.SysLogAdd(in)
}

func (s *SysLogServiceServer) SysLogList(ctx context.Context, in *sysclient.SysLogListReq) (*sysclient.SysLogListResp, error) {
	l := syslogservicelogic.NewSysLogListLogic(ctx, s.svcCtx)
	return l.SysLogList(in)
}

func (s *SysLogServiceServer) SysLogDelete(ctx context.Context, in *sysclient.SysLogDeleteReq) (*sysclient.SysLogDeleteResp, error) {
	l := syslogservicelogic.NewSysLogDeleteLogic(ctx, s.svcCtx)
	return l.SysLogDelete(in)
}
