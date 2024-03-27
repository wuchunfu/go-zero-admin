// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/logic/memberrulesettingservice"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
)

type MemberRuleSettingServiceServer struct {
	svcCtx *svc.ServiceContext
	umsclient.UnimplementedMemberRuleSettingServiceServer
}

func NewMemberRuleSettingServiceServer(svcCtx *svc.ServiceContext) *MemberRuleSettingServiceServer {
	return &MemberRuleSettingServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *MemberRuleSettingServiceServer) MemberRuleSettingAdd(ctx context.Context, in *umsclient.MemberRuleSettingAddReq) (*umsclient.MemberRuleSettingAddResp, error) {
	l := memberrulesettingservicelogic.NewMemberRuleSettingAddLogic(ctx, s.svcCtx)
	return l.MemberRuleSettingAdd(in)
}

func (s *MemberRuleSettingServiceServer) MemberRuleSettingList(ctx context.Context, in *umsclient.MemberRuleSettingListReq) (*umsclient.MemberRuleSettingListResp, error) {
	l := memberrulesettingservicelogic.NewMemberRuleSettingListLogic(ctx, s.svcCtx)
	return l.MemberRuleSettingList(in)
}

func (s *MemberRuleSettingServiceServer) MemberRuleSettingUpdate(ctx context.Context, in *umsclient.MemberRuleSettingUpdateReq) (*umsclient.MemberRuleSettingUpdateResp, error) {
	l := memberrulesettingservicelogic.NewMemberRuleSettingUpdateLogic(ctx, s.svcCtx)
	return l.MemberRuleSettingUpdate(in)
}

func (s *MemberRuleSettingServiceServer) MemberRuleSettingDelete(ctx context.Context, in *umsclient.MemberRuleSettingDeleteReq) (*umsclient.MemberRuleSettingDeleteResp, error) {
	l := memberrulesettingservicelogic.NewMemberRuleSettingDeleteLogic(ctx, s.svcCtx)
	return l.MemberRuleSettingDelete(in)
}
