package memberreadhistoryservicelogic

import (
	"context"
	"database/sql"
	"github.com/feihua/zero-admin/rpc/model/umsmodel"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberReadHistoryAddLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 16:43
*/
type MemberReadHistoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReadHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReadHistoryAddLogic {
	return &MemberReadHistoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberReadHistoryAdd 添加浏览商品记录
func (l *MemberReadHistoryAddLogic) MemberReadHistoryAdd(in *umsclient.MemberReadHistoryAddReq) (*umsclient.MemberReadHistoryAddResp, error) {
	_, err := l.svcCtx.UmsMemberReadHistoryModel.Insert(l.ctx, &umsmodel.UmsMemberReadHistory{
		MemberId:        in.MemberId,
		MemberNickName:  in.MemberNickName,
		MemberIcon:      in.MemberIcon,
		ProductId:       in.ProductId,
		ProductName:     in.ProductName,
		ProductPic:      in.ProductPic,
		ProductSubTitle: sql.NullString{String: in.ProductSubTitle, Valid: true},
		ProductPrice:    in.ProductPrice,
		CreateTime:      time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberReadHistoryAddResp{}, nil
}
