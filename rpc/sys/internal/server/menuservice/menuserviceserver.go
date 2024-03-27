// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/internal/logic/menuservice"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
)

type MenuServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedMenuServiceServer
}

func NewMenuServiceServer(svcCtx *svc.ServiceContext) *MenuServiceServer {
	return &MenuServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *MenuServiceServer) MenuAdd(ctx context.Context, in *sysclient.MenuAddReq) (*sysclient.MenuAddResp, error) {
	l := menuservicelogic.NewMenuAddLogic(ctx, s.svcCtx)
	return l.MenuAdd(in)
}

func (s *MenuServiceServer) MenuList(ctx context.Context, in *sysclient.MenuListReq) (*sysclient.MenuListResp, error) {
	l := menuservicelogic.NewMenuListLogic(ctx, s.svcCtx)
	return l.MenuList(in)
}

func (s *MenuServiceServer) MenuUpdate(ctx context.Context, in *sysclient.MenuUpdateReq) (*sysclient.MenuUpdateResp, error) {
	l := menuservicelogic.NewMenuUpdateLogic(ctx, s.svcCtx)
	return l.MenuUpdate(in)
}

func (s *MenuServiceServer) MenuDelete(ctx context.Context, in *sysclient.MenuDeleteReq) (*sysclient.MenuDeleteResp, error) {
	l := menuservicelogic.NewMenuDeleteLogic(ctx, s.svcCtx)
	return l.MenuDelete(in)
}
