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

func newOmsOrderItem(db *gorm.DB, opts ...gen.DOOption) omsOrderItem {
	_omsOrderItem := omsOrderItem{}

	_omsOrderItem.omsOrderItemDo.UseDB(db, opts...)
	_omsOrderItem.omsOrderItemDo.UseModel(&model.OmsOrderItem{})

	tableName := _omsOrderItem.omsOrderItemDo.TableName()
	_omsOrderItem.ALL = field.NewAsterisk(tableName)
	_omsOrderItem.ID = field.NewInt64(tableName, "id")
	_omsOrderItem.OrderID = field.NewInt64(tableName, "order_id")
	_omsOrderItem.OrderSn = field.NewString(tableName, "order_sn")
	_omsOrderItem.ProductID = field.NewInt64(tableName, "product_id")
	_omsOrderItem.ProductPic = field.NewString(tableName, "product_pic")
	_omsOrderItem.ProductName = field.NewString(tableName, "product_name")
	_omsOrderItem.ProductBrand = field.NewString(tableName, "product_brand")
	_omsOrderItem.ProductSn = field.NewString(tableName, "product_sn")
	_omsOrderItem.ProductPrice = field.NewInt64(tableName, "product_price")
	_omsOrderItem.ProductQuantity = field.NewInt32(tableName, "product_quantity")
	_omsOrderItem.ProductSkuID = field.NewInt64(tableName, "product_sku_id")
	_omsOrderItem.ProductSkuCode = field.NewString(tableName, "product_sku_code")
	_omsOrderItem.ProductCategoryID = field.NewInt64(tableName, "product_category_id")
	_omsOrderItem.PromotionName = field.NewString(tableName, "promotion_name")
	_omsOrderItem.PromotionAmount = field.NewInt64(tableName, "promotion_amount")
	_omsOrderItem.CouponAmount = field.NewInt64(tableName, "coupon_amount")
	_omsOrderItem.IntegrationAmount = field.NewInt64(tableName, "integration_amount")
	_omsOrderItem.RealAmount = field.NewInt64(tableName, "real_amount")
	_omsOrderItem.GiftIntegration = field.NewInt32(tableName, "gift_integration")
	_omsOrderItem.GiftGrowth = field.NewInt32(tableName, "gift_growth")
	_omsOrderItem.ProductAttr = field.NewString(tableName, "product_attr")

	_omsOrderItem.fillFieldMap()

	return _omsOrderItem
}

// omsOrderItem 订单中所包含的商品
type omsOrderItem struct {
	omsOrderItemDo omsOrderItemDo

	ALL               field.Asterisk
	ID                field.Int64
	OrderID           field.Int64  // 订单id
	OrderSn           field.String // 订单编号
	ProductID         field.Int64  // 商品id
	ProductPic        field.String // 商品图片
	ProductName       field.String // 商品名称
	ProductBrand      field.String // 商品品牌
	ProductSn         field.String // 货号
	ProductPrice      field.Int64  // 销售价格
	ProductQuantity   field.Int32  // 购买数量
	ProductSkuID      field.Int64  // 商品sku编号
	ProductSkuCode    field.String // 商品sku条码
	ProductCategoryID field.Int64  // 商品分类id
	PromotionName     field.String // 商品促销名称
	PromotionAmount   field.Int64  // 商品促销分解金额
	CouponAmount      field.Int64  // 优惠券优惠分解金额
	IntegrationAmount field.Int64  // 积分优惠分解金额
	RealAmount        field.Int64  // 该商品经过优惠后的分解金额
	GiftIntegration   field.Int32
	GiftGrowth        field.Int32
	ProductAttr       field.String // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]

	fieldMap map[string]field.Expr
}

