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

func newUmsMemberLevel(db *gorm.DB, opts ...gen.DOOption) umsMemberLevel {
	_umsMemberLevel := umsMemberLevel{}

	_umsMemberLevel.umsMemberLevelDo.UseDB(db, opts...)
	_umsMemberLevel.umsMemberLevelDo.UseModel(&model.UmsMemberLevel{})

	tableName := _umsMemberLevel.umsMemberLevelDo.TableName()
	_umsMemberLevel.ALL = field.NewAsterisk(tableName)
	_umsMemberLevel.ID = field.NewInt64(tableName, "id")
	_umsMemberLevel.LevelName = field.NewString(tableName, "level_name")
	_umsMemberLevel.GrowthPoint = field.NewInt32(tableName, "growth_point")
	_umsMemberLevel.DefaultStatus = field.NewInt32(tableName, "default_status")
	_umsMemberLevel.FreeFreightPoint = field.NewInt64(tableName, "free_freight_point")
	_umsMemberLevel.CommentGrowthPoint = field.NewInt32(tableName, "comment_growth_point")
	_umsMemberLevel.IsFreeFreight = field.NewInt32(tableName, "is_free_freight")
	_umsMemberLevel.IsSignIn = field.NewInt32(tableName, "is_sign_in")
	_umsMemberLevel.IsComment = field.NewInt32(tableName, "is_comment")
	_umsMemberLevel.IsPromotion = field.NewInt32(tableName, "is_promotion")
	_umsMemberLevel.IsMemberPrice = field.NewInt32(tableName, "is_member_price")
	_umsMemberLevel.IsBirthday = field.NewInt32(tableName, "is_birthday")
	_umsMemberLevel.Remark = field.NewString(tableName, "remark")

	_umsMemberLevel.fillFieldMap()

	return _umsMemberLevel
}

// umsMemberLevel 会员等级表
type umsMemberLevel struct {
	umsMemberLevelDo umsMemberLevelDo

	ALL                field.Asterisk
	ID                 field.Int64
	LevelName          field.String // 等级名称
	GrowthPoint        field.Int32  // 成长点
	DefaultStatus      field.Int32  // 是否为默认等级：0->不是；1->是
	FreeFreightPoint   field.Int64  // 免运费标准
	CommentGrowthPoint field.Int32  // 每次评价获取的成长值
	IsFreeFreight      field.Int32  // 是否有免邮特权
	IsSignIn           field.Int32  // 是否有签到特权
	IsComment          field.Int32  // 是否有评论获奖励特权
	IsPromotion        field.Int32  // 是否有专享活动特权
	IsMemberPrice      field.Int32  // 是否有会员价格特权
	IsBirthday         field.Int32  // 是否有生日特权
	Remark             field.String // 备注

	fieldMap map[string]field.Expr
}

