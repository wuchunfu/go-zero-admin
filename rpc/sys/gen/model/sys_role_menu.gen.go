// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysRoleMenu = "sys_role_menu"

// SysRoleMenu 角色菜单关联表
type SysRoleMenu struct {
	ID     int64 `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"` // 编号
	RoleID int64 `gorm:"column:role_id;not null;comment:角色Id" json:"role_id"`          // 角色Id
	MenuID int64 `gorm:"column:menu_id;not null;comment:菜单Id" json:"menu_id"`          // 菜单Id
}

// TableName SysRoleMenu's table name
func (*SysRoleMenu) TableName() string {
	return TableNameSysRoleMenu
}
