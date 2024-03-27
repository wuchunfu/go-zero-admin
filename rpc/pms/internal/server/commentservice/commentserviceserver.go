// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/logic/commentservice"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
)

type CommentServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedCommentServiceServer
}

func NewCommentServiceServer(svcCtx *svc.ServiceContext) *CommentServiceServer {
	return &CommentServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentServiceServer) CommentAdd(ctx context.Context, in *pmsclient.CommentAddReq) (*pmsclient.CommentAddResp, error) {
	l := commentservicelogic.NewCommentAddLogic(ctx, s.svcCtx)
	return l.CommentAdd(in)
}

func (s *CommentServiceServer) CommentList(ctx context.Context, in *pmsclient.CommentListReq) (*pmsclient.CommentListResp, error) {
	l := commentservicelogic.NewCommentListLogic(ctx, s.svcCtx)
	return l.CommentList(in)
}

func (s *CommentServiceServer) CommentUpdate(ctx context.Context, in *pmsclient.CommentUpdateReq) (*pmsclient.CommentUpdateResp, error) {
	l := commentservicelogic.NewCommentUpdateLogic(ctx, s.svcCtx)
	return l.CommentUpdate(in)
}

func (s *CommentServiceServer) CommentDelete(ctx context.Context, in *pmsclient.CommentDeleteReq) (*pmsclient.CommentDeleteResp, error) {
	l := commentservicelogic.NewCommentDeleteLogic(ctx, s.svcCtx)
	return l.CommentDelete(in)
}
