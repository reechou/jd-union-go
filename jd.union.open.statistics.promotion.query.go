package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenStatisticsPromotionQueryTopLevel struct {
	JdUnionOpenStatisticsPromotionQueryResponse JdUnionOpenStatisticsPromotionQueryResponse `json:"jd_union_open_statistics_Promotion_query_responce"`
}

type JdUnionOpenStatisticsPromotionQueryResponse struct {
	Result string `json:"queryResult"`
	Code   string `json:"code"`
}

type JdUnionOpenStatisticsPromotionQueryResult struct {
	Code      int64                  `json:"code"`
	Message   string                 `json:"message"`
	Data      []*PromotionEffectData `json:"data"`
	RequestID string                 `json:"requestId"`
}

type PromotionEffectDataResp struct {
	PromotionEffectDataResp *PromotionEffectData `json:"promotionEffectDataResp"`
}

type PromotionEffectData struct {
	UnionId             int64  `json:"unionId"`
	SkuId               int64  `json:"skuId"`
	ActivityUrl         string `json:"activityUrl"`
	TimeType            int    `json:"timeType"`
	DataType            int    `json:"dataType"`
	Time                string `json:"time"`
	ClickPv             int64  `json:"clickPv"`             // 点击量
	EstimateValidOrders int64  `json:"estimateValidOrders"` // 有效订单量
	EstimateValidFee    int64  `json:"estimateValidFee"`    // 预估收入
	EstimateValidGmv    int64  `json:"estimateValidGmv"`    // 有效订单金额
	RefundOrders        int64  `json:"refundOrders"`        // 退款订单量： 当日下单付款后又取消的订单量
	CompleteOrders      int64  `json:"completeOrders"`      // 完成订单量
	CompleteGmv         int64  `json:"completeGmv"`         // 完成订单金额
	ActualFee           int64  `json:"actualFee"`           // 实际收入
	ItemId              string `json:"itemId"`
}

func (app *App) JdUnionOpenStatisticsPromotionQuery(params map[string]interface{}) (result *JdUnionOpenStatisticsPromotionQueryResult, err error) {
	body, err := app.Request("jd.union.open.statistics.promotion.query", map[string]interface{}{"promotionEffectDataReq": params})
	resp := &JdUnionOpenStatisticsPromotionQueryTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenStatisticsPromotionQueryResponse.Result != "" {
		result = &JdUnionOpenStatisticsPromotionQueryResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenStatisticsPromotionQueryResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
