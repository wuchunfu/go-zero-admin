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

	"github.com/feihua/zero-admin/rpc/oms/gen/model"
)

func newOmsCartItem(db *gorm.DB, opts ...gen.DOOption) omsCartItem {
	_omsCartItem := omsCartItem{}

	_omsCartItem.omsCartItemDo.UseDB(db, opts...)
	_omsCartItem.omsCartItemDo.UseModel(&model.OmsCartItem{})

	tableName := _omsCartItem.omsCartItemDo.TableName()
	_omsCartItem.ALL = field.NewAsterisk(tableName)
	_omsCartItem.ID = field.NewInt64(tableName, "id")
	_omsCartItem.ProductID = field.NewInt64(tableName, "product_id")
	_omsCartItem.ProductSkuID = field.NewInt64(tableName, "product_sku_id")
	_omsCartItem.MemberID = field.NewInt64(tableName, "member_id")
	_omsCartItem.Quantity = field.NewInt32(tableName, "quantity")
	_omsCartItem.Price = field.NewInt64(tableName, "price")
	_omsCartItem.ProductPic = field.NewString(tableName, "product_pic")
	_omsCartItem.ProductName = field.NewString(tableName, "product_name")
	_omsCartItem.ProductSubTitle = field.NewString(tableName, "product_sub_title")
	_omsCartItem.ProductSkuCode = field.NewString(tableName, "product_sku_code")
	_omsCartItem.MemberNickname = field.NewString(tableName, "member_nickname")
	_omsCartItem.CreateDate = field.NewTime(tableName, "create_date")
	_omsCartItem.UpdateDate = field.NewTime(tableName, "update_date")
	_omsCartItem.DeleteStatus = field.NewInt32(tableName, "delete_status")
	_omsCartItem.ProductCategoryID = field.NewInt64(tableName, "product_category_id")
	_omsCartItem.ProductBrand = field.NewString(tableName, "product_brand")
	_omsCartItem.ProductSn = field.NewString(tableName, "product_sn")
	_omsCartItem.ProductAttr = field.NewString(tableName, "product_attr")

	_omsCartItem.fillFieldMap()

	return _omsCartItem
}

// omsCartItem 购物车表
type omsCartItem struct {
	omsCartItemDo omsCartItemDo

	ALL               field.Asterisk
	ID                field.Int64
	ProductID         field.Int64  // 商品id
	ProductSkuID      field.Int64  // 商品库存id
	MemberID          field.Int64  // 会员id
	Quantity          field.Int32  // 购买数量
	Price             field.Int64  // 添加到购物车的价格
	ProductPic        field.String // 商品主图
	ProductName       field.String // 商品名称
	ProductSubTitle   field.String // 商品副标题（卖点）
	ProductSkuCode    field.String // 商品sku条码
	MemberNickname    field.String // 会员昵称
	CreateDate        field.Time   // 创建时间
	UpdateDate        field.Time   // 修改时间
	DeleteStatus      field.Int32  // 是否删除
	ProductCategoryID field.Int64  // 商品分类
	ProductBrand      field.String // 商品品牌
	ProductSn         field.String // 货号
	ProductAttr       field.String // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]

	fieldMap map[string]field.Expr
}

