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

	"github.com/feihua/zero-admin/rpc/cms/gen/model"
)

func newCmsPreferredAreaProductRelation(db *gorm.DB, opts ...gen.DOOption) cmsPreferredAreaProductRelation {
	_cmsPreferredAreaProductRelation := cmsPreferredAreaProductRelation{}

	_cmsPreferredAreaProductRelation.cmsPreferredAreaProductRelationDo.UseDB(db, opts...)
	_cmsPreferredAreaProductRelation.cmsPreferredAreaProductRelationDo.UseModel(&model.CmsPreferredAreaProductRelation{})

	tableName := _cmsPreferredAreaProductRelation.cmsPreferredAreaProductRelationDo.TableName()
	_cmsPreferredAreaProductRelation.ALL = field.NewAsterisk(tableName)
	_cmsPreferredAreaProductRelation.ID = field.NewInt64(tableName, "id")
	_cmsPreferredAreaProductRelation.PreferredAreaID = field.NewInt64(tableName, "preferred_area_id")
	_cmsPreferredAreaProductRelation.ProductID = field.NewInt64(tableName, "product_id")

	_cmsPreferredAreaProductRelation.fillFieldMap()

	return _cmsPreferredAreaProductRelation
}

// cmsPreferredAreaProductRelation 优选专区和产品关系表
type cmsPreferredAreaProductRelation struct {
	cmsPreferredAreaProductRelationDo cmsPreferredAreaProductRelationDo

	ALL             field.Asterisk
	ID              field.Int64
	PreferredAreaID field.Int64
	ProductID       field.Int64

	fieldMap map[string]field.Expr
}

func (c cmsPreferredAreaProductRelation) Table(newTableName string) *cmsPreferredAreaProductRelation {
	c.cmsPreferredAreaProductRelationDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cmsPreferredAreaProductRelation) As(alias string) *cmsPreferredAreaProductRelation {
	c.cmsPreferredAreaProductRelationDo.DO = *(c.cmsPreferredAreaProductRelationDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cmsPreferredAreaProductRelation) updateTableName(table string) *cmsPreferredAreaProductRelation {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.PreferredAreaID = field.NewInt64(table, "preferred_area_id")
	c.ProductID = field.NewInt64(table, "product_id")

	c.fillFieldMap()

	return c
}

func (c *cmsPreferredAreaProductRelation) WithContext(ctx context.Context) ICmsPreferredAreaProductRelationDo {
	return c.cmsPreferredAreaProductRelationDo.WithContext(ctx)
}

func (c cmsPreferredAreaProductRelation) TableName() string {
	return c.cmsPreferredAreaProductRelationDo.TableName()
}

func (c cmsPreferredAreaProductRelation) Alias() string {
	return c.cmsPreferredAreaProductRelationDo.Alias()
}

func (c cmsPreferredAreaProductRelation) Columns(cols ...field.Expr) gen.Columns {
	return c.cmsPreferredAreaProductRelationDo.Columns(cols...)
}

func (c *cmsPreferredAreaProductRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cmsPreferredAreaProductRelation) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 3)
	c.fieldMap["id"] = c.ID
	c.fieldMap["preferred_area_id"] = c.PreferredAreaID
	c.fieldMap["product_id"] = c.ProductID
}

func (c cmsPreferredAreaProductRelation) clone(db *gorm.DB) cmsPreferredAreaProductRelation {
	c.cmsPreferredAreaProductRelationDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cmsPreferredAreaProductRelation) replaceDB(db *gorm.DB) cmsPreferredAreaProductRelation {
	c.cmsPreferredAreaProductRelationDo.ReplaceDB(db)
	return c
}

type cmsPreferredAreaProductRelationDo struct{ gen.DO }

type ICmsPreferredAreaProductRelationDo interface {
	gen.SubQuery
	Debug() ICmsPreferredAreaProductRelationDo
	WithContext(ctx context.Context) ICmsPreferredAreaProductRelationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICmsPreferredAreaProductRelationDo
	WriteDB() ICmsPreferredAreaProductRelationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICmsPreferredAreaProductRelationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICmsPreferredAreaProductRelationDo
	Not(conds ...gen.Condition) ICmsPreferredAreaProductRelationDo
	Or(conds ...gen.Condition) ICmsPreferredAreaProductRelationDo
	Select(conds ...field.Expr) ICmsPreferredAreaProductRelationDo
	Where(conds ...gen.Condition) ICmsPreferredAreaProductRelationDo
	Order(conds ...field.Expr) ICmsPreferredAreaProductRelationDo
	Distinct(cols ...field.Expr) ICmsPreferredAreaProductRelationDo
	Omit(cols ...field.Expr) ICmsPreferredAreaProductRelationDo
	Join(table schema.Tabler, on ...field.Expr) ICmsPreferredAreaProductRelationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICmsPreferredAreaProductRelationDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICmsPreferredAreaProductRelationDo
	Group(cols ...field.Expr) ICmsPreferredAreaProductRelationDo
	Having(conds ...gen.Condition) ICmsPreferredAreaProductRelationDo
	Limit(limit int) ICmsPreferredAreaProductRelationDo
	Offset(offset int) ICmsPreferredAreaProductRelationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICmsPreferredAreaProductRelationDo
	Unscoped() ICmsPreferredAreaProductRelationDo
	Create(values ...*model.CmsPreferredAreaProductRelation) error
	CreateInBatches(values []*model.CmsPreferredAreaProductRelation, batchSize int) error
	Save(values ...*model.CmsPreferredAreaProductRelation) error
	First() (*model.CmsPreferredAreaProductRelation, error)
	Take() (*model.CmsPreferredAreaProductRelation, error)
	Last() (*model.CmsPreferredAreaProductRelation, error)
	Find() ([]*model.CmsPreferredAreaProductRelation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CmsPreferredAreaProductRelation, err error)
	FindInBatches(result *[]*model.CmsPreferredAreaProductRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CmsPreferredAreaProductRelation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICmsPreferredAreaProductRelationDo
	Assign(attrs ...field.AssignExpr) ICmsPreferredAreaProductRelationDo
	Joins(fields ...field.RelationField) ICmsPreferredAreaProductRelationDo
	Preload(fields ...field.RelationField) ICmsPreferredAreaProductRelationDo
	FirstOrInit() (*model.CmsPreferredAreaProductRelation, error)
	FirstOrCreate() (*model.CmsPreferredAreaProductRelation, error)
	FindByPage(offset int, limit int) (result []*model.CmsPreferredAreaProductRelation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICmsPreferredAreaProductRelationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cmsPreferredAreaProductRelationDo) Debug() ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Debug())
}

