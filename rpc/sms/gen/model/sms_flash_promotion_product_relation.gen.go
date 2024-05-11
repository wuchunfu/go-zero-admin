// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSmsFlashPromotionProductRelation = "sms_flash_promotion_product_relation"

// SmsFlashPromotionProductRelation 商品限时购与商品关系表
type SmsFlashPromotionProductRelation struct {
	ID                      int64 `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`                            // 编号
	FlashPromotionID        int64 `gorm:"column:flash_promotion_id;not null;comment:限时购id" json:"flash_promotion_id"`              // 限时购id
	FlashPromotionSessionID int64 `gorm:"column:flash_promotion_session_id;not null;comment:编号" json:"flash_promotion_session_id"` // 编号
	ProductID               int64 `gorm:"column:product_id;not null;comment:商品id" json:"product_id"`                               // 商品id
	FlashPromotionPrice     int64 `gorm:"column:flash_promotion_price;not null;comment:限时购价格" json:"flash_promotion_price"`        // 限时购价格
	FlashPromotionCount     int32 `gorm:"column:flash_promotion_count;not null;comment:限时购数量" json:"flash_promotion_count"`        // 限时购数量
	FlashPromotionLimit     int32 `gorm:"column:flash_promotion_limit;not null;comment:每人限购数量" json:"flash_promotion_limit"`       // 每人限购数量
	Sort                    int32 `gorm:"column:sort;not null;comment:排序" json:"sort"`                                             // 排序
}

// TableName SmsFlashPromotionProductRelation's table name
func (*SmsFlashPromotionProductRelation) TableName() string {
	return TableNameSmsFlashPromotionProductRelation
}