func (o omsCartItem) Table(newTableName string) *omsCartItem {
	o.omsCartItemDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o omsCartItem) As(alias string) *omsCartItem {
	o.omsCartItemDo.DO = *(o.omsCartItemDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *omsCartItem) updateTableName(table string) *omsCartItem {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.ProductID = field.NewInt64(table, "product_id")
	o.ProductSkuID = field.NewInt64(table, "product_sku_id")
	o.MemberID = field.NewInt64(table, "member_id")
	o.Quantity = field.NewInt32(table, "quantity")
	o.Price = field.NewInt64(table, "price")
	o.ProductPic = field.NewString(table, "product_pic")
	o.ProductName = field.NewString(table, "product_name")
	o.ProductSubTitle = field.NewString(table, "product_sub_title")
	o.ProductSkuCode = field.NewString(table, "product_sku_code")
	o.MemberNickname = field.NewString(table, "member_nickname")
	o.CreateDate = field.NewTime(table, "create_date")
	o.UpdateDate = field.NewTime(table, "update_date")
	o.DeleteStatus = field.NewInt32(table, "delete_status")
	o.ProductCategoryID = field.NewInt64(table, "product_category_id")
	o.ProductBrand = field.NewString(table, "product_brand")
	o.ProductSn = field.NewString(table, "product_sn")
	o.ProductAttr = field.NewString(table, "product_attr")

	o.fillFieldMap()

	return o
}

func (o *omsCartItem) WithContext(ctx context.Context) IOmsCartItemDo {
	return o.omsCartItemDo.WithContext(ctx)
}

func (o omsCartItem) TableName() string { return o.omsCartItemDo.TableName() }

func (o omsCartItem) Alias() string { return o.omsCartItemDo.Alias() }

func (o omsCartItem) Columns(cols ...field.Expr) gen.Columns { return o.omsCartItemDo.Columns(cols...) }

func (o *omsCartItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *omsCartItem) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 18)
	o.fieldMap["id"] = o.ID
	o.fieldMap["product_id"] = o.ProductID
	o.fieldMap["product_sku_id"] = o.ProductSkuID
	o.fieldMap["member_id"] = o.MemberID
	o.fieldMap["quantity"] = o.Quantity
	o.fieldMap["price"] = o.Price
	o.fieldMap["product_pic"] = o.ProductPic
	o.fieldMap["product_name"] = o.ProductName
	o.fieldMap["product_sub_title"] = o.ProductSubTitle
	o.fieldMap["product_sku_code"] = o.ProductSkuCode
	o.fieldMap["member_nickname"] = o.MemberNickname
	o.fieldMap["create_date"] = o.CreateDate
	o.fieldMap["update_date"] = o.UpdateDate
	o.fieldMap["delete_status"] = o.DeleteStatus
	o.fieldMap["product_category_id"] = o.ProductCategoryID
	o.fieldMap["product_brand"] = o.ProductBrand
	o.fieldMap["product_sn"] = o.ProductSn
	o.fieldMap["product_attr"] = o.ProductAttr
}

func (o omsCartItem) clone(db *gorm.DB) omsCartItem {
	o.omsCartItemDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o omsCartItem) replaceDB(db *gorm.DB) omsCartItem {
	o.omsCartItemDo.ReplaceDB(db)
	return o
}

type omsCartItemDo struct{ gen.DO }

type IOmsCartItemDo interface {
	gen.SubQuery
	Debug() IOmsCartItemDo
	WithContext(ctx context.Context) IOmsCartItemDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOmsCartItemDo
	WriteDB() IOmsCartItemDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOmsCartItemDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOmsCartItemDo
	Not(conds ...gen.Condition) IOmsCartItemDo
	Or(conds ...gen.Condition) IOmsCartItemDo
	Select(conds ...field.Expr) IOmsCartItemDo
	Where(conds ...gen.Condition) IOmsCartItemDo
	Order(conds ...field.Expr) IOmsCartItemDo
	Distinct(cols ...field.Expr) IOmsCartItemDo
	Omit(cols ...field.Expr) IOmsCartItemDo
	Join(table schema.Tabler, on ...field.Expr) IOmsCartItemDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOmsCartItemDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOmsCartItemDo
	Group(cols ...field.Expr) IOmsCartItemDo
	Having(conds ...gen.Condition) IOmsCartItemDo
	Limit(limit int) IOmsCartItemDo
	Offset(offset int) IOmsCartItemDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsCartItemDo
	Unscoped() IOmsCartItemDo
	Create(values ...*model.OmsCartItem) error
	CreateInBatches(values []*model.OmsCartItem, batchSize int) error
	Save(values ...*model.OmsCartItem) error
	First() (*model.OmsCartItem, error)
	Take() (*model.OmsCartItem, error)
	Last() (*model.OmsCartItem, error)
	Find() ([]*model.OmsCartItem, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsCartItem, err error)
	FindInBatches(result *[]*model.OmsCartItem, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OmsCartItem) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOmsCartItemDo
	Assign(attrs ...field.AssignExpr) IOmsCartItemDo
	Joins(fields ...field.RelationField) IOmsCartItemDo
	Preload(fields ...field.RelationField) IOmsCartItemDo
	FirstOrInit() (*model.OmsCartItem, error)
	FirstOrCreate() (*model.OmsCartItem, error)
	FindByPage(offset int, limit int) (result []*model.OmsCartItem, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOmsCartItemDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o omsCartItemDo) Debug() IOmsCartItemDo {
	return o.withDO(o.DO.Debug())
}

func (o omsCartItemDo) WithContext(ctx context.Context) IOmsCartItemDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o omsCartItemDo) ReadDB() IOmsCartItemDo {
	return o.Clauses(dbresolver.Read)
}

