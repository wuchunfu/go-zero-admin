// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/internal/logic/jobservice"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
)

type JobServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedJobServiceServer
}

func NewJobServiceServer(svcCtx *svc.ServiceContext) *JobServiceServer {
	return &JobServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *JobServiceServer) JobAdd(ctx context.Context, in *sysclient.JobAddReq) (*sysclient.JobAddResp, error) {
	l := jobservicelogic.NewJobAddLogic(ctx, s.svcCtx)
	return l.JobAdd(in)
}

func (s *JobServiceServer) JobList(ctx context.Context, in *sysclient.JobListReq) (*sysclient.JobListResp, error) {
	l := jobservicelogic.NewJobListLogic(ctx, s.svcCtx)
	return l.JobList(in)
}

func (s *JobServiceServer) JobUpdate(ctx context.Context, in *sysclient.JobUpdateReq) (*sysclient.JobUpdateResp, error) {
	l := jobservicelogic.NewJobUpdateLogic(ctx, s.svcCtx)
	return l.JobUpdate(in)
}

func (s *JobServiceServer) JobDelete(ctx context.Context, in *sysclient.JobDeleteReq) (*sysclient.JobDeleteResp, error) {
	l := jobservicelogic.NewJobDeleteLogic(ctx, s.svcCtx)
	return l.JobDelete(in)
}
