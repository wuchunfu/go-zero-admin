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

func newCmsMemberReport(db *gorm.DB, opts ...gen.DOOption) cmsMemberReport {
	_cmsMemberReport := cmsMemberReport{}

	_cmsMemberReport.cmsMemberReportDo.UseDB(db, opts...)
	_cmsMemberReport.cmsMemberReportDo.UseModel(&model.CmsMemberReport{})

	tableName := _cmsMemberReport.cmsMemberReportDo.TableName()
	_cmsMemberReport.ALL = field.NewAsterisk(tableName)
	_cmsMemberReport.ID = field.NewInt64(tableName, "id")
	_cmsMemberReport.ReportType = field.NewInt32(tableName, "report_type")
	_cmsMemberReport.ReportMemberName = field.NewString(tableName, "report_member_name")
	_cmsMemberReport.ReportObject = field.NewString(tableName, "report_object")
	_cmsMemberReport.ReportStatus = field.NewInt32(tableName, "report_status")
	_cmsMemberReport.HandleStatus = field.NewInt32(tableName, "handle_status")
	_cmsMemberReport.Note = field.NewString(tableName, "note")
	_cmsMemberReport.CreateTime = field.NewTime(tableName, "create_time")

	_cmsMemberReport.fillFieldMap()

	return _cmsMemberReport
}

// cmsMemberReport 用户举报表
type cmsMemberReport struct {
	cmsMemberReportDo cmsMemberReportDo

	ALL              field.Asterisk
	ID               field.Int64
	ReportType       field.Int32  // 举报类型：0->商品评价；1->话题内容；2->用户评论
	ReportMemberName field.String // 举报人
	ReportObject     field.String
	ReportStatus     field.Int32 // 举报状态：0->未处理；1->已处理
	HandleStatus     field.Int32 // 处理结果：0->无效；1->有效；2->恶意
	Note             field.String
	CreateTime       field.Time // 创建时间

	fieldMap map[string]field.Expr
}

func (c cmsMemberReport) Table(newTableName string) *cmsMemberReport {
	c.cmsMemberReportDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cmsMemberReport) As(alias string) *cmsMemberReport {
	c.cmsMemberReportDo.DO = *(c.cmsMemberReportDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cmsMemberReport) updateTableName(table string) *cmsMemberReport {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.ReportType = field.NewInt32(table, "report_type")
	c.ReportMemberName = field.NewString(table, "report_member_name")
	c.ReportObject = field.NewString(table, "report_object")
	c.ReportStatus = field.NewInt32(table, "report_status")
	c.HandleStatus = field.NewInt32(table, "handle_status")
	c.Note = field.NewString(table, "note")
	c.CreateTime = field.NewTime(table, "create_time")

	c.fillFieldMap()

	return c
}

func (c *cmsMemberReport) WithContext(ctx context.Context) ICmsMemberReportDo {
	return c.cmsMemberReportDo.WithContext(ctx)
}

func (c cmsMemberReport) TableName() string { return c.cmsMemberReportDo.TableName() }

func (c cmsMemberReport) Alias() string { return c.cmsMemberReportDo.Alias() }

func (c cmsMemberReport) Columns(cols ...field.Expr) gen.Columns {
	return c.cmsMemberReportDo.Columns(cols...)
}

func (c *cmsMemberReport) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cmsMemberReport) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 8)
	c.fieldMap["id"] = c.ID
	c.fieldMap["report_type"] = c.ReportType
	c.fieldMap["report_member_name"] = c.ReportMemberName
	c.fieldMap["report_object"] = c.ReportObject
	c.fieldMap["report_status"] = c.ReportStatus
	c.fieldMap["handle_status"] = c.HandleStatus
	c.fieldMap["note"] = c.Note
	c.fieldMap["create_time"] = c.CreateTime
}

func (c cmsMemberReport) clone(db *gorm.DB) cmsMemberReport {
	c.cmsMemberReportDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cmsMemberReport) replaceDB(db *gorm.DB) cmsMemberReport {
	c.cmsMemberReportDo.ReplaceDB(db)
	return c
}

type cmsMemberReportDo struct{ gen.DO }

