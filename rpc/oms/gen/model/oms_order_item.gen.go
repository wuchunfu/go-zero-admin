// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOmsOrderItem = "oms_order_item"

// OmsOrderItem 订单中所包含的商品
type OmsOrderItem struct {
	ID                int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OrderID           int64  `gorm:"column:order_id;not null;comment:订单id" json:"order_id"`                         // 订单id
	OrderSn           string `gorm:"column:order_sn;not null;comment:订单编号" json:"order_sn"`                         // 订单编号
	ProductID         int64  `gorm:"column:product_id;not null;comment:商品id" json:"product_id"`                     // 商品id
	ProductPic        string `gorm:"column:product_pic;not null;comment:商品图片" json:"product_pic"`                   // 商品图片
	ProductName       string `gorm:"column:product_name;not null;comment:商品名称" json:"product_name"`                 // 商品名称
	ProductBrand      string `gorm:"column:product_brand;not null;comment:商品品牌" json:"product_brand"`               // 商品品牌
	ProductSn         string `gorm:"column:product_sn;not null;comment:货号" json:"product_sn"`                       // 货号
	ProductPrice      int64  `gorm:"column:product_price;not null;comment:销售价格" json:"product_price"`               // 销售价格
	ProductQuantity   int32  `gorm:"column:product_quantity;not null;comment:购买数量" json:"product_quantity"`         // 购买数量
	ProductSkuID      int64  `gorm:"column:product_sku_id;not null;comment:商品sku编号" json:"product_sku_id"`          // 商品sku编号
	ProductSkuCode    string `gorm:"column:product_sku_code;not null;comment:商品sku条码" json:"product_sku_code"`      // 商品sku条码
	ProductCategoryID int64  `gorm:"column:product_category_id;not null;comment:商品分类id" json:"product_category_id"` // 商品分类id
	PromotionName     string `gorm:"column:promotion_name;not null;comment:商品促销名称" json:"promotion_name"`           // 商品促销名称
	PromotionAmount   int64  `gorm:"column:promotion_amount;not null;comment:商品促销分解金额" json:"promotion_amount"`     // 商品促销分解金额
	CouponAmount      int64  `gorm:"column:coupon_amount;not null;comment:优惠券优惠分解金额" json:"coupon_amount"`          // 优惠券优惠分解金额
	IntegrationAmount int64  `gorm:"column:integration_amount;not null;comment:积分优惠分解金额" json:"integration_amount"` // 积分优惠分解金额
	RealAmount        int64  `gorm:"column:real_amount;not null;comment:该商品经过优惠后的分解金额" json:"real_amount"`          // 该商品经过优惠后的分解金额
	GiftIntegration   int32  `gorm:"column:gift_integration;not null" json:"gift_integration"`
	GiftGrowth        int32  `gorm:"column:gift_growth;not null" json:"gift_growth"`
	ProductAttr       string `gorm:"column:product_attr;not null;comment:商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]" json:"product_attr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

// TableName OmsOrderItem's table name
func (*OmsOrderItem) TableName() string {
	return TableNameOmsOrderItem
}