func (o omsOrderItem) Table(newTableName string) *omsOrderItem {
	o.omsOrderItemDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o omsOrderItem) As(alias string) *omsOrderItem {
	o.omsOrderItemDo.DO = *(o.omsOrderItemDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *omsOrderItem) updateTableName(table string) *omsOrderItem {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.OrderID = field.NewInt64(table, "order_id")
	o.OrderSn = field.NewString(table, "order_sn")
	o.ProductID = field.NewInt64(table, "product_id")
	o.ProductPic = field.NewString(table, "product_pic")
	o.ProductName = field.NewString(table, "product_name")
	o.ProductBrand = field.NewString(table, "product_brand")
	o.ProductSn = field.NewString(table, "product_sn")
	o.ProductPrice = field.NewInt64(table, "product_price")
	o.ProductQuantity = field.NewInt32(table, "product_quantity")
	o.ProductSkuID = field.NewInt64(table, "product_sku_id")
	o.ProductSkuCode = field.NewString(table, "product_sku_code")
	o.ProductCategoryID = field.NewInt64(table, "product_category_id")
	o.PromotionName = field.NewString(table, "promotion_name")
	o.PromotionAmount = field.NewInt64(table, "promotion_amount")
	o.CouponAmount = field.NewInt64(table, "coupon_amount")
	o.IntegrationAmount = field.NewInt64(table, "integration_amount")
	o.RealAmount = field.NewInt64(table, "real_amount")
	o.GiftIntegration = field.NewInt32(table, "gift_integration")
	o.GiftGrowth = field.NewInt32(table, "gift_growth")
	o.ProductAttr = field.NewString(table, "product_attr")

	o.fillFieldMap()

	return o
}

func (o *omsOrderItem) WithContext(ctx context.Context) IOmsOrderItemDo {
	return o.omsOrderItemDo.WithContext(ctx)
}

func (o omsOrderItem) TableName() string { return o.omsOrderItemDo.TableName() }

func (o omsOrderItem) Alias() string { return o.omsOrderItemDo.Alias() }

func (o omsOrderItem) Columns(cols ...field.Expr) gen.Columns {
	return o.omsOrderItemDo.Columns(cols...)
}

func (o *omsOrderItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *omsOrderItem) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 21)
	o.fieldMap["id"] = o.ID
	o.fieldMap["order_id"] = o.OrderID
	o.fieldMap["order_sn"] = o.OrderSn
	o.fieldMap["product_id"] = o.ProductID
	o.fieldMap["product_pic"] = o.ProductPic
	o.fieldMap["product_name"] = o.ProductName
	o.fieldMap["product_brand"] = o.ProductBrand
	o.fieldMap["product_sn"] = o.ProductSn
	o.fieldMap["product_price"] = o.ProductPrice
	o.fieldMap["product_quantity"] = o.ProductQuantity
	o.fieldMap["product_sku_id"] = o.ProductSkuID
	o.fieldMap["product_sku_code"] = o.ProductSkuCode
	o.fieldMap["product_category_id"] = o.ProductCategoryID
	o.fieldMap["promotion_name"] = o.PromotionName
	o.fieldMap["promotion_amount"] = o.PromotionAmount
	o.fieldMap["coupon_amount"] = o.CouponAmount
	o.fieldMap["integration_amount"] = o.IntegrationAmount
	o.fieldMap["real_amount"] = o.RealAmount
	o.fieldMap["gift_integration"] = o.GiftIntegration
	o.fieldMap["gift_growth"] = o.GiftGrowth
	o.fieldMap["product_attr"] = o.ProductAttr
}

func (o omsOrderItem) clone(db *gorm.DB) omsOrderItem {
	o.omsOrderItemDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o omsOrderItem) replaceDB(db *gorm.DB) omsOrderItem {
	o.omsOrderItemDo.ReplaceDB(db)
	return o
}

type omsOrderItemDo struct{ gen.DO }

type IOmsOrderItemDo interface {
	gen.SubQuery
	Debug() IOmsOrderItemDo
	WithContext(ctx context.Context) IOmsOrderItemDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOmsOrderItemDo
	WriteDB() IOmsOrderItemDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOmsOrderItemDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOmsOrderItemDo
	Not(conds ...gen.Condition) IOmsOrderItemDo
	Or(conds ...gen.Condition) IOmsOrderItemDo
	Select(conds ...field.Expr) IOmsOrderItemDo
	Where(conds ...gen.Condition) IOmsOrderItemDo
	Order(conds ...field.Expr) IOmsOrderItemDo
	Distinct(cols ...field.Expr) IOmsOrderItemDo
	Omit(cols ...field.Expr) IOmsOrderItemDo
	Join(table schema.Tabler, on ...field.Expr) IOmsOrderItemDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderItemDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderItemDo
	Group(cols ...field.Expr) IOmsOrderItemDo
	Having(conds ...gen.Condition) IOmsOrderItemDo
	Limit(limit int) IOmsOrderItemDo
	Offset(offset int) IOmsOrderItemDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderItemDo
	Unscoped() IOmsOrderItemDo
	Create(values ...*model.OmsOrderItem) error
	CreateInBatches(values []*model.OmsOrderItem, batchSize int) error
	Save(values ...*model.OmsOrderItem) error
	First() (*model.OmsOrderItem, error)
	Take() (*model.OmsOrderItem, error)
	Last() (*model.OmsOrderItem, error)
	Find() ([]*model.OmsOrderItem, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderItem, err error)
	FindInBatches(result *[]*model.OmsOrderItem, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OmsOrderItem) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOmsOrderItemDo
	Assign(attrs ...field.AssignExpr) IOmsOrderItemDo
	Joins(fields ...field.RelationField) IOmsOrderItemDo
	Preload(fields ...field.RelationField) IOmsOrderItemDo
	FirstOrInit() (*model.OmsOrderItem, error)
	FirstOrCreate() (*model.OmsOrderItem, error)
	FindByPage(offset int, limit int) (result []*model.OmsOrderItem, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOmsOrderItemDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o omsOrderItemDo) Debug() IOmsOrderItemDo {
	return o.withDO(o.DO.Debug())
}

func (o omsOrderItemDo) WithContext(ctx context.Context) IOmsOrderItemDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o omsOrderItemDo) ReadDB() IOmsOrderItemDo {
	return o.Clauses(dbresolver.Read)
}

