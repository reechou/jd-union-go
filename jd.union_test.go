package jd_union_go

import (
	"log"
	"testing"
)

var app = &App{
	ID:     "xx",
	Key:    "xx",
	Secret: "xx",
	debug:  true,
}

// positionId 推广位 id
var positionId = 3101765874
var orderId = 315522556318

// 获取商品类目
func TestOpenCategoryGoodsGet(t *testing.T) {
	res, err := app.JdUnionOpenCategoryGoodsGet(map[string]interface{}{
		"parentId": 0,
		"grade":    0,
	})
	log.Println(res, err)
}

// 获取活动列表
func TestOpenActivityQuery(t *testing.T) {
	res, err := app.JdUnionOpenActivityQuery(map[string]interface{}{
		"pageIndex":  1,
		"pageSize":   50,
		"poolId":     6, //活动物料ID，1：营销日历热门会场；2：营销日历热门榜单；6：PC站长端官方活动
		"activeDate": "20210421",
	})
	log.Println(len(res.Data), res.TotalCount, err)
}

func TestOpenGoodsJingfenQuery(t *testing.T) {
	res, err := app.JdUnionOpenGoodsJingfenQuery(map[string]interface{}{
		"eliteId":   1,
		"sortName":  "price",
		"sort":      "asc",
		"pageIndex": 1,
		"pageSize":  10,
	})
	log.Println(res, err)
}

func TestOpenGoodsQuery(t *testing.T) {
	//单品查询
	res, err := app.JdUnionOpenGoodsQuery(map[string]interface{}{
		"skuIds":   []int{30881878056},
		"isCoupon": "1",
	})
	log.Println(res, err)

	//列表查询
	res, err = app.JdUnionOpenGoodsQuery(map[string]interface{}{
		"sort":                 "asc",   // asc desc
		"sortName":             "price", //price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30Days：30天引单量， inOrderComm30Days：30天支出佣金
		"isCoupon":             1,
		"commissionShareStart": 20, //佣金比例开始区间
		"pageIndex":            1,
		"pageSize":             10,
		"cid1":                 1315,
	})
	log.Println(res, err)
}

// 获取单品信息
func TestOpenGoodsPromotionGoodsInfoQuery(t *testing.T) {
	res, err := app.JdUnionOpenGoodsPromotiongoodsinfoQuery(map[string]interface{}{
		"skuIds": "30881878056",
	})
	log.Println(res, err)
}

// 获取通用推广链接
func TestOpenPromotionCommonGet(t *testing.T) {
	res, err := app.JdUnionOpenPromotionCommonGet(map[string]interface{}{
		"subUnionId": "xc618",
		"ext1":       "100_618_618",
		"siteId":     app.ID,
		"materialId": "https://daojia.jd.com/activity/union/middlePage/index.html?channel=wm38094",
		"positionId": positionId,
		"command":    1,
	})
	log.Println(res, err)
}

// 获取商品订单
func TestOpenOrderQuery(t *testing.T) {
	//单品查询
	res, err := app.JdUnionOpenOrderQuery(map[string]interface{}{
		"type":     "1", //1 下单时间  2 完成时间 3 更新时间
		"time":     "201906141811",
		"pageNo":   1,
		"pagesize": 500,
	})
	log.Println(res, err)
}

func TestJdUnionOpenOrderRowQuery(t *testing.T) {
	// 订单行查询
	res, err := app.JdUnionOpenOrderRowQuery(map[string]interface{}{
		"type":      3, // 订单时间查询类型(1：下单时间，2：完成时间（购买用户确认收货时间），3：更新时间
		"startTime": "2025-05-24 22:00:00",
		"endTime":   "2025-05-24 22:59:00",
		"pageIndex": 1,
		"pageSize":  200,
		//"orderId": orderId,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res, res.Data)
	if len(res.Data) > 0 {
		log.Printf("%+v\n", res.Data[0])
	}
}

// 通过subUnionid获取推广链接
// https://wqitem.jd.com/item/view?sku=
func TestOpenPromotionBySubUnionIdGet(t *testing.T) {
	res, err := app.JdUnionOpenPromotionBysubunionidGet(map[string]interface{}{
		// 子渠道标识，仅支持传入字母、数字、下划线或中划线，最多80个字符（不可包含空格），该参数会在订单行查询接口中展示
		"subUnionId": "xc618",
		"positionId": positionId, // 推广位ID
		"chainType":  2,          // 转链类型（必填），1：长链， 2 ：短链 ，3： 长链+短链，默认短链，短链有效期60天
		// 推广物料url
		"materialId": "https://hour.jd.com/activity/hourbuyh5/pages/channelPage/index?rn_args-channelId=13355&cpsAutomaticCouponPagesource=71&flag=cps&eParam=618%7Cpc%7C&cu=true&utm_source=lianmeng__2__kong&utm_medium=tuiguang&utm_campaign=t_288552067_&utm_term=4f74f71977cf4614b91353aacbfcd7a5&hasLogin=1",
	})
	log.Println(res, err)
}

func TestJdUnionOpenStatisticsPromotionQuery(t *testing.T) {
	res, err := app.JdUnionOpenStatisticsPromotionQuery(map[string]interface{}{
		"activityUrl": "https://u.jd.com/r6PEcbm",
		"fields":      "clickPv,estimateValidOrders",
	})
	log.Println(res, err)
}
