// Code generated by goctl. DO NOT EDIT.
// Source: cms.proto

package subjectproductrelationservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PrefrenceAreaAddReq                    = cmsclient.PrefrenceAreaAddReq
	PrefrenceAreaAddResp                   = cmsclient.PrefrenceAreaAddResp
	PrefrenceAreaDeleteReq                 = cmsclient.PrefrenceAreaDeleteReq
	PrefrenceAreaDeleteResp                = cmsclient.PrefrenceAreaDeleteResp
	PrefrenceAreaListData                  = cmsclient.PrefrenceAreaListData
	PrefrenceAreaListReq                   = cmsclient.PrefrenceAreaListReq
	PrefrenceAreaListResp                  = cmsclient.PrefrenceAreaListResp
	PrefrenceAreaProductRelationAddReq     = cmsclient.PrefrenceAreaProductRelationAddReq
	PrefrenceAreaProductRelationAddResp    = cmsclient.PrefrenceAreaProductRelationAddResp
	PrefrenceAreaProductRelationDeleteReq  = cmsclient.PrefrenceAreaProductRelationDeleteReq
	PrefrenceAreaProductRelationDeleteResp = cmsclient.PrefrenceAreaProductRelationDeleteResp
	PrefrenceAreaProductRelationListReq    = cmsclient.PrefrenceAreaProductRelationListReq
	PrefrenceAreaProductRelationListResp   = cmsclient.PrefrenceAreaProductRelationListResp
	PrefrenceAreaProductRelationUpdateReq  = cmsclient.PrefrenceAreaProductRelationUpdateReq
	PrefrenceAreaProductRelationUpdateResp = cmsclient.PrefrenceAreaProductRelationUpdateResp
	PrefrenceAreaUpdateReq                 = cmsclient.PrefrenceAreaUpdateReq
	PrefrenceAreaUpdateResp                = cmsclient.PrefrenceAreaUpdateResp
	SubjectAddReq                          = cmsclient.SubjectAddReq
	SubjectAddResp                         = cmsclient.SubjectAddResp
	SubjectDeleteReq                       = cmsclient.SubjectDeleteReq
	SubjectDeleteResp                      = cmsclient.SubjectDeleteResp
	SubjectListByIdsReq                    = cmsclient.SubjectListByIdsReq
	SubjectListData                        = cmsclient.SubjectListData
	SubjectListReq                         = cmsclient.SubjectListReq
	SubjectListResp                        = cmsclient.SubjectListResp
	SubjectProductRelationAddReq           = cmsclient.SubjectProductRelationAddReq
	SubjectProductRelationAddResp          = cmsclient.SubjectProductRelationAddResp
	SubjectProductRelationDeleteReq        = cmsclient.SubjectProductRelationDeleteReq
	SubjectProductRelationDeleteResp       = cmsclient.SubjectProductRelationDeleteResp
	SubjectProductRelationListReq          = cmsclient.SubjectProductRelationListReq
	SubjectProductRelationListResp         = cmsclient.SubjectProductRelationListResp
	SubjectProductRelationUpdateReq        = cmsclient.SubjectProductRelationUpdateReq
	SubjectProductRelationUpdateResp       = cmsclient.SubjectProductRelationUpdateResp
	SubjectUpdateReq                       = cmsclient.SubjectUpdateReq
	SubjectUpdateResp                      = cmsclient.SubjectUpdateResp

	SubjectProductRelationService interface {
		// 专题关联
		SubjectProductRelationAdd(ctx context.Context, in *SubjectProductRelationAddReq, opts ...grpc.CallOption) (*SubjectProductRelationAddResp, error)
		SubjectProductRelationDelete(ctx context.Context, in *SubjectProductRelationDeleteReq, opts ...grpc.CallOption) (*SubjectProductRelationDeleteResp, error)
		SubjectProductRelationUpdate(ctx context.Context, in *SubjectProductRelationUpdateReq, opts ...grpc.CallOption) (*SubjectProductRelationUpdateResp, error)
		SubjectProductRelationList(ctx context.Context, in *SubjectProductRelationListReq, opts ...grpc.CallOption) (*SubjectProductRelationListResp, error)
	}

	defaultSubjectProductRelationService struct {
		cli zrpc.Client
	}
)

func NewSubjectProductRelationService(cli zrpc.Client) SubjectProductRelationService {
	return &defaultSubjectProductRelationService{
		cli: cli,
	}
}

// 专题关联
func (m *defaultSubjectProductRelationService) SubjectProductRelationAdd(ctx context.Context, in *SubjectProductRelationAddReq, opts ...grpc.CallOption) (*SubjectProductRelationAddResp, error) {
	client := cmsclient.NewSubjectProductRelationServiceClient(m.cli.Conn())
	return client.SubjectProductRelationAdd(ctx, in, opts...)
}

func (m *defaultSubjectProductRelationService) SubjectProductRelationDelete(ctx context.Context, in *SubjectProductRelationDeleteReq, opts ...grpc.CallOption) (*SubjectProductRelationDeleteResp, error) {
	client := cmsclient.NewSubjectProductRelationServiceClient(m.cli.Conn())
	return client.SubjectProductRelationDelete(ctx, in, opts...)
}

func (m *defaultSubjectProductRelationService) SubjectProductRelationUpdate(ctx context.Context, in *SubjectProductRelationUpdateReq, opts ...grpc.CallOption) (*SubjectProductRelationUpdateResp, error) {
	client := cmsclient.NewSubjectProductRelationServiceClient(m.cli.Conn())
	return client.SubjectProductRelationUpdate(ctx, in, opts...)
}

func (m *defaultSubjectProductRelationService) SubjectProductRelationList(ctx context.Context, in *SubjectProductRelationListReq, opts ...grpc.CallOption) (*SubjectProductRelationListResp, error) {
	client := cmsclient.NewSubjectProductRelationServiceClient(m.cli.Conn())
	return client.SubjectProductRelationList(ctx, in, opts...)
}
