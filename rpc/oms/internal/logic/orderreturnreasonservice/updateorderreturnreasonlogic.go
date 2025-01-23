package orderreturnreasonservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderReturnReasonLogic 更新退货原因表
type UpdateOrderReturnReasonLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderReturnReasonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderReturnReasonLogic {
	return &UpdateOrderReturnReasonLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderReturnReason 更新退货原因表
func (l *UpdateOrderReturnReasonLogic) UpdateOrderReturnReason(in *omsclient.UpdateOrderReturnReasonReq) (*omsclient.UpdateOrderReturnReasonResp, error) {
	q := query.OmsOrderReturnReason
	_, err := q.WithContext(l.ctx).Updates(&model.OmsOrderReturnReason{
		ID:     in.Id,     //
		Name:   in.Name,   // 退货类型
		Sort:   in.Sort,   // 排序
		Status: in.Status, // 状态：0->不启用；1->启用
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateOrderReturnReasonResp{}, nil
}
