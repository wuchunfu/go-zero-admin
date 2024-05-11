// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCmsTopicCategory = "cms_topic_category"

// CmsTopicCategory 话题分类表
type CmsTopicCategory struct {
	ID           int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name         string     `gorm:"column:name;not null" json:"name"`
	Icon         string     `gorm:"column:icon;not null;comment:分类图标" json:"icon"`                   // 分类图标
	SubjectCount int32      `gorm:"column:subject_count;not null;comment:专题数量" json:"subject_count"` // 专题数量
	ShowStatus   int32      `gorm:"column:show_status;not null" json:"show_status"`
	Sort         int32      `gorm:"column:sort;not null" json:"sort"`
	CreateBy     string     `gorm:"column:create_by;not null;comment:创建者" json:"create_by"`                                // 创建者
	CreateTime   time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateBy     *string    `gorm:"column:update_by;comment:更新者" json:"update_by"`                                         // 更新者
	UpdateTime   *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                    // 更新时间
}

// TableName CmsTopicCategory's table name
func (*CmsTopicCategory) TableName() string {
	return TableNameCmsTopicCategory
}