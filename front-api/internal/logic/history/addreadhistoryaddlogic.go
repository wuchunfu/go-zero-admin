package history

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"time"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddReadHistoryAddLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 16:29
*/
type AddReadHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddReadHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReadHistoryAddLogic {
	return &AddReadHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddReadHistoryAdd 添加浏览商品记录
func (l *AddReadHistoryAddLogic) AddReadHistoryAdd(req *types.AddReadHistoryReq) (resp *types.AddReadHistoryResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	member, _ := l.svcCtx.MemberService.QueryMemberById(l.ctx, &umsclient.MemberByIdReq{Id: memberId})

	_, err = l.svcCtx.MemberReadHistoryService.MemberReadHistoryAdd(l.ctx, &umsclient.MemberReadHistoryAddReq{
		MemberId:        member.Id,
		MemberNickName:  member.Nickname,
		MemberIcon:      member.Icon,
		ProductId:       req.ProductId,
		ProductName:     req.ProductName,
		ProductPic:      req.ProductPic,
		ProductSubTitle: req.ProductSubTitle,
		ProductPrice:    req.ProductPrice,
		CreateTime:      time.Now().Format("2006-01-02 15:04:05"),
	})

	if err != nil {
		return nil, err
	}
	return &types.AddReadHistoryResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
