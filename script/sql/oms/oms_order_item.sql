create table oms_order_item
(
    id                  bigint auto_increment
        primary key,
    order_id            bigint        not null comment '订单id',
    order_sn            varchar(64)   not null comment '订单编号',
    product_id          bigint        not null comment '商品id',
    product_pic         varchar(500)  not null comment '商品图片',
    product_name        varchar(200)  not null comment '商品名称',
    product_brand       varchar(200)  not null comment '商品品牌',
    product_sn          varchar(64)   not null comment '货号',
    product_price       bigint        not null comment '销售价格',
    product_quantity    int           not null comment '购买数量',
    product_sku_id      bigint        not null comment '商品sku编号',
    product_sku_code    varchar(50)   not null comment '商品sku条码',
    product_category_id bigint        not null comment '商品分类id',
    promotion_name      varchar(200)  not null comment '商品促销名称',
    promotion_amount    bigint        not null comment '商品促销分解金额',
    coupon_amount       bigint        not null comment '优惠券优惠分解金额',
    integration_amount  bigint        not null comment '积分优惠分解金额',
    real_amount         bigint        not null comment '该商品经过优惠后的分解金额',
    gift_integration    int default 0 not null,
    gift_growth         int default 0 not null,
    product_attr        varchar(500)  not null comment '商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]'
)
    comment '订单中所包含的商品';

INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (21, 12, '201809150101000001', 26,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20',
        '华为', '6946605', 3788.00, 1, 90, '201806070026001', 19, '单品促销', 200.00, 2.02, 0.00, 3585.98, 3788, 3788,
        '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (22, 12, '201809150101000001', 27,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '7437788',
        2699.00, 3, 98, '201808270027001', 19, '打折优惠：满3件，打7.50折', 674.75, 1.44, 0.00, 2022.81, 2699, 2699,
        '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (23, 12, '201809150101000001', 28,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米',
        '7437789', 649.00, 1, 102, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 57.60, 0.35, 0.00, 591.05,
        649, 649, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (24, 12, '201809150101000001', 28,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米',
        '7437789', 699.00, 1, 103, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 62.40, 0.37, 0.00, 636.23,
        649, 649, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (25, 12, '201809150101000001', 29,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5acc5248N6a5f81cd.jpg',
        'Apple iPhone 8 Plus', '苹果', '7437799', 5499.00, 1, 106, '201808270029001', 19, '无优惠', 0.00, 2.94, 0.00,
        5496.06, 5499, 5499, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (26, 13, '201809150102000002', 26,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20',
        '华为', '6946605', 3788.00, 1, 90, '201806070026001', 19, '单品促销', 200.00, 2.02, 0.00, 3585.98, 3788, 3788,
        '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (27, 13, '201809150102000002', 27,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '7437788',
        2699.00, 3, 98, '201808270027001', 19, '打折优惠：满3件，打7.50折', 674.75, 1.44, 0.00, 2022.81, 2699, 2699,
        '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (28, 13, '201809150102000002', 28,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米',
        '7437789', 649.00, 1, 102, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 57.60, 0.35, 0.00, 591.05,
        649, 649, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (29, 13, '201809150102000002', 28,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米',
        '7437789', 699.00, 1, 103, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 62.40, 0.37, 0.00, 636.23,
        649, 649, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (30, 13, '201809150102000002', 29,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5acc5248N6a5f81cd.jpg',
        'Apple iPhone 8 Plus', '苹果', '7437799', 5499.00, 1, 106, '201808270029001', 19, '无优惠', 0.00, 2.94, 0.00,
        5496.06, 5499, 5499, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (31, 14, '201809130101000001', 26,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20',
        '华为', '6946605', 3788.00, 1, 90, '201806070026001', 19, '单品促销', 200.00, 2.02, 0.00, 3585.98, 3788, 3788,
        '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (32, 14, '201809130101000001', 27,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '7437788',
        2699.00, 3, 98, '201808270027001', 19, '打折优惠：满3件，打7.50折', 674.75, 1.44, 0.00, 2022.81, 2699, 2699,
        '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (33, 14, '201809130101000001', 28,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米',
        '7437789', 649.00, 1, 102, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 57.60, 0.35, 0.00, 591.05,
        649, 649, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (34, 14, '201809130101000001', 28,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米',
        '7437789', 699.00, 1, 103, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 62.40, 0.37, 0.00, 636.23,
        649, 649, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (35, 14, '201809130101000001', 29,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5acc5248N6a5f81cd.jpg',
        'Apple iPhone 8 Plus', '苹果', '7437799', 5499.00, 1, 106, '201808270029001', 19, '无优惠', 0.00, 2.94, 0.00,
        5496.06, 5499, 5499, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (36, 15, '201809130101000001', 26,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20',
        '华为', '6946605', 3788.00, 1, 90, '201806070026001', 19, '单品促销', 200.00, 2.02, 0.00, 3585.98, 3788, 3788,
        '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (37, 15, '201809130101000001', 27,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '7437788',
        2699.00, 3, 98, '201808270027001', 19, '打折优惠：满3件，打7.50折', 674.75, 1.44, 0.00, 2022.81, 2699, 2699,
        '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (38, 15, '201809130101000001', 28,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米',
        '7437789', 649.00, 1, 102, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 57.60, 0.35, 0.00, 591.05,
        649, 649, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (39, 15, '201809130101000001', 28,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米',
        '7437789', 699.00, 1, 103, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 62.40, 0.37, 0.00, 636.23,
        649, 649, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (40, 15, '201809130101000001', 29,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5acc5248N6a5f81cd.jpg',
        'Apple iPhone 8 Plus', '苹果', '7437799', 5499.00, 1, 106, '201808270029001', 19, '无优惠', 0.00, 2.94, 0.00,
        5496.06, 5499, 5499, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (41, 16, '201809140101000001', 26,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20',
        '华为', '6946605', 3788.00, 1, 90, '201806070026001', 19, '单品促销', 200.00, 2.02, 0.00, 3585.98, 3788, 3788,
        '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (42, 16, '201809140101000001', 27,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '7437788',
        2699.00, 3, 98, '201808270027001', 19, '打折优惠：满3件，打7.50折', 674.75, 1.44, 0.00, 2022.81, 2699, 2699,
        '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (43, 16, '201809140101000001', 28,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米',
        '7437789', 649.00, 1, 102, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 57.60, 0.35, 0.00, 591.05,
        649, 649, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (44, 16, '201809140101000001', 28,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米',
        '7437789', 699.00, 1, 103, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 62.40, 0.37, 0.00, 636.23,
        649, 649, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (45, 16, '201809140101000001', 29,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5acc5248N6a5f81cd.jpg',
        'Apple iPhone 8 Plus', '苹果', '7437799', 5499.00, 1, 106, '201808270029001', 19, '无优惠', 0.00, 2.94, 0.00,
        5496.06, 5499, 5499, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (46, 27, '202002250100000001', 36,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5b19403eN9f0b3cb8.jpg',
        '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 100.00, 3, 163,
        '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 100.00, 0, 0, ' ');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (47, 27, '202002250100000001', 36,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5b19403eN9f0b3cb8.jpg',
        '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 120.00, 2, 164,
        '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 120.00, 0, 0, ' ');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (48, 28, '202002250100000002', 36,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5b19403eN9f0b3cb8.jpg',
        '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 100.00, 3, 163,
        '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 100.00, 0, 0, ' ');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (49, 28, '202002250100000002', 36,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5b19403eN9f0b3cb8.jpg',
        '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 120.00, 2, 164,
        '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 120.00, 0, 0, ' ');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (50, 29, '202002250100000003', 36,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5b19403eN9f0b3cb8.jpg',
        '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 100.00, 3, 163,
        '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 100.00, 0, 0,
        '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (51, 29, '202002250100000003', 36,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5b19403eN9f0b3cb8.jpg',
        '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 120.00, 2, 164,
        '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 120.00, 0, 0,
        '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
INSERT INTO oms_order_item (id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn,
                            product_price, product_quantity, product_sku_id, product_sku_code, product_category_id,
                            promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount,
                            gift_integration, gift_growth, product_attr)
VALUES (52, 30, '202002250100000004', 36,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5b19403eN9f0b3cb8.jpg',
        '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 120.00, 2, 164,
        '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 120.00, 0, 0,
        '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]');
