// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsFeightTemplate = "pms_feight_template"

// PmsFeightTemplate 运费模版
type PmsFeightTemplate struct {
	ID             int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name           string `gorm:"column:name;not null;comment:运费模版名称" json:"name"`                           // 运费模版名称
	ChargeType     int32  `gorm:"column:charge_type;not null;comment:计费类型:0->按重量；1->按件数" json:"charge_type"` // 计费类型:0->按重量；1->按件数
	FirstWeight    int64  `gorm:"column:first_weight;not null;comment:首重kg" json:"first_weight"`             // 首重kg
	FirstFee       int64  `gorm:"column:first_fee;not null;comment:首费（元）" json:"first_fee"`                  // 首费（元）
	ContinueWeight int64  `gorm:"column:continue_weight;not null" json:"continue_weight"`
	ContinueFee    int64  `gorm:"column:continue_fee;not null" json:"continue_fee"`
	Dest           string `gorm:"column:dest;not null;comment:目的地（省、市）" json:"dest"` // 目的地（省、市）
}

// TableName PmsFeightTemplate's table name
func (*PmsFeightTemplate) TableName() string {
	return TableNamePmsFeightTemplate
}
