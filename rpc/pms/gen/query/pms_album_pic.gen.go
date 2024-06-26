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

	"github.com/feihua/zero-admin/rpc/pms/gen/model"
)

func newPmsAlbumPic(db *gorm.DB, opts ...gen.DOOption) pmsAlbumPic {
	_pmsAlbumPic := pmsAlbumPic{}

	_pmsAlbumPic.pmsAlbumPicDo.UseDB(db, opts...)
	_pmsAlbumPic.pmsAlbumPicDo.UseModel(&model.PmsAlbumPic{})

	tableName := _pmsAlbumPic.pmsAlbumPicDo.TableName()
	_pmsAlbumPic.ALL = field.NewAsterisk(tableName)
	_pmsAlbumPic.ID = field.NewInt64(tableName, "id")
	_pmsAlbumPic.AlbumID = field.NewInt64(tableName, "album_id")
	_pmsAlbumPic.Pic = field.NewString(tableName, "pic")

	_pmsAlbumPic.fillFieldMap()

	return _pmsAlbumPic
}

// pmsAlbumPic 画册图片表
type pmsAlbumPic struct {
	pmsAlbumPicDo pmsAlbumPicDo

	ALL     field.Asterisk
	ID      field.Int64
	AlbumID field.Int64
	Pic     field.String

	fieldMap map[string]field.Expr
}

func (p pmsAlbumPic) Table(newTableName string) *pmsAlbumPic {
	p.pmsAlbumPicDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsAlbumPic) As(alias string) *pmsAlbumPic {
	p.pmsAlbumPicDo.DO = *(p.pmsAlbumPicDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsAlbumPic) updateTableName(table string) *pmsAlbumPic {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.AlbumID = field.NewInt64(table, "album_id")
	p.Pic = field.NewString(table, "pic")

	p.fillFieldMap()

	return p
}

func (p *pmsAlbumPic) WithContext(ctx context.Context) IPmsAlbumPicDo {
	return p.pmsAlbumPicDo.WithContext(ctx)
}

func (p pmsAlbumPic) TableName() string { return p.pmsAlbumPicDo.TableName() }

func (p pmsAlbumPic) Alias() string { return p.pmsAlbumPicDo.Alias() }

func (p pmsAlbumPic) Columns(cols ...field.Expr) gen.Columns { return p.pmsAlbumPicDo.Columns(cols...) }

func (p *pmsAlbumPic) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsAlbumPic) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 3)
	p.fieldMap["id"] = p.ID
	p.fieldMap["album_id"] = p.AlbumID
	p.fieldMap["pic"] = p.Pic
}

func (p pmsAlbumPic) clone(db *gorm.DB) pmsAlbumPic {
	p.pmsAlbumPicDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsAlbumPic) replaceDB(db *gorm.DB) pmsAlbumPic {
	p.pmsAlbumPicDo.ReplaceDB(db)
	return p
}

type pmsAlbumPicDo struct{ gen.DO }