func (o omsOrderItemDo) WriteDB() IOmsOrderItemDo {
	return o.Clauses(dbresolver.Write)
}

func (o omsOrderItemDo) Session(config *gorm.Session) IOmsOrderItemDo {
	return o.withDO(o.DO.Session(config))
}

func (o omsOrderItemDo) Clauses(conds ...clause.Expression) IOmsOrderItemDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o omsOrderItemDo) Returning(value interface{}, columns ...string) IOmsOrderItemDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o omsOrderItemDo) Not(conds ...gen.Condition) IOmsOrderItemDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o omsOrderItemDo) Or(conds ...gen.Condition) IOmsOrderItemDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o omsOrderItemDo) Select(conds ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o omsOrderItemDo) Where(conds ...gen.Condition) IOmsOrderItemDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o omsOrderItemDo) Order(conds ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o omsOrderItemDo) Distinct(cols ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o omsOrderItemDo) Omit(cols ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o omsOrderItemDo) Join(table schema.Tabler, on ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o omsOrderItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o omsOrderItemDo) RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o omsOrderItemDo) Group(cols ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o omsOrderItemDo) Having(conds ...gen.Condition) IOmsOrderItemDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o omsOrderItemDo) Limit(limit int) IOmsOrderItemDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o omsOrderItemDo) Offset(offset int) IOmsOrderItemDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o omsOrderItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderItemDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o omsOrderItemDo) Unscoped() IOmsOrderItemDo {
	return o.withDO(o.DO.Unscoped())
}

func (o omsOrderItemDo) Create(values ...*model.OmsOrderItem) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o omsOrderItemDo) CreateInBatches(values []*model.OmsOrderItem, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o omsOrderItemDo) Save(values ...*model.OmsOrderItem) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o omsOrderItemDo) First() (*model.OmsOrderItem, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderItem), nil
	}
}

func (o omsOrderItemDo) Take() (*model.OmsOrderItem, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderItem), nil
	}
}

func (o omsOrderItemDo) Last() (*model.OmsOrderItem, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderItem), nil
	}
}

func (o omsOrderItemDo) Find() ([]*model.OmsOrderItem, error) {
	result, err := o.DO.Find()
	return result.([]*model.OmsOrderItem), err
}

func (o omsOrderItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderItem, err error) {
	buf := make([]*model.OmsOrderItem, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o omsOrderItemDo) FindInBatches(result *[]*model.OmsOrderItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o omsOrderItemDo) Attrs(attrs ...field.AssignExpr) IOmsOrderItemDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o omsOrderItemDo) Assign(attrs ...field.AssignExpr) IOmsOrderItemDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o omsOrderItemDo) Joins(fields ...field.RelationField) IOmsOrderItemDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o omsOrderItemDo) Preload(fields ...field.RelationField) IOmsOrderItemDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o omsOrderItemDo) FirstOrInit() (*model.OmsOrderItem, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderItem), nil
	}
}

func (o omsOrderItemDo) FirstOrCreate() (*model.OmsOrderItem, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderItem), nil
	}
}

func (o omsOrderItemDo) FindByPage(offset int, limit int) (result []*model.OmsOrderItem, count int64, err error) {
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

func (o omsOrderItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o omsOrderItemDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o omsOrderItemDo) Delete(models ...*model.OmsOrderItem) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *omsOrderItemDo) withDO(do gen.Dao) *omsOrderItemDo {
	o.DO = *do.(*gen.DO)
	return o
}