func (o omsCartItemDo) WriteDB() IOmsCartItemDo {
	return o.Clauses(dbresolver.Write)
}

func (o omsCartItemDo) Session(config *gorm.Session) IOmsCartItemDo {
	return o.withDO(o.DO.Session(config))
}

func (o omsCartItemDo) Clauses(conds ...clause.Expression) IOmsCartItemDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o omsCartItemDo) Returning(value interface{}, columns ...string) IOmsCartItemDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o omsCartItemDo) Not(conds ...gen.Condition) IOmsCartItemDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o omsCartItemDo) Or(conds ...gen.Condition) IOmsCartItemDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o omsCartItemDo) Select(conds ...field.Expr) IOmsCartItemDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o omsCartItemDo) Where(conds ...gen.Condition) IOmsCartItemDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o omsCartItemDo) Order(conds ...field.Expr) IOmsCartItemDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o omsCartItemDo) Distinct(cols ...field.Expr) IOmsCartItemDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o omsCartItemDo) Omit(cols ...field.Expr) IOmsCartItemDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o omsCartItemDo) Join(table schema.Tabler, on ...field.Expr) IOmsCartItemDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o omsCartItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOmsCartItemDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o omsCartItemDo) RightJoin(table schema.Tabler, on ...field.Expr) IOmsCartItemDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o omsCartItemDo) Group(cols ...field.Expr) IOmsCartItemDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o omsCartItemDo) Having(conds ...gen.Condition) IOmsCartItemDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o omsCartItemDo) Limit(limit int) IOmsCartItemDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o omsCartItemDo) Offset(offset int) IOmsCartItemDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o omsCartItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsCartItemDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o omsCartItemDo) Unscoped() IOmsCartItemDo {
	return o.withDO(o.DO.Unscoped())
}

func (o omsCartItemDo) Create(values ...*model.OmsCartItem) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o omsCartItemDo) CreateInBatches(values []*model.OmsCartItem, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o omsCartItemDo) Save(values ...*model.OmsCartItem) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o omsCartItemDo) First() (*model.OmsCartItem, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsCartItem), nil
	}
}

func (o omsCartItemDo) Take() (*model.OmsCartItem, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsCartItem), nil
	}
}

func (o omsCartItemDo) Last() (*model.OmsCartItem, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsCartItem), nil
	}
}

func (o omsCartItemDo) Find() ([]*model.OmsCartItem, error) {
	result, err := o.DO.Find()
	return result.([]*model.OmsCartItem), err
}

func (o omsCartItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsCartItem, err error) {
	buf := make([]*model.OmsCartItem, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o omsCartItemDo) FindInBatches(result *[]*model.OmsCartItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o omsCartItemDo) Attrs(attrs ...field.AssignExpr) IOmsCartItemDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o omsCartItemDo) Assign(attrs ...field.AssignExpr) IOmsCartItemDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o omsCartItemDo) Joins(fields ...field.RelationField) IOmsCartItemDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o omsCartItemDo) Preload(fields ...field.RelationField) IOmsCartItemDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o omsCartItemDo) FirstOrInit() (*model.OmsCartItem, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsCartItem), nil
	}
}

func (o omsCartItemDo) FirstOrCreate() (*model.OmsCartItem, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsCartItem), nil
	}
}

func (o omsCartItemDo) FindByPage(offset int, limit int) (result []*model.OmsCartItem, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o omsCartItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o omsCartItemDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o omsCartItemDo) Delete(models ...*model.OmsCartItem) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *omsCartItemDo) withDO(do gen.Dao) *omsCartItemDo {
	o.DO = *do.(*gen.DO)
	return o
}
