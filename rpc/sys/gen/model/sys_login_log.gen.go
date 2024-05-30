// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysLoginLog = "sys_login_log"

// SysLoginLog 系统登录日志表
type SysLoginLog struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`                        // 编号
	UserName    string    `gorm:"column:user_name;not null;comment:用户名" json:"user_name"`                              // 用户名
	LoginStatus string    `gorm:"column:login_status;not null;comment:登录状态" json:"login_status"`                       // 登录状态
	IPAddress   string    `gorm:"column:ip_address;not null;comment:IP地址" json:"ip_address"`                           // IP地址
	Browser     string    `gorm:"column:browser;not null;comment:浏览器" json:"browser"`                                  // 浏览器
	Os          string    `gorm:"column:os;not null;comment:操作信息" json:"os"`                                           // 操作信息
	LoginTime   time.Time `gorm:"column:login_time;not null;default:CURRENT_TIMESTAMP;comment:登录时间" json:"login_time"` // 登录时间
}

// TableName SysLoginLog's table name
func (*SysLoginLog) TableName() string {
	return TableNameSysLoginLog
}
