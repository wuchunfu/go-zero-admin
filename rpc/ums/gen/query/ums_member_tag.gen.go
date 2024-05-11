// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/feihua/zero-admin/rpc/ums/gen/model"
)

func newUmsMemberTag(db *gorm.DB, opts ...gen.DOOption) umsMemberTag {
	_umsMemberTag := umsMemberTag{}

	_umsMemberTag.umsMemberTagDo.UseDB(db, opts...)
	_umsMemberTag.umsMemberTagDo.UseModel(&model.UmsMemberTag{})

	tableName := _umsMemberTag.umsMemberTagDo.TableName()
	_umsMemberTag.ALL = field.NewAsterisk(tableName)
	_umsMemberTag.ID = field.NewInt64(tableName, "id")
	_umsMemberTag.TagName = field.NewString(tableName, "tag_name")
	_umsMemberTag.FinishOrderCount = field.NewInt32(tableName, "finish_order_count")
	_umsMemberTag.FinishOrderAmount = field.NewInt64(tableName, "finish_order_amount")

	_umsMemberTag.fillFieldMap()

	return _umsMemberTag
}

// umsMemberTag 用户标签表
type umsMemberTag struct {
	umsMemberTagDo umsMemberTagDo

	ALL               field.Asterisk
	ID                field.Int64
	TagName           field.String // 标签名称
	FinishOrderCount  field.Int32  // 自动打标签完成订单数量
	FinishOrderAmount field.Int64  // 自动打标签完成订单金额

	fieldMap map[string]field.Expr
}

func (u umsMemberTag) Table(newTableName string) *umsMemberTag {
	u.umsMemberTagDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u umsMemberTag) As(alias string) *umsMemberTag {
	u.umsMemberTagDo.DO = *(u.umsMemberTagDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *umsMemberTag) updateTableName(table string) *umsMemberTag {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.TagName = field.NewString(table, "tag_name")
	u.FinishOrderCount = field.NewInt32(table, "finish_order_count")
	u.FinishOrderAmount = field.NewInt64(table, "finish_order_amount")

	u.fillFieldMap()

	return u
}

func (u *umsMemberTag) WithContext(ctx context.Context) IUmsMemberTagDo {
	return u.umsMemberTagDo.WithContext(ctx)
}

func (u umsMemberTag) TableName() string { return u.umsMemberTagDo.TableName() }

func (u umsMemberTag) Alias() string { return u.umsMemberTagDo.Alias() }

func (u umsMemberTag) Columns(cols ...field.Expr) gen.Columns {
	return u.umsMemberTagDo.Columns(cols...)
}

func (u *umsMemberTag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *umsMemberTag) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 4)
	u.fieldMap["id"] = u.ID
	u.fieldMap["tag_name"] = u.TagName
	u.fieldMap["finish_order_count"] = u.FinishOrderCount
	u.fieldMap["finish_order_amount"] = u.FinishOrderAmount
}

func (u umsMemberTag) clone(db *gorm.DB) umsMemberTag {
	u.umsMemberTagDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u umsMemberTag) replaceDB(db *gorm.DB) umsMemberTag {
	u.umsMemberTagDo.ReplaceDB(db)
	return u
}

type umsMemberTagDo struct{ gen.DO }

type IUmsMemberTagDo interface {
	gen.SubQuery
	Debug() IUmsMemberTagDo
	WithContext(ctx context.Context) IUmsMemberTagDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUmsMemberTagDo
	WriteDB() IUmsMemberTagDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUmsMemberTagDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUmsMemberTagDo
	Not(conds ...gen.Condition) IUmsMemberTagDo
	Or(conds ...gen.Condition) IUmsMemberTagDo
	Select(conds ...field.Expr) IUmsMemberTagDo
	Where(conds ...gen.Condition) IUmsMemberTagDo
	Order(conds ...field.Expr) IUmsMemberTagDo
	Distinct(cols ...field.Expr) IUmsMemberTagDo
	Omit(cols ...field.Expr) IUmsMemberTagDo
	Join(table schema.Tabler, on ...field.Expr) IUmsMemberTagDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberTagDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberTagDo
	Group(cols ...field.Expr) IUmsMemberTagDo
	Having(conds ...gen.Condition) IUmsMemberTagDo
	Limit(limit int) IUmsMemberTagDo
	Offset(offset int) IUmsMemberTagDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberTagDo
	Unscoped() IUmsMemberTagDo
	Create(values ...*model.UmsMemberTag) error
	CreateInBatches(values []*model.UmsMemberTag, batchSize int) error
	Save(values ...*model.UmsMemberTag) error
	First() (*model.UmsMemberTag, error)
	Take() (*model.UmsMemberTag, error)
	Last() (*model.UmsMemberTag, error)
	Find() ([]*model.UmsMemberTag, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMemberTag, err error)
	FindInBatches(result *[]*model.UmsMemberTag, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UmsMemberTag) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUmsMemberTagDo
	Assign(attrs ...field.AssignExpr) IUmsMemberTagDo
	Joins(fields ...field.RelationField) IUmsMemberTagDo
	Preload(fields ...field.RelationField) IUmsMemberTagDo
	FirstOrInit() (*model.UmsMemberTag, error)
	FirstOrCreate() (*model.UmsMemberTag, error)
	FindByPage(offset int, limit int) (result []*model.UmsMemberTag, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUmsMemberTagDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u umsMemberTagDo) Debug() IUmsMemberTagDo {
	return u.withDO(u.DO.Debug())
}

func (u umsMemberTagDo) WithContext(ctx context.Context) IUmsMemberTagDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u umsMemberTagDo) ReadDB() IUmsMemberTagDo {
	return u.Clauses(dbresolver.Read)
}