func (u umsMemberLevel) Table(newTableName string) *umsMemberLevel {
	u.umsMemberLevelDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u umsMemberLevel) As(alias string) *umsMemberLevel {
	u.umsMemberLevelDo.DO = *(u.umsMemberLevelDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *umsMemberLevel) updateTableName(table string) *umsMemberLevel {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.LevelName = field.NewString(table, "level_name")
	u.GrowthPoint = field.NewInt32(table, "growth_point")
	u.DefaultStatus = field.NewInt32(table, "default_status")
	u.FreeFreightPoint = field.NewInt64(table, "free_freight_point")
	u.CommentGrowthPoint = field.NewInt32(table, "comment_growth_point")
	u.IsFreeFreight = field.NewInt32(table, "is_free_freight")
	u.IsSignIn = field.NewInt32(table, "is_sign_in")
	u.IsComment = field.NewInt32(table, "is_comment")
	u.IsPromotion = field.NewInt32(table, "is_promotion")
	u.IsMemberPrice = field.NewInt32(table, "is_member_price")
	u.IsBirthday = field.NewInt32(table, "is_birthday")
	u.Remark = field.NewString(table, "remark")

	u.fillFieldMap()

	return u
}

func (u *umsMemberLevel) WithContext(ctx context.Context) IUmsMemberLevelDo {
	return u.umsMemberLevelDo.WithContext(ctx)
}

func (u umsMemberLevel) TableName() string { return u.umsMemberLevelDo.TableName() }

func (u umsMemberLevel) Alias() string { return u.umsMemberLevelDo.Alias() }

func (u umsMemberLevel) Columns(cols ...field.Expr) gen.Columns {
	return u.umsMemberLevelDo.Columns(cols...)
}

func (u *umsMemberLevel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *umsMemberLevel) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 13)
	u.fieldMap["id"] = u.ID
	u.fieldMap["level_name"] = u.LevelName
	u.fieldMap["growth_point"] = u.GrowthPoint
	u.fieldMap["default_status"] = u.DefaultStatus
	u.fieldMap["free_freight_point"] = u.FreeFreightPoint
	u.fieldMap["comment_growth_point"] = u.CommentGrowthPoint
	u.fieldMap["is_free_freight"] = u.IsFreeFreight
	u.fieldMap["is_sign_in"] = u.IsSignIn
	u.fieldMap["is_comment"] = u.IsComment
	u.fieldMap["is_promotion"] = u.IsPromotion
	u.fieldMap["is_member_price"] = u.IsMemberPrice
	u.fieldMap["is_birthday"] = u.IsBirthday
	u.fieldMap["remark"] = u.Remark
}

func (u umsMemberLevel) clone(db *gorm.DB) umsMemberLevel {
	u.umsMemberLevelDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u umsMemberLevel) replaceDB(db *gorm.DB) umsMemberLevel {
	u.umsMemberLevelDo.ReplaceDB(db)
	return u
}

type umsMemberLevelDo struct{ gen.DO }

type IUmsMemberLevelDo interface {
	gen.SubQuery
	Debug() IUmsMemberLevelDo
	WithContext(ctx context.Context) IUmsMemberLevelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUmsMemberLevelDo
	WriteDB() IUmsMemberLevelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUmsMemberLevelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUmsMemberLevelDo
	Not(conds ...gen.Condition) IUmsMemberLevelDo
	Or(conds ...gen.Condition) IUmsMemberLevelDo
	Select(conds ...field.Expr) IUmsMemberLevelDo
	Where(conds ...gen.Condition) IUmsMemberLevelDo
	Order(conds ...field.Expr) IUmsMemberLevelDo
	Distinct(cols ...field.Expr) IUmsMemberLevelDo
	Omit(cols ...field.Expr) IUmsMemberLevelDo
	Join(table schema.Tabler, on ...field.Expr) IUmsMemberLevelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberLevelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberLevelDo
	Group(cols ...field.Expr) IUmsMemberLevelDo
	Having(conds ...gen.Condition) IUmsMemberLevelDo
	Limit(limit int) IUmsMemberLevelDo
	Offset(offset int) IUmsMemberLevelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberLevelDo
	Unscoped() IUmsMemberLevelDo
	Create(values ...*model.UmsMemberLevel) error
	CreateInBatches(values []*model.UmsMemberLevel, batchSize int) error
	Save(values ...*model.UmsMemberLevel) error
	First() (*model.UmsMemberLevel, error)
	Take() (*model.UmsMemberLevel, error)
	Last() (*model.UmsMemberLevel, error)
	Find() ([]*model.UmsMemberLevel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMemberLevel, err error)
	FindInBatches(result *[]*model.UmsMemberLevel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UmsMemberLevel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUmsMemberLevelDo
	Assign(attrs ...field.AssignExpr) IUmsMemberLevelDo
	Joins(fields ...field.RelationField) IUmsMemberLevelDo
	Preload(fields ...field.RelationField) IUmsMemberLevelDo
	FirstOrInit() (*model.UmsMemberLevel, error)
	FirstOrCreate() (*model.UmsMemberLevel, error)
	FindByPage(offset int, limit int) (result []*model.UmsMemberLevel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUmsMemberLevelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u umsMemberLevelDo) Debug() IUmsMemberLevelDo {
	return u.withDO(u.DO.Debug())
}

func (u umsMemberLevelDo) WithContext(ctx context.Context) IUmsMemberLevelDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u umsMemberLevelDo) ReadDB() IUmsMemberLevelDo {
	return u.Clauses(dbresolver.Read)
}

func (u umsMemberLevelDo) WriteDB() IUmsMemberLevelDo {
	return u.Clauses(dbresolver.Write)
}

func (u umsMemberLevelDo) Session(config *gorm.Session) IUmsMemberLevelDo {
	return u.withDO(u.DO.Session(config))
}

func (u umsMemberLevelDo) Clauses(conds ...clause.Expression) IUmsMemberLevelDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u umsMemberLevelDo) Returning(value interface{}, columns ...string) IUmsMemberLevelDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u umsMemberLevelDo) Not(conds ...gen.Condition) IUmsMemberLevelDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u umsMemberLevelDo) Or(conds ...gen.Condition) IUmsMemberLevelDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u umsMemberLevelDo) Select(conds ...field.Expr) IUmsMemberLevelDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u umsMemberLevelDo) Where(conds ...gen.Condition) IUmsMemberLevelDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u umsMemberLevelDo) Order(conds ...field.Expr) IUmsMemberLevelDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u umsMemberLevelDo) Distinct(cols ...field.Expr) IUmsMemberLevelDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u umsMemberLevelDo) Omit(cols ...field.Expr) IUmsMemberLevelDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u umsMemberLevelDo) Join(table schema.Tabler, on ...field.Expr) IUmsMemberLevelDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u umsMemberLevelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberLevelDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u umsMemberLevelDo) RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberLevelDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u umsMemberLevelDo) Group(cols ...field.Expr) IUmsMemberLevelDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u umsMemberLevelDo) Having(conds ...gen.Condition) IUmsMemberLevelDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u umsMemberLevelDo) Limit(limit int) IUmsMemberLevelDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u umsMemberLevelDo) Offset(offset int) IUmsMemberLevelDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u umsMemberLevelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberLevelDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u umsMemberLevelDo) Unscoped() IUmsMemberLevelDo {
	return u.withDO(u.DO.Unscoped())
}