func (c cmsPreferredAreaProductRelationDo) WithContext(ctx context.Context) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cmsPreferredAreaProductRelationDo) ReadDB() ICmsPreferredAreaProductRelationDo {
	return c.Clauses(dbresolver.Read)
}

func (c cmsPreferredAreaProductRelationDo) WriteDB() ICmsPreferredAreaProductRelationDo {
	return c.Clauses(dbresolver.Write)
}

func (c cmsPreferredAreaProductRelationDo) Session(config *gorm.Session) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Session(config))
}

func (c cmsPreferredAreaProductRelationDo) Clauses(conds ...clause.Expression) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cmsPreferredAreaProductRelationDo) Returning(value interface{}, columns ...string) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cmsPreferredAreaProductRelationDo) Not(conds ...gen.Condition) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cmsPreferredAreaProductRelationDo) Or(conds ...gen.Condition) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cmsPreferredAreaProductRelationDo) Select(conds ...field.Expr) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cmsPreferredAreaProductRelationDo) Where(conds ...gen.Condition) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cmsPreferredAreaProductRelationDo) Order(conds ...field.Expr) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cmsPreferredAreaProductRelationDo) Distinct(cols ...field.Expr) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cmsPreferredAreaProductRelationDo) Omit(cols ...field.Expr) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cmsPreferredAreaProductRelationDo) Join(table schema.Tabler, on ...field.Expr) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cmsPreferredAreaProductRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cmsPreferredAreaProductRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cmsPreferredAreaProductRelationDo) Group(cols ...field.Expr) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cmsPreferredAreaProductRelationDo) Having(conds ...gen.Condition) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cmsPreferredAreaProductRelationDo) Limit(limit int) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cmsPreferredAreaProductRelationDo) Offset(offset int) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cmsPreferredAreaProductRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cmsPreferredAreaProductRelationDo) Unscoped() ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cmsPreferredAreaProductRelationDo) Create(values ...*model.CmsPreferredAreaProductRelation) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cmsPreferredAreaProductRelationDo) CreateInBatches(values []*model.CmsPreferredAreaProductRelation, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cmsPreferredAreaProductRelationDo) Save(values ...*model.CmsPreferredAreaProductRelation) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cmsPreferredAreaProductRelationDo) First() (*model.CmsPreferredAreaProductRelation, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsPreferredAreaProductRelation), nil
	}
}

func (c cmsPreferredAreaProductRelationDo) Take() (*model.CmsPreferredAreaProductRelation, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsPreferredAreaProductRelation), nil
	}
}

func (c cmsPreferredAreaProductRelationDo) Last() (*model.CmsPreferredAreaProductRelation, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsPreferredAreaProductRelation), nil
	}
}

func (c cmsPreferredAreaProductRelationDo) Find() ([]*model.CmsPreferredAreaProductRelation, error) {
	result, err := c.DO.Find()
	return result.([]*model.CmsPreferredAreaProductRelation), err
}

func (c cmsPreferredAreaProductRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CmsPreferredAreaProductRelation, err error) {
	buf := make([]*model.CmsPreferredAreaProductRelation, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cmsPreferredAreaProductRelationDo) FindInBatches(result *[]*model.CmsPreferredAreaProductRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cmsPreferredAreaProductRelationDo) Attrs(attrs ...field.AssignExpr) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cmsPreferredAreaProductRelationDo) Assign(attrs ...field.AssignExpr) ICmsPreferredAreaProductRelationDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cmsPreferredAreaProductRelationDo) Joins(fields ...field.RelationField) ICmsPreferredAreaProductRelationDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cmsPreferredAreaProductRelationDo) Preload(fields ...field.RelationField) ICmsPreferredAreaProductRelationDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cmsPreferredAreaProductRelationDo) FirstOrInit() (*model.CmsPreferredAreaProductRelation, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsPreferredAreaProductRelation), nil
	}
}

func (c cmsPreferredAreaProductRelationDo) FirstOrCreate() (*model.CmsPreferredAreaProductRelation, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsPreferredAreaProductRelation), nil
	}
}

func (c cmsPreferredAreaProductRelationDo) FindByPage(offset int, limit int) (result []*model.CmsPreferredAreaProductRelation, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cmsPreferredAreaProductRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cmsPreferredAreaProductRelationDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cmsPreferredAreaProductRelationDo) Delete(models ...*model.CmsPreferredAreaProductRelation) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cmsPreferredAreaProductRelationDo) withDO(do gen.Dao) *cmsPreferredAreaProductRelationDo {
	c.DO = *do.(*gen.DO)
	return c
}
