// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsProductFullReduction = "pms_product_full_reduction"

// PmsProductFullReduction 产品满减表(只针对同商品)
type PmsProductFullReduction struct {
	ID          int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProductID   int64 `gorm:"column:product_id;not null" json:"product_id"`
	FullPrice   int64 `gorm:"column:full_price;not null" json:"full_price"`
	ReducePrice int64 `gorm:"column:reduce_price;not null" json:"reduce_price"`
}

// TableName PmsProductFullReduction's table name
func (*PmsProductFullReduction) TableName() string {
	return TableNamePmsProductFullReduction
}