func (u umsMemberLevelDo) Create(values ...*model.UmsMemberLevel) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u umsMemberLevelDo) CreateInBatches(values []*model.UmsMemberLevel, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u umsMemberLevelDo) Save(values ...*model.UmsMemberLevel) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u umsMemberLevelDo) First() (*model.UmsMemberLevel, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberLevel), nil
	}
}

func (u umsMemberLevelDo) Take() (*model.UmsMemberLevel, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberLevel), nil
	}
}

func (u umsMemberLevelDo) Last() (*model.UmsMemberLevel, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberLevel), nil
	}
}

func (u umsMemberLevelDo) Find() ([]*model.UmsMemberLevel, error) {
	result, err := u.DO.Find()
	return result.([]*model.UmsMemberLevel), err
}

func (u umsMemberLevelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMemberLevel, err error) {
	buf := make([]*model.UmsMemberLevel, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u umsMemberLevelDo) FindInBatches(result *[]*model.UmsMemberLevel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u umsMemberLevelDo) Attrs(attrs ...field.AssignExpr) IUmsMemberLevelDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u umsMemberLevelDo) Assign(attrs ...field.AssignExpr) IUmsMemberLevelDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u umsMemberLevelDo) Joins(fields ...field.RelationField) IUmsMemberLevelDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u umsMemberLevelDo) Preload(fields ...field.RelationField) IUmsMemberLevelDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u umsMemberLevelDo) FirstOrInit() (*model.UmsMemberLevel, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberLevel), nil
	}
}

func (u umsMemberLevelDo) FirstOrCreate() (*model.UmsMemberLevel, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberLevel), nil
	}
}

func (u umsMemberLevelDo) FindByPage(offset int, limit int) (result []*model.UmsMemberLevel, count int64, err error) {
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

func (u umsMemberLevelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u umsMemberLevelDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u umsMemberLevelDo) Delete(models ...*model.UmsMemberLevel) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *umsMemberLevelDo) withDO(do gen.Dao) *umsMemberLevelDo {
	u.DO = *do.(*gen.DO)
	return u
}
