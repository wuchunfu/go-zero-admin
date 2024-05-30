// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q             = new(Query)
	SysDept       *sysDept
	SysDictItem   *sysDictItem
	SysDictType   *sysDictType
	SysLoginLog   *sysLoginLog
	SysMenu       *sysMenu
	SysOperateLog *sysOperateLog
	SysPost       *sysPost
	SysRole       *sysRole
	SysRoleMenu   *sysRoleMenu
	SysUser       *sysUser
	SysUserRole   *sysUserRole
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	SysDept = &Q.SysDept
	SysDictItem = &Q.SysDictItem
	SysDictType = &Q.SysDictType
	SysLoginLog = &Q.SysLoginLog
	SysMenu = &Q.SysMenu
	SysOperateLog = &Q.SysOperateLog
	SysPost = &Q.SysPost
	SysRole = &Q.SysRole
	SysRoleMenu = &Q.SysRoleMenu
	SysUser = &Q.SysUser
	SysUserRole = &Q.SysUserRole
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:            db,
		SysDept:       newSysDept(db, opts...),
		SysDictItem:   newSysDictItem(db, opts...),
		SysDictType:   newSysDictType(db, opts...),
		SysLoginLog:   newSysLoginLog(db, opts...),
		SysMenu:       newSysMenu(db, opts...),
		SysOperateLog: newSysOperateLog(db, opts...),
		SysPost:       newSysPost(db, opts...),
		SysRole:       newSysRole(db, opts...),
		SysRoleMenu:   newSysRoleMenu(db, opts...),
		SysUser:       newSysUser(db, opts...),
		SysUserRole:   newSysUserRole(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	SysDept       sysDept
	SysDictItem   sysDictItem
	SysDictType   sysDictType
	SysLoginLog   sysLoginLog
	SysMenu       sysMenu
	SysOperateLog sysOperateLog
	SysPost       sysPost
	SysRole       sysRole
	SysRoleMenu   sysRoleMenu
	SysUser       sysUser
	SysUserRole   sysUserRole
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		SysDept:       q.SysDept.clone(db),
		SysDictItem:   q.SysDictItem.clone(db),
		SysDictType:   q.SysDictType.clone(db),
		SysLoginLog:   q.SysLoginLog.clone(db),
		SysMenu:       q.SysMenu.clone(db),
		SysOperateLog: q.SysOperateLog.clone(db),
		SysPost:       q.SysPost.clone(db),
		SysRole:       q.SysRole.clone(db),
		SysRoleMenu:   q.SysRoleMenu.clone(db),
		SysUser:       q.SysUser.clone(db),
		SysUserRole:   q.SysUserRole.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		SysDept:       q.SysDept.replaceDB(db),
		SysDictItem:   q.SysDictItem.replaceDB(db),
		SysDictType:   q.SysDictType.replaceDB(db),
		SysLoginLog:   q.SysLoginLog.replaceDB(db),
		SysMenu:       q.SysMenu.replaceDB(db),
		SysOperateLog: q.SysOperateLog.replaceDB(db),
		SysPost:       q.SysPost.replaceDB(db),
		SysRole:       q.SysRole.replaceDB(db),
		SysRoleMenu:   q.SysRoleMenu.replaceDB(db),
		SysUser:       q.SysUser.replaceDB(db),
		SysUserRole:   q.SysUserRole.replaceDB(db),
	}
}

type queryCtx struct {
	SysDept       ISysDeptDo
	SysDictItem   ISysDictItemDo
	SysDictType   ISysDictTypeDo
	SysLoginLog   ISysLoginLogDo
	SysMenu       ISysMenuDo
	SysOperateLog ISysOperateLogDo
	SysPost       ISysPostDo
	SysRole       ISysRoleDo
	SysRoleMenu   ISysRoleMenuDo
	SysUser       ISysUserDo
	SysUserRole   ISysUserRoleDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		SysDept:       q.SysDept.WithContext(ctx),
		SysDictItem:   q.SysDictItem.WithContext(ctx),
		SysDictType:   q.SysDictType.WithContext(ctx),
		SysLoginLog:   q.SysLoginLog.WithContext(ctx),
		SysMenu:       q.SysMenu.WithContext(ctx),
		SysOperateLog: q.SysOperateLog.WithContext(ctx),
		SysPost:       q.SysPost.WithContext(ctx),
		SysRole:       q.SysRole.WithContext(ctx),
		SysRoleMenu:   q.SysRoleMenu.WithContext(ctx),
		SysUser:       q.SysUser.WithContext(ctx),
		SysUserRole:   q.SysUserRole.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
