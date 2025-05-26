package jd_union_go

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/shopspring/decimal"
)

type JdUnionOpenOrderRowQueryResponseTopLevel struct {
	JdUnionOpenOrderRowQueryResponse JdUnionOpenOrderRowQueryResponse `json:"jd_union_open_order_row_query_responce"`
}

type JdUnionOpenOrderRowQueryResponse struct {
	Result string `json:"queryResult"`
	Code   string `json:"code"`
}

type JdUnionOpenOrderRowQueryResult struct {
	Code      int64       `json:"code"`
	Data      []*OrderRow `json:"data"`
	HasMore   bool        `json:"hasMore"`
	Message   string      `json:"message"`
	RequestID string      `json:"requestId"`
}

type OrderRow struct {
	ActualCosPrice      decimal.Decimal    `json:"actualCosPrice"`
	ActualFee           decimal.Decimal    `json:"actualFee"`
	BalanceExt          string             `json:"balanceExt"`
	Cid1                int                `json:"cid1"`
	Cid2                int                `json:"cid2"`
	Cid3                int                `json:"cid3"`
	CommissionRate      decimal.Decimal    `json:"commissionRate"`
	CpActID             int                `json:"cpActId"`
	EstimateCosPrice    decimal.Decimal    `json:"estimateCosPrice"` // 预估计佣金额
	EstimateFee         decimal.Decimal    `json:"estimateFee"`
	ExpressStatus       int                `json:"expressStatus"` // 发货状态（10：待发货，20：已发货）
	Ext1                string             `json:"ext1"`
	FinalRate           decimal.Decimal    `json:"finalRate"`
	FinishTime          string             `json:"finishTime"` // 完成时间（购买用户确认收货时间）,格式yyyy-MM-dd HH:mm:ss
	GiftCouponKey       string             `json:"giftCouponKey"`
	GiftCouponOcsAmount decimal.Decimal    `json:"giftCouponOcsAmount"`
	GoodsInfo           *OrderRowGoodsInfo `json:"goodsInfo"`
	ID                  string             `json:"id"`
	ModifyTime          string             `json:"modifyTime"` // 更新时间,格式yyyy-MM-dd HH:mm:ss
	OrderEmt            int                `json:"orderEmt"`   // 下单设备 1.pc 2.无线
	OrderID             int64              `json:"orderId"`    // 订单号
	OrderTime           string             `json:"orderTime"`  // 下单时间,格式yyyy-MM-dd HH:mm:ss
	ParentID            int                `json:"parentId"`   // 主单的订单号
	PayMonth            int                `json:"payMonth"`
	Pid                 string             `json:"pid"`
	Plus                int                `json:"plus"`       // 下单用户是否为PLUS会员 0：否，1：是
	PopId               int                `json:"popId"`      // 商家ID
	PositionId          int                `json:"positionId"` // 推广位ID
	Price               decimal.Decimal    `json:"price"`
	ProPriceAmount      decimal.Decimal    `json:"proPriceAmount"`
	Rid                 int                `json:"rid"`
	SiteId              int                `json:"siteId"`
	SkuFrozenNum        int                `json:"skuFrozenNum"`
	SkuId               int                `json:"skuId"`
	SkuName             string             `json:"skuName"`
	SkuNum              int                `json:"skuNum"`
	SkuReturnNum        int                `json:"skuReturnNum"` // 商品已退货数量
	SubSideRate         decimal.Decimal    `json:"subSideRate"`  // 分成比例
	SubUnionID          string             `json:"subUnionId"`
	SubSidyRate         decimal.Decimal    `json:"subsidyRate"` // 补贴比例
	TraceType           int                `json:"traceType"`   // 同跨店：2同店 3跨店
	UnionAlias          string             `json:"unionAlias"`
	UnionId             int                `json:"unionId"`
	UnionRole           int                `json:"unionRole"`
	UnionTag            string             `json:"unionTag"`
	// sku维度的有效码（-1：未知,2.无效-拆单,3.无效-取消,4.无效-京东帮帮主订单,5.无效-账号异常,6.无效-赠品类目不返佣,7.无效-校园订单,8.无效-企业订单,9.无效-团购订单,11.无效-乡村推广员下单,13. 违规订单-其他,14.无效-来源与备案网址不符,15.待付款,16.已付款,17.已完成（购买用户确认收货）,19.无效-佣金比例为0,20.无效-此复购订单对应的首购订单无效,21.无效-云店订单,22.无效-PLUS会员佣金比例为0,23.无效-支付有礼,24.已付定金,25. 违规订单-流量劫持,26. 违规订单-流量异常,27. 违规订单-违反京东平台规则,28. 违规订单-多笔交易异常,29.无效-跨屏跨店,30.无效-累计件数超出类目上限,31.无效-黑名单sku,33.超市卡充值订单,34.无效-推卡订单无效
	ValidCode int `json:"validCode"`
}

type OrderRowGoodsInfo struct {
	ImageUrl  string `json:"imageUrl"`
	MainSkuId int    `json:"mainSkuId"`
	Owner     string `json:"owner"`
	ProductId int    `json:"productId"`
	ShopId    int    `json:"shopId"`
	ShopName  string `json:"shopName"`
}

func (app *App) JdUnionOpenOrderRowQuery(params map[string]interface{}) (result *JdUnionOpenOrderRowQueryResult, err error) {
	body, err := app.Request("jd.union.open.order.row.query", map[string]interface{}{"orderReq": params})

	resp := &JdUnionOpenOrderRowQueryResponseTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenOrderRowQueryResponse.Result != "" {
		result = &JdUnionOpenOrderRowQueryResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenOrderRowQueryResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