func (u umsMemberTagDo) WriteDB() IUmsMemberTagDo {
	return u.Clauses(dbresolver.Write)
}

func (u umsMemberTagDo) Session(config *gorm.Session) IUmsMemberTagDo {
	return u.withDO(u.DO.Session(config))
}

func (u umsMemberTagDo) Clauses(conds ...clause.Expression) IUmsMemberTagDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u umsMemberTagDo) Returning(value interface{}, columns ...string) IUmsMemberTagDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u umsMemberTagDo) Not(conds ...gen.Condition) IUmsMemberTagDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u umsMemberTagDo) Or(conds ...gen.Condition) IUmsMemberTagDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u umsMemberTagDo) Select(conds ...field.Expr) IUmsMemberTagDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u umsMemberTagDo) Where(conds ...gen.Condition) IUmsMemberTagDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u umsMemberTagDo) Order(conds ...field.Expr) IUmsMemberTagDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u umsMemberTagDo) Distinct(cols ...field.Expr) IUmsMemberTagDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u umsMemberTagDo) Omit(cols ...field.Expr) IUmsMemberTagDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u umsMemberTagDo) Join(table schema.Tabler, on ...field.Expr) IUmsMemberTagDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u umsMemberTagDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberTagDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u umsMemberTagDo) RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberTagDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u umsMemberTagDo) Group(cols ...field.Expr) IUmsMemberTagDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u umsMemberTagDo) Having(conds ...gen.Condition) IUmsMemberTagDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u umsMemberTagDo) Limit(limit int) IUmsMemberTagDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u umsMemberTagDo) Offset(offset int) IUmsMemberTagDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u umsMemberTagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberTagDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u umsMemberTagDo) Unscoped() IUmsMemberTagDo {
	return u.withDO(u.DO.Unscoped())
}

func (u umsMemberTagDo) Create(values ...*model.UmsMemberTag) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u umsMemberTagDo) CreateInBatches(values []*model.UmsMemberTag, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u umsMemberTagDo) Save(values ...*model.UmsMemberTag) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u umsMemberTagDo) First() (*model.UmsMemberTag, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberTag), nil
	}
}

func (u umsMemberTagDo) Take() (*model.UmsMemberTag, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberTag), nil
	}
}

func (u umsMemberTagDo) Last() (*model.UmsMemberTag, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberTag), nil
	}
}

func (u umsMemberTagDo) Find() ([]*model.UmsMemberTag, error) {
	result, err := u.DO.Find()
	return result.([]*model.UmsMemberTag), err
}

func (u umsMemberTagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMemberTag, err error) {
	buf := make([]*model.UmsMemberTag, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u umsMemberTagDo) FindInBatches(result *[]*model.UmsMemberTag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u umsMemberTagDo) Attrs(attrs ...field.AssignExpr) IUmsMemberTagDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u umsMemberTagDo) Assign(attrs ...field.AssignExpr) IUmsMemberTagDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u umsMemberTagDo) Joins(fields ...field.RelationField) IUmsMemberTagDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u umsMemberTagDo) Preload(fields ...field.RelationField) IUmsMemberTagDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u umsMemberTagDo) FirstOrInit() (*model.UmsMemberTag, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberTag), nil
	}
}

func (u umsMemberTagDo) FirstOrCreate() (*model.UmsMemberTag, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberTag), nil
	}
}

func (u umsMemberTagDo) FindByPage(offset int, limit int) (result []*model.UmsMemberTag, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u umsMemberTagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u umsMemberTagDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u umsMemberTagDo) Delete(models ...*model.UmsMemberTag) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *umsMemberTagDo) withDO(do gen.Dao) *umsMemberTagDo {
	u.DO = *do.(*gen.DO)
	return u
}
