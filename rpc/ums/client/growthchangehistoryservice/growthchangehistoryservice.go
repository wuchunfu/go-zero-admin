// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package growthchangehistoryservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GrowthChangeHistoryAddReq               = umsclient.GrowthChangeHistoryAddReq
	GrowthChangeHistoryAddResp              = umsclient.GrowthChangeHistoryAddResp
	GrowthChangeHistoryListData             = umsclient.GrowthChangeHistoryListData
	GrowthChangeHistoryListReq              = umsclient.GrowthChangeHistoryListReq
	GrowthChangeHistoryListResp             = umsclient.GrowthChangeHistoryListResp
	IntegrationChangeHistoryAddReq          = umsclient.IntegrationChangeHistoryAddReq
	IntegrationChangeHistoryAddResp         = umsclient.IntegrationChangeHistoryAddResp
	IntegrationChangeHistoryListData        = umsclient.IntegrationChangeHistoryListData
	IntegrationChangeHistoryListReq         = umsclient.IntegrationChangeHistoryListReq
	IntegrationChangeHistoryListResp        = umsclient.IntegrationChangeHistoryListResp
	IntegrationConsumeSettingAddReq         = umsclient.IntegrationConsumeSettingAddReq
	IntegrationConsumeSettingAddResp        = umsclient.IntegrationConsumeSettingAddResp
	IntegrationConsumeSettingDeleteReq      = umsclient.IntegrationConsumeSettingDeleteReq
	IntegrationConsumeSettingDeleteResp     = umsclient.IntegrationConsumeSettingDeleteResp
	IntegrationConsumeSettingListData       = umsclient.IntegrationConsumeSettingListData
	IntegrationConsumeSettingListReq        = umsclient.IntegrationConsumeSettingListReq
	IntegrationConsumeSettingListResp       = umsclient.IntegrationConsumeSettingListResp
	IntegrationConsumeSettingUpdateReq      = umsclient.IntegrationConsumeSettingUpdateReq
	IntegrationConsumeSettingUpdateResp     = umsclient.IntegrationConsumeSettingUpdateResp
	MemberAddReq                            = umsclient.MemberAddReq
	MemberAddResp                           = umsclient.MemberAddResp
	MemberBrandAttentionAddReq              = umsclient.MemberBrandAttentionAddReq
	MemberBrandAttentionAddResp             = umsclient.MemberBrandAttentionAddResp
	MemberBrandAttentionDeleteReq           = umsclient.MemberBrandAttentionDeleteReq
	MemberBrandAttentionDeleteResp          = umsclient.MemberBrandAttentionDeleteResp
	MemberBrandAttentionListData            = umsclient.MemberBrandAttentionListData
	MemberBrandAttentionListReq             = umsclient.MemberBrandAttentionListReq
	MemberBrandAttentionListResp            = umsclient.MemberBrandAttentionListResp
	MemberByIdReq                           = umsclient.MemberByIdReq
	MemberDeleteReq                         = umsclient.MemberDeleteReq
	MemberDeleteResp                        = umsclient.MemberDeleteResp
	MemberLevelAddReq                       = umsclient.MemberLevelAddReq
	MemberLevelAddResp                      = umsclient.MemberLevelAddResp
	MemberLevelDeleteReq                    = umsclient.MemberLevelDeleteReq
	MemberLevelDeleteResp                   = umsclient.MemberLevelDeleteResp
	MemberLevelListData                     = umsclient.MemberLevelListData
	MemberLevelListReq                      = umsclient.MemberLevelListReq
	MemberLevelListResp                     = umsclient.MemberLevelListResp
	MemberLevelUpdateReq                    = umsclient.MemberLevelUpdateReq
	MemberLevelUpdateResp                   = umsclient.MemberLevelUpdateResp
	MemberListData                          = umsclient.MemberListData
	MemberListReq                           = umsclient.MemberListReq
	MemberListResp                          = umsclient.MemberListResp
	MemberLoginLogAddReq                    = umsclient.MemberLoginLogAddReq
	MemberLoginLogAddResp                   = umsclient.MemberLoginLogAddResp
	MemberLoginLogDeleteReq                 = umsclient.MemberLoginLogDeleteReq
	MemberLoginLogDeleteResp                = umsclient.MemberLoginLogDeleteResp
	MemberLoginLogListData                  = umsclient.MemberLoginLogListData
	MemberLoginLogListReq                   = umsclient.MemberLoginLogListReq
	MemberLoginLogListResp                  = umsclient.MemberLoginLogListResp
	MemberLoginReq                          = umsclient.MemberLoginReq
	MemberLoginResp                         = umsclient.MemberLoginResp
	MemberMemberTagRelationAddReq           = umsclient.MemberMemberTagRelationAddReq
	MemberMemberTagRelationAddResp          = umsclient.MemberMemberTagRelationAddResp
	MemberMemberTagRelationDeleteReq        = umsclient.MemberMemberTagRelationDeleteReq
	MemberMemberTagRelationDeleteResp       = umsclient.MemberMemberTagRelationDeleteResp
	MemberMemberTagRelationListData         = umsclient.MemberMemberTagRelationListData
	MemberMemberTagRelationListReq          = umsclient.MemberMemberTagRelationListReq
	MemberMemberTagRelationListResp         = umsclient.MemberMemberTagRelationListResp
	MemberMemberTagRelationUpdateReq        = umsclient.MemberMemberTagRelationUpdateReq
	MemberMemberTagRelationUpdateResp       = umsclient.MemberMemberTagRelationUpdateResp
	MemberProductCategoryRelationAddReq     = umsclient.MemberProductCategoryRelationAddReq
	MemberProductCategoryRelationAddResp    = umsclient.MemberProductCategoryRelationAddResp
	MemberProductCategoryRelationDeleteReq  = umsclient.MemberProductCategoryRelationDeleteReq
	MemberProductCategoryRelationDeleteResp = umsclient.MemberProductCategoryRelationDeleteResp
	MemberProductCategoryRelationListData   = umsclient.MemberProductCategoryRelationListData
	MemberProductCategoryRelationListReq    = umsclient.MemberProductCategoryRelationListReq
	MemberProductCategoryRelationListResp   = umsclient.MemberProductCategoryRelationListResp
	MemberProductCategoryRelationUpdateReq  = umsclient.MemberProductCategoryRelationUpdateReq
	MemberProductCategoryRelationUpdateResp = umsclient.MemberProductCategoryRelationUpdateResp
	MemberProductCollectionAddReq           = umsclient.MemberProductCollectionAddReq
	MemberProductCollectionAddResp          = umsclient.MemberProductCollectionAddResp
	MemberProductCollectionDeleteReq        = umsclient.MemberProductCollectionDeleteReq
	MemberProductCollectionDeleteResp       = umsclient.MemberProductCollectionDeleteResp
	MemberProductCollectionListData         = umsclient.MemberProductCollectionListData
	MemberProductCollectionListReq          = umsclient.MemberProductCollectionListReq
	MemberProductCollectionListResp         = umsclient.MemberProductCollectionListResp
	MemberReadHistoryAddReq                 = umsclient.MemberReadHistoryAddReq
	MemberReadHistoryAddResp                = umsclient.MemberReadHistoryAddResp
	MemberReadHistoryDeleteReq              = umsclient.MemberReadHistoryDeleteReq
	MemberReadHistoryDeleteResp             = umsclient.MemberReadHistoryDeleteResp
	MemberReadHistoryListData               = umsclient.MemberReadHistoryListData
	MemberReadHistoryListReq                = umsclient.MemberReadHistoryListReq
	MemberReadHistoryListResp               = umsclient.MemberReadHistoryListResp
	MemberReceiveAddressAddReq              = umsclient.MemberReceiveAddressAddReq
	MemberReceiveAddressAddResp             = umsclient.MemberReceiveAddressAddResp
	MemberReceiveAddressDeleteReq           = umsclient.MemberReceiveAddressDeleteReq
	MemberReceiveAddressDeleteResp          = umsclient.MemberReceiveAddressDeleteResp
	MemberReceiveAddressListData            = umsclient.MemberReceiveAddressListData
	MemberReceiveAddressListReq             = umsclient.MemberReceiveAddressListReq
	MemberReceiveAddressListResp            = umsclient.MemberReceiveAddressListResp
	MemberReceiveAddressQueryDetailReq      = umsclient.MemberReceiveAddressQueryDetailReq
	MemberReceiveAddressQueryDetailResp     = umsclient.MemberReceiveAddressQueryDetailResp
	MemberReceiveAddressUpdateReq           = umsclient.MemberReceiveAddressUpdateReq
	MemberReceiveAddressUpdateResp          = umsclient.MemberReceiveAddressUpdateResp
	MemberRuleSettingAddReq                 = umsclient.MemberRuleSettingAddReq
	MemberRuleSettingAddResp                = umsclient.MemberRuleSettingAddResp
	MemberRuleSettingDeleteReq              = umsclient.MemberRuleSettingDeleteReq
	MemberRuleSettingDeleteResp             = umsclient.MemberRuleSettingDeleteResp
	MemberRuleSettingListData               = umsclient.MemberRuleSettingListData
	MemberRuleSettingListReq                = umsclient.MemberRuleSettingListReq
	MemberRuleSettingListResp               = umsclient.MemberRuleSettingListResp
	MemberRuleSettingUpdateReq              = umsclient.MemberRuleSettingUpdateReq
	MemberRuleSettingUpdateResp             = umsclient.MemberRuleSettingUpdateResp
	MemberStatisticsInfoAddReq              = umsclient.MemberStatisticsInfoAddReq
	MemberStatisticsInfoAddResp             = umsclient.MemberStatisticsInfoAddResp
	MemberStatisticsInfoDeleteReq           = umsclient.MemberStatisticsInfoDeleteReq
	MemberStatisticsInfoDeleteResp          = umsclient.MemberStatisticsInfoDeleteResp
	MemberStatisticsInfoListData            = umsclient.MemberStatisticsInfoListData
	MemberStatisticsInfoListReq             = umsclient.MemberStatisticsInfoListReq
	MemberStatisticsInfoListResp            = umsclient.MemberStatisticsInfoListResp
	MemberStatisticsInfoUpdateReq           = umsclient.MemberStatisticsInfoUpdateReq
	MemberStatisticsInfoUpdateResp          = umsclient.MemberStatisticsInfoUpdateResp
	MemberTagAddReq                         = umsclient.MemberTagAddReq
	MemberTagAddResp                        = umsclient.MemberTagAddResp
	MemberTagDeleteReq                      = umsclient.MemberTagDeleteReq
	MemberTagDeleteResp                     = umsclient.MemberTagDeleteResp
	MemberTagListData                       = umsclient.MemberTagListData
	MemberTagListReq                        = umsclient.MemberTagListReq
	MemberTagListResp                       = umsclient.MemberTagListResp
	MemberTagUpdateReq                      = umsclient.MemberTagUpdateReq
	MemberTagUpdateResp                     = umsclient.MemberTagUpdateResp
	MemberTaskAddReq                        = umsclient.MemberTaskAddReq
	MemberTaskAddResp                       = umsclient.MemberTaskAddResp
	MemberTaskDeleteReq                     = umsclient.MemberTaskDeleteReq
	MemberTaskDeleteResp                    = umsclient.MemberTaskDeleteResp
	MemberTaskListData                      = umsclient.MemberTaskListData
	MemberTaskListReq                       = umsclient.MemberTaskListReq
	MemberTaskListResp                      = umsclient.MemberTaskListResp
	MemberTaskUpdateReq                     = umsclient.MemberTaskUpdateReq
	MemberTaskUpdateResp                    = umsclient.MemberTaskUpdateResp
	MemberUpdatePasswordReq                 = umsclient.MemberUpdatePasswordReq
	MemberUpdateReq                         = umsclient.MemberUpdateReq
	MemberUpdateResp                        = umsclient.MemberUpdateResp
	QueryIntegrationConsumeSettingByIdReq   = umsclient.QueryIntegrationConsumeSettingByIdReq
	UpdateMemberIntegrationReq              = umsclient.UpdateMemberIntegrationReq
	UpdateMemberIntegrationResp             = umsclient.UpdateMemberIntegrationResp

	GrowthChangeHistoryService interface {
		// 添加成长值变化历史记录
		GrowthChangeHistoryAdd(ctx context.Context, in *GrowthChangeHistoryAddReq, opts ...grpc.CallOption) (*GrowthChangeHistoryAddResp, error)
		// 查询成长值变化历史记录列表
		GrowthChangeHistoryList(ctx context.Context, in *GrowthChangeHistoryListReq, opts ...grpc.CallOption) (*GrowthChangeHistoryListResp, error)
	}

	defaultGrowthChangeHistoryService struct {
		cli zrpc.Client
	}
)

func NewGrowthChangeHistoryService(cli zrpc.Client) GrowthChangeHistoryService {
	return &defaultGrowthChangeHistoryService{
		cli: cli,
	}
}

// 添加成长值变化历史记录
func (m *defaultGrowthChangeHistoryService) GrowthChangeHistoryAdd(ctx context.Context, in *GrowthChangeHistoryAddReq, opts ...grpc.CallOption) (*GrowthChangeHistoryAddResp, error) {
	client := umsclient.NewGrowthChangeHistoryServiceClient(m.cli.Conn())
	return client.GrowthChangeHistoryAdd(ctx, in, opts...)
}

// 查询成长值变化历史记录列表
func (m *defaultGrowthChangeHistoryService) GrowthChangeHistoryList(ctx context.Context, in *GrowthChangeHistoryListReq, opts ...grpc.CallOption) (*GrowthChangeHistoryListResp, error) {
	client := umsclient.NewGrowthChangeHistoryServiceClient(m.cli.Conn())
	return client.GrowthChangeHistoryList(ctx, in, opts...)
}