type ICmsMemberReportDo interface {
	gen.SubQuery
	Debug() ICmsMemberReportDo
	WithContext(ctx context.Context) ICmsMemberReportDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICmsMemberReportDo
	WriteDB() ICmsMemberReportDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICmsMemberReportDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICmsMemberReportDo
	Not(conds ...gen.Condition) ICmsMemberReportDo
	Or(conds ...gen.Condition) ICmsMemberReportDo
	Select(conds ...field.Expr) ICmsMemberReportDo
	Where(conds ...gen.Condition) ICmsMemberReportDo
	Order(conds ...field.Expr) ICmsMemberReportDo
	Distinct(cols ...field.Expr) ICmsMemberReportDo
	Omit(cols ...field.Expr) ICmsMemberReportDo
	Join(table schema.Tabler, on ...field.Expr) ICmsMemberReportDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICmsMemberReportDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICmsMemberReportDo
	Group(cols ...field.Expr) ICmsMemberReportDo
	Having(conds ...gen.Condition) ICmsMemberReportDo
	Limit(limit int) ICmsMemberReportDo
	Offset(offset int) ICmsMemberReportDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICmsMemberReportDo
	Unscoped() ICmsMemberReportDo
	Create(values ...*model.CmsMemberReport) error
	CreateInBatches(values []*model.CmsMemberReport, batchSize int) error
	Save(values ...*model.CmsMemberReport) error
	First() (*model.CmsMemberReport, error)
	Take() (*model.CmsMemberReport, error)
	Last() (*model.CmsMemberReport, error)
	Find() ([]*model.CmsMemberReport, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CmsMemberReport, err error)
	FindInBatches(result *[]*model.CmsMemberReport, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CmsMemberReport) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICmsMemberReportDo
	Assign(attrs ...field.AssignExpr) ICmsMemberReportDo
	Joins(fields ...field.RelationField) ICmsMemberReportDo
	Preload(fields ...field.RelationField) ICmsMemberReportDo
	FirstOrInit() (*model.CmsMemberReport, error)
	FirstOrCreate() (*model.CmsMemberReport, error)
	FindByPage(offset int, limit int) (result []*model.CmsMemberReport, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICmsMemberReportDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cmsMemberReportDo) Debug() ICmsMemberReportDo {
	return c.withDO(c.DO.Debug())
}

func (c cmsMemberReportDo) WithContext(ctx context.Context) ICmsMemberReportDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cmsMemberReportDo) ReadDB() ICmsMemberReportDo {
	return c.Clauses(dbresolver.Read)
}

func (c cmsMemberReportDo) WriteDB() ICmsMemberReportDo {
	return c.Clauses(dbresolver.Write)
}

func (c cmsMemberReportDo) Session(config *gorm.Session) ICmsMemberReportDo {
	return c.withDO(c.DO.Session(config))
}

func (c cmsMemberReportDo) Clauses(conds ...clause.Expression) ICmsMemberReportDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cmsMemberReportDo) Returning(value interface{}, columns ...string) ICmsMemberReportDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cmsMemberReportDo) Not(conds ...gen.Condition) ICmsMemberReportDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cmsMemberReportDo) Or(conds ...gen.Condition) ICmsMemberReportDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cmsMemberReportDo) Select(conds ...field.Expr) ICmsMemberReportDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cmsMemberReportDo) Where(conds ...gen.Condition) ICmsMemberReportDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cmsMemberReportDo) Order(conds ...field.Expr) ICmsMemberReportDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cmsMemberReportDo) Distinct(cols ...field.Expr) ICmsMemberReportDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cmsMemberReportDo) Omit(cols ...field.Expr) ICmsMemberReportDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cmsMemberReportDo) Join(table schema.Tabler, on ...field.Expr) ICmsMemberReportDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cmsMemberReportDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICmsMemberReportDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cmsMemberReportDo) RightJoin(table schema.Tabler, on ...field.Expr) ICmsMemberReportDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cmsMemberReportDo) Group(cols ...field.Expr) ICmsMemberReportDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cmsMemberReportDo) Having(conds ...gen.Condition) ICmsMemberReportDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cmsMemberReportDo) Limit(limit int) ICmsMemberReportDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cmsMemberReportDo) Offset(offset int) ICmsMemberReportDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cmsMemberReportDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICmsMemberReportDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cmsMemberReportDo) Unscoped() ICmsMemberReportDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cmsMemberReportDo) Create(values ...*model.CmsMemberReport) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cmsMemberReportDo) CreateInBatches(values []*model.CmsMemberReport, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cmsMemberReportDo) Save(values ...*model.CmsMemberReport) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cmsMemberReportDo) First() (*model.CmsMemberReport, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsMemberReport), nil
	}
}

func (c cmsMemberReportDo) Take() (*model.CmsMemberReport, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsMemberReport), nil
	}
}

func (c cmsMemberReportDo) Last() (*model.CmsMemberReport, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsMemberReport), nil
	}
}

func (c cmsMemberReportDo) Find() ([]*model.CmsMemberReport, error) {
	result, err := c.DO.Find()
	return result.([]*model.CmsMemberReport), err
}

func (c cmsMemberReportDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CmsMemberReport, err error) {
	buf := make([]*model.CmsMemberReport, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cmsMemberReportDo) FindInBatches(result *[]*model.CmsMemberReport, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cmsMemberReportDo) Attrs(attrs ...field.AssignExpr) ICmsMemberReportDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cmsMemberReportDo) Assign(attrs ...field.AssignExpr) ICmsMemberReportDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cmsMemberReportDo) Joins(fields ...field.RelationField) ICmsMemberReportDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cmsMemberReportDo) Preload(fields ...field.RelationField) ICmsMemberReportDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cmsMemberReportDo) FirstOrInit() (*model.CmsMemberReport, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsMemberReport), nil
	}
}

func (c cmsMemberReportDo) FirstOrCreate() (*model.CmsMemberReport, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsMemberReport), nil
	}
}

func (c cmsMemberReportDo) FindByPage(offset int, limit int) (result []*model.CmsMemberReport, count int64, err error) {
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

func (c cmsMemberReportDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cmsMemberReportDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cmsMemberReportDo) Delete(models ...*model.CmsMemberReport) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cmsMemberReportDo) withDO(do gen.Dao) *cmsMemberReportDo {
	c.DO = *do.(*gen.DO)
	return c
}