type IPmsAlbumPicDo interface {
	gen.SubQuery
	Debug() IPmsAlbumPicDo
	WithContext(ctx context.Context) IPmsAlbumPicDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPmsAlbumPicDo
	WriteDB() IPmsAlbumPicDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPmsAlbumPicDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPmsAlbumPicDo
	Not(conds ...gen.Condition) IPmsAlbumPicDo
	Or(conds ...gen.Condition) IPmsAlbumPicDo
	Select(conds ...field.Expr) IPmsAlbumPicDo
	Where(conds ...gen.Condition) IPmsAlbumPicDo
	Order(conds ...field.Expr) IPmsAlbumPicDo
	Distinct(cols ...field.Expr) IPmsAlbumPicDo
	Omit(cols ...field.Expr) IPmsAlbumPicDo
	Join(table schema.Tabler, on ...field.Expr) IPmsAlbumPicDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPmsAlbumPicDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPmsAlbumPicDo
	Group(cols ...field.Expr) IPmsAlbumPicDo
	Having(conds ...gen.Condition) IPmsAlbumPicDo
	Limit(limit int) IPmsAlbumPicDo
	Offset(offset int) IPmsAlbumPicDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsAlbumPicDo
	Unscoped() IPmsAlbumPicDo
	Create(values ...*model.PmsAlbumPic) error
	CreateInBatches(values []*model.PmsAlbumPic, batchSize int) error
	Save(values ...*model.PmsAlbumPic) error
	First() (*model.PmsAlbumPic, error)
	Take() (*model.PmsAlbumPic, error)
	Last() (*model.PmsAlbumPic, error)
	Find() ([]*model.PmsAlbumPic, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsAlbumPic, err error)
	FindInBatches(result *[]*model.PmsAlbumPic, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PmsAlbumPic) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPmsAlbumPicDo
	Assign(attrs ...field.AssignExpr) IPmsAlbumPicDo
	Joins(fields ...field.RelationField) IPmsAlbumPicDo
	Preload(fields ...field.RelationField) IPmsAlbumPicDo
	FirstOrInit() (*model.PmsAlbumPic, error)
	FirstOrCreate() (*model.PmsAlbumPic, error)
	FindByPage(offset int, limit int) (result []*model.PmsAlbumPic, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPmsAlbumPicDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pmsAlbumPicDo) Debug() IPmsAlbumPicDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsAlbumPicDo) WithContext(ctx context.Context) IPmsAlbumPicDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsAlbumPicDo) ReadDB() IPmsAlbumPicDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsAlbumPicDo) WriteDB() IPmsAlbumPicDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsAlbumPicDo) Session(config *gorm.Session) IPmsAlbumPicDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsAlbumPicDo) Clauses(conds ...clause.Expression) IPmsAlbumPicDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsAlbumPicDo) Returning(value interface{}, columns ...string) IPmsAlbumPicDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsAlbumPicDo) Not(conds ...gen.Condition) IPmsAlbumPicDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsAlbumPicDo) Or(conds ...gen.Condition) IPmsAlbumPicDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsAlbumPicDo) Select(conds ...field.Expr) IPmsAlbumPicDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsAlbumPicDo) Where(conds ...gen.Condition) IPmsAlbumPicDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsAlbumPicDo) Order(conds ...field.Expr) IPmsAlbumPicDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsAlbumPicDo) Distinct(cols ...field.Expr) IPmsAlbumPicDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsAlbumPicDo) Omit(cols ...field.Expr) IPmsAlbumPicDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsAlbumPicDo) Join(table schema.Tabler, on ...field.Expr) IPmsAlbumPicDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsAlbumPicDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPmsAlbumPicDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsAlbumPicDo) RightJoin(table schema.Tabler, on ...field.Expr) IPmsAlbumPicDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsAlbumPicDo) Group(cols ...field.Expr) IPmsAlbumPicDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsAlbumPicDo) Having(conds ...gen.Condition) IPmsAlbumPicDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsAlbumPicDo) Limit(limit int) IPmsAlbumPicDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsAlbumPicDo) Offset(offset int) IPmsAlbumPicDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsAlbumPicDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsAlbumPicDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsAlbumPicDo) Unscoped() IPmsAlbumPicDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsAlbumPicDo) Create(values ...*model.PmsAlbumPic) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsAlbumPicDo) CreateInBatches(values []*model.PmsAlbumPic, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsAlbumPicDo) Save(values ...*model.PmsAlbumPic) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsAlbumPicDo) First() (*model.PmsAlbumPic, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAlbumPic), nil
	}
}

func (p pmsAlbumPicDo) Take() (*model.PmsAlbumPic, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAlbumPic), nil
	}
}

func (p pmsAlbumPicDo) Last() (*model.PmsAlbumPic, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAlbumPic), nil
	}
}

func (p pmsAlbumPicDo) Find() ([]*model.PmsAlbumPic, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsAlbumPic), err
}

func (p pmsAlbumPicDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsAlbumPic, err error) {
	buf := make([]*model.PmsAlbumPic, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsAlbumPicDo) FindInBatches(result *[]*model.PmsAlbumPic, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsAlbumPicDo) Attrs(attrs ...field.AssignExpr) IPmsAlbumPicDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsAlbumPicDo) Assign(attrs ...field.AssignExpr) IPmsAlbumPicDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsAlbumPicDo) Joins(fields ...field.RelationField) IPmsAlbumPicDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsAlbumPicDo) Preload(fields ...field.RelationField) IPmsAlbumPicDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsAlbumPicDo) FirstOrInit() (*model.PmsAlbumPic, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAlbumPic), nil
	}
}

func (p pmsAlbumPicDo) FirstOrCreate() (*model.PmsAlbumPic, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAlbumPic), nil
	}
}

func (p pmsAlbumPicDo) FindByPage(offset int, limit int) (result []*model.PmsAlbumPic, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pmsAlbumPicDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsAlbumPicDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsAlbumPicDo) Delete(models ...*model.PmsAlbumPic) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsAlbumPicDo) withDO(do gen.Dao) *pmsAlbumPicDo {
	p.DO = *do.(*gen.DO)
	return p
}
