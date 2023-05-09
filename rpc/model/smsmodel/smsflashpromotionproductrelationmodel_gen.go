// Code generated by goctl. DO NOT EDIT.

package smsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	smsFlashPromotionProductRelationFieldNames          = builder.RawFieldNames(&SmsFlashPromotionProductRelation{})
	smsFlashPromotionProductRelationRows                = strings.Join(smsFlashPromotionProductRelationFieldNames, ",")
	smsFlashPromotionProductRelationRowsExpectAutoSet   = strings.Join(stringx.Remove(smsFlashPromotionProductRelationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	smsFlashPromotionProductRelationRowsWithPlaceHolder = strings.Join(stringx.Remove(smsFlashPromotionProductRelationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	smsFlashPromotionProductRelationModel interface {
		Insert(ctx context.Context, data *SmsFlashPromotionProductRelation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SmsFlashPromotionProductRelation, error)
		Update(ctx context.Context, data *SmsFlashPromotionProductRelation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSmsFlashPromotionProductRelationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SmsFlashPromotionProductRelation struct {
		Id                      int64   `db:"id"`                         // 编号
		FlashPromotionId        int64   `db:"flash_promotion_id"`         // 限时购id
		FlashPromotionSessionId int64   `db:"flash_promotion_session_id"` // 编号
		ProductId               int64   `db:"product_id"`                 // 商品id
		FlashPromotionPrice     float64 `db:"flash_promotion_price"`      // 限时购价格
		FlashPromotionCount     int64   `db:"flash_promotion_count"`      // 限时购数量
		FlashPromotionLimit     int64   `db:"flash_promotion_limit"`      // 每人限购数量
		Sort                    int64   `db:"sort"`                       // 排序
	}
)

func newSmsFlashPromotionProductRelationModel(conn sqlx.SqlConn) *defaultSmsFlashPromotionProductRelationModel {
	return &defaultSmsFlashPromotionProductRelationModel{
		conn:  conn,
		table: "`sms_flash_promotion_product_relation`",
	}
}

func (m *defaultSmsFlashPromotionProductRelationModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSmsFlashPromotionProductRelationModel) FindOne(ctx context.Context, id int64) (*SmsFlashPromotionProductRelation, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", smsFlashPromotionProductRelationRows, m.table)
	var resp SmsFlashPromotionProductRelation
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSmsFlashPromotionProductRelationModel) Insert(ctx context.Context, data *SmsFlashPromotionProductRelation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, smsFlashPromotionProductRelationRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.FlashPromotionId, data.FlashPromotionSessionId, data.ProductId, data.FlashPromotionPrice, data.FlashPromotionCount, data.FlashPromotionLimit, data.Sort)
	return ret, err
}

func (m *defaultSmsFlashPromotionProductRelationModel) Update(ctx context.Context, data *SmsFlashPromotionProductRelation) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, smsFlashPromotionProductRelationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.FlashPromotionId, data.FlashPromotionSessionId, data.ProductId, data.FlashPromotionPrice, data.FlashPromotionCount, data.FlashPromotionLimit, data.Sort, data.Id)
	return err
}

func (m *defaultSmsFlashPromotionProductRelationModel) tableName() string {
	return m.table
}
