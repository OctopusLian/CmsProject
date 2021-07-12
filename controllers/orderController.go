package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"CmsProject/util"
	"encoding/json"
	"CmsProject/models"
	"math/rand"
	"CmsProject/entity"
	"strings"
	"time"
)

/**
 * 订单控制器
 */
type OrderController struct {
	beego.Controller
}

const (
	ORDERTABLENAME   = "user_order"
	ADDRESSTABLENAME = "address"
)

/**
 * 获取某一日的订单的订单增长数量
 */
func (this *OrderController) GetOrderStatis() {

	util.LogInfo("获取某一日的订单的订单增长数量")

	reJson := make(map[string]interface{})
	this.Data["json"] = reJson
	defer this.ServeJSON()

	//判断权限
	if !this.IsLogin() {
		reJson["status"] = util.RECODE_UNLOGIN
		reJson["type"] = util.EEROR_UNLOGIN
		reJson["message"] = util.Recode2Text(util.EEROR_UNLOGIN)
		return
	}

	/**
	 * 2018-11-29日新增逻辑：
	 *   先从redis缓存中获取日增长量的数据，能够获取到数据，直接返回
	 *   在redis中获取不到数据，再利用orm操作查询数据库，然后存放在redis中，同时返回数据
	 */
	//先获取redis缓存实例
	redisConn, err := util.GetRedis()
	if err != nil {
		reJson["status"] = util.RECODE_FAIL
		reJson["count"] = 0
		return
	}

	//获取参数
	paths := strings.Split(this.Ctx.Input.URL(), "/")

	//从redis中根据key值获取对应的数据缓存
	statis := redisConn.Get("orderStatis" + paths[3])

	if statis != nil {
		var statisCount entity.StatisEntity
		json.Unmarshal(statis.([]byte), &statisCount)
		reJson["status"] = util.RECODE_OK
		reJson["count"] = statisCount.StatisCount
		return
	}

	om := orm.NewOrm()

	//得到要进行统计的日期
	//datetime := paths[3]

	//采用字符串来进行比较查询某一日的订单
	// select * from user_order where time like '2018-11-28%'
	//select count(*) from user_order ...  Count()方法对应的sql语句
	//orderCount, err := om.QueryTable(ORDERTABLENAME).Filter("time__istartswith", datetime).Count()
	//beego.Info("订单日增长数量：", orderCount)

	// 采用秒数来进行比较计算查询某日的订单 gt 大于 lt 小于
	//timeLayout := "2006-01-02"           //转化所需模板
	//loc, _ := time.LoadLocation("Local") //获取时区
	//tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	//beego.Info(tmp.Unix(), tmp.UnixNano())
	//start := tmp.Unix() * 1000 //转化为时间戳 类型是int64

	//一天是 1*60 * 60 * 24 = 86400秒
	//end := (start + 86400*1000)

	//select * from user_order where order_time > ? and order_time < ? values ( start, end )
	//select count(*) from ....
	//orderCount, err := om.QueryTable(ORDERTABLENAME).
	//	Filter("order_time__gt", start).
	//	Filter("order_time__lt", end).
	//	Count()
	//beego.Info("查询结果：", orderCount)

	//仅做测试效果使用，正确的代码应该是上面的注释代码
	orderCount, err := om.QueryTable(ORDERTABLENAME).Count()
	if err != nil {
		beego.Info(orderCount)
		reJson["status"] = util.RECODE_FAIL
		reJson["count"] = 0
		return
	}

	/**
	 * 2018-11-29新增逻辑:
	 *  因为每日增长数据这个数据要求优先级并不是特别高，因此把查询到的数据存放于redis中，每隔1分钟更新一次
	 */

	//分两种情况进行redis缓存
	// statis/order/NaN-NaN-NaN/count
	// statis/order/2018-11-23/count
	// statis/order/2018-11-24/count
	// statis/order/2018-11-25/count
	// statis/order/2018-11-26/count
	// statis/order/2018-11-27/count
	// statis/order/2018-11-28/count

	// statis/admin/2018-11-29/count
	//日期必须是2006-01-02
	todayStr := time.Now().Format("2006-01-02")
	statisCount := &entity.StatisEntity{
		StatisDate:  "orderStatis" + paths[3],
		StatisCount: rand.Intn(150),
	}
	bytes, _ := json.Marshal(statisCount)

	if todayStr == paths[3] {
		redisConn.Put("orderStatis"+paths[3], bytes, 60*time.Second)
	} else {
		redisConn.Put("orderStatis"+paths[3], bytes, 60*60*24*time.Second)
	}

	reJson["status"] = util.RECODE_OK
	//仅做测试效果使用，正确的代码应该是下的注释代码
	reJson["count"] = statisCount.StatisCount
	//reJson["count"] = orderCount

}

/**
 * 获取某个商家的用户订单的总数量
 */
func (this *OrderController) GetOrderCount() {

	util.LogInfo("获取用户订单总数量")

	resp := make(map[string]interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	//判断用户是否已经登录 未登录，返回没有权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"] = util.EEROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.EEROR_UNLOGIN)
		return
	}

	om := orm.NewOrm()
	//0 代表正常 1代表已删除
	count, err := om.QueryTable(ORDERTABLENAME).Filter("del_flag", 0).Count()

	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_ORDERCOUNT
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_ORDERCOUNT)
		return
	}

	resp["status"] = util.RECODE_OK
	resp["count"] = count
}

/**
 * 获取用户订单列表
 */
func (this *OrderController) GetOrderList() {

	util.LogInfo("获取用户订单列表")

	resp := make(map[string]interface{})
	defer this.ServeJSON()

	//判断用户是否已经登录 未登录，返回没有权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"] = util.EEROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.EEROR_UNLOGIN)
		return
	}

	var orderList []*models.UserOrder
	offset, _ := this.GetInt("offset")
	limit, _ := this.GetInt("limit")

	om := orm.NewOrm()
	num, err := om.QueryTable(ORDERTABLENAME).Limit(limit, offset).All(&orderList)

	if num <= 0 || err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_ORDERLIST
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_ORDERLIST)
		this.Data["json"] = resp
		return
	}

	var respList []interface{}
	for _, order := range orderList {
		om.LoadRelated(order, "Shop")        //关联查询  Shop为 Order结构体中声明的字段名 而非 表名 ！切记！
		om.LoadRelated(order, "User")        //关联查询  User为 Order结构体中声明的字段名 而非 表名 ！切结！
		om.LoadRelated(order, "OrderStatus") //关联查询  OrderStatus 为Order结构体中声明的字段名 而非 表名 ！切记！
		respList = append(respList, order.UserOrder2Resp())
	}
	this.Data["json"] = respList
}

/**
 * 获取订单地址
 */
func (this *OrderController) GetOrderAddress() {

	util.LogInfo("获取订单地址")

	resp := make(map[string]interface{})
	defer this.ServeJSON()

	addressId := this.Ctx.Input.Param(":id")
	var address models.Address

	om := orm.NewOrm()
	om.QueryTable(ADDRESSTABLENAME).Filter("id", addressId).One(&address)

	if (address.Id <= 0) {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_ORDERINFO
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_ORDERINFO)
		this.Data["json"] = resp
		return
	}
	this.Data["json"] = &address
}

/**
 * 判断用户是否已经登录的方法
 */
func (this *OrderController) IsLogin() bool {
	adminByte := this.GetSession(ADMIN)
	if adminByte != nil {
		var admin models.Admin
		json.Unmarshal(adminByte.([]byte), &admin)
		return admin.Id > 0
	}
	return false
}
