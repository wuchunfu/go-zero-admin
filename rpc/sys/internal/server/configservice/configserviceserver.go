// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/internal/logic/configservice"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
)

type ConfigServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedConfigServiceServer
}

func NewConfigServiceServer(svcCtx *svc.ServiceContext) *ConfigServiceServer {
	return &ConfigServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ConfigServiceServer) ConfigAdd(ctx context.Context, in *sysclient.ConfigAddReq) (*sysclient.ConfigAddResp, error) {
	l := configservicelogic.NewConfigAddLogic(ctx, s.svcCtx)
	return l.ConfigAdd(in)
}

func (s *ConfigServiceServer) ConfigList(ctx context.Context, in *sysclient.ConfigListReq) (*sysclient.ConfigListResp, error) {
	l := configservicelogic.NewConfigListLogic(ctx, s.svcCtx)
	return l.ConfigList(in)
}

func (s *ConfigServiceServer) ConfigUpdate(ctx context.Context, in *sysclient.ConfigUpdateReq) (*sysclient.ConfigUpdateResp, error) {
	l := configservicelogic.NewConfigUpdateLogic(ctx, s.svcCtx)
	return l.ConfigUpdate(in)
}

func (s *ConfigServiceServer) ConfigDelete(ctx context.Context, in *sysclient.ConfigDeleteReq) (*sysclient.ConfigDeleteResp, error) {
	l := configservicelogic.NewConfigDeleteLogic(ctx, s.svcCtx)
	return l.ConfigDelete(in)
}

func (s *ConfigServiceServer) UpdateConfigRole(ctx context.Context, in *sysclient.UpdateConfigRoleReq) (*sysclient.UpdateConfigRoleResp, error) {
	l := configservicelogic.NewUpdateConfigRoleLogic(ctx, s.svcCtx)
	return l.UpdateConfigRole(in)
}
