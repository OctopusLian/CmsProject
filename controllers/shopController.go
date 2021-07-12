package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"CmsProject/util"
	"CmsProject/models"
	"strconv"
	"encoding/json"
	"math/rand"
	"net/http"
	"io/ioutil"
	"CmsProject/entity"
)

/**
 * 商铺控制器
 */
type ShopController struct {
	beego.Controller
}

const (
	SHOPTABLENAME = "shop"
)

/**
 * 根据关键词搜索地图地址信息
 */
func (this *ShopController) PoiSearch() {

	util.LogInfo("搜索地图地址信息")

	resp := make(map[string]interface{})
	defer this.ServeJSON()

	rs, err := http.Get("https://elm.cangdu.org" + this.Ctx.Request.URL.String())

	//重要
	defer rs.Body.Close()
	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_SEARCHADDRESS
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_SEARCHADDRESS)
		this.Data["json"] = resp
		return
	}

	//读取请求数据并封装返回到前端
	body, err := ioutil.ReadAll(rs.Body)
	var searchList []*entity.PoiSearch
	json.Unmarshal(body, &searchList)
	this.Data["json"] = searchList
}

/**
 * 获取商家店铺总数
 */
func (this *ShopController) GetRestaurantCount() {

	util.LogInfo("获取商家店铺总数")

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
	restaurantCount, err := om.QueryTable(SHOPTABLENAME).Filter("dele", 0).Count()

	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["count"] = 0
	} else {
		resp["status"] = util.RECODE_OK
		resp["count"] = restaurantCount
	}
}

/**
 * 获取商家店铺列表
 */
func (this *ShopController) GetRestaurantList() {

	util.LogInfo("获取商家店铺列表")

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

	var sellerList []*models.Shop
	offset, _ := this.GetInt("offset") //偏移量
	limit, _ := this.GetInt("limit")   //本次查询所需要的记录条数

	om := orm.NewOrm()
	om.QueryTable(SHOPTABLENAME).Filter("dele", 0).Limit(limit, offset).All(&sellerList)

	var respList []interface{}
	for _, shop := range sellerList {
		respList = append(respList, shop.ShopToRespDesc())
	}

	if len(respList) > 0 {
		this.Data["json"] = &respList
	} else {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_RESTLIST
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_RESTLIST)
	}
}

/**
 * 获取某个店铺的信息
 */
func (this *ShopController) GetRestaurantInfo() {

	util.LogInfo("获取某个商家店铺的信息")

	resp := make(map[string]interface{})
	defer this.ServeJSON()

	//判断用户是否已经登录 未登录，返回没有权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"] = util.EEROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.EEROR_UNLOGIN)
		return
	}

	//获取要查询的商家的id
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	var restaurant models.Shop
	om := orm.NewOrm()
	om.QueryTable(SHOPTABLENAME).Filter("id", id).Filter("dele", 0).One(&restaurant)
	if (restaurant.Id > 0) {
		this.Data["json"] = restaurant
	} else {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_RESTAURANTINFO
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_RESTAURANTINFO)
		this.Data["json"] = resp
	}
}

/**
 * 删除商家店铺信息
 */
func (this *ShopController) DeleteRestaurant() {
	util.LogInfo("删除商家店铺信息")
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

	//查询当前管理员的权限
	adminByte := this.GetSession(ADMIN)
	var admin models.Admin
	json.Unmarshal(adminByte.([]byte), &admin)
	beego.Warn("管理员权限:", len(admin.Permission))

	om := orm.NewOrm()
	om.LoadRelated(&admin, "Permission") //Permission 是Admin实体中的字段名，而非表名
	beego.Warn("关联查询后的管理员权限：", len(admin.Permission))

	var authority bool
	var permissions []*models.Permission
	permissions = admin.Permission
	for _, permission := range permissions {
		if permission.Level == "DELETE" {
			authority = true
			break
		}
	}

	//现判断当前管理员的权限，如果权限不够，则直接返回没有删除权限
	if !authority {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_HASNOACCESS
		resp["message"] = util.Recode2Text(util.RESPMSG_HASNOACCESS)
		return
	}

	//删除.修改一个字段，dele = 0 变为dele = 1
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	seller := models.Shop{Id: id, Dele: 1}
	_, err := om.Update(&seller, "dele")

	if err != nil {
		//删除失败
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_HASNOACCESS
		resp["message"] = util.Recode2Text(util.RESPMSG_HASNOACCESS)
	} else {
		resp["status"] = util.RECODE_OK
		resp["type"] = util.RESPMSG_SUCCESS_DELETESHOP
		resp["message"] = util.Recode2Text(util.RESPMSG_SUCCESS_DELETESHOP)
	}
}

/**
 * 添加商家店铺信息
 */
func (this *ShopController) AddRestaurant() {

	util.LogInfo("添加商家店铺信息")

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

	//获取客户端传值方法：
	//1.this.GetString : get,post等请求直接传递参数
	//2.this.Ctx.Input.Param(key) : 路由中设置的正则表达式或者正则变量
	//3.this.Ctx.Input.RequstBody : 获取到的是字节数组数据, 适合json类型数据格式传输请求  要求：必须要在app.conf中设置，否则无法接受

	var restaurantEntity models.Shop
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &restaurantEntity)
	restaurantEntity.Status = 1                      //店铺状态为正常
	restaurantEntity.Rating = rand.Intn(10)          //评分
	restaurantEntity.RatingCount = rand.Intn(1000)   //评分总次数
	restaurantEntity.RecentOrderNum = rand.Intn(500) //最近的订单数量

	//重要 //

	om := orm.NewOrm()

	om.Begin() //事务开始

	//添加记录的方法
	id, err := om.Insert(&restaurantEntity)

	//添加店铺所支持的活动
	activities := restaurantEntity.Activities
	//多对多的高级查询
	//第一个参数：对象，主键必须有值
	//第二个参数：对象需要操作的 M2M 字段
	m2m := om.QueryM2M(&restaurantEntity, "activities")

	for _, activity := range activities {
		om.Insert(activity)
		m2m.Add(activity)
	}
	if err == nil {
		om.Commit() //事务提交
	}

	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["message"] = util.Recode2Text(util.RESPMSG_FAIL_ADDREST)
	} else {
		var restuarant models.Shop
		om.QueryTable(SHOPTABLENAME).Filter("id", id).One(&restuarant)
		if restuarant.Id > 0 {
			resp["status"] = util.RECODE_OK
			resp["sussess"] = util.Recode2Text(util.RESPMSG_SUCCESS_ADDREST)
			resp["shopDetail"] = restuarant
		} else {
			resp["status"] = util.RECODE_FAIL
			resp["sussess"] = util.Recode2Text(util.RESPMSG_FAIL_ADDREST)
		}
	}
}

/**
 * 判断用户是否已经登录
 */
func (this *ShopController) IsLogin() bool {
	adminByte := this.GetSession(ADMIN)
	if adminByte != nil {
		var admin models.Admin
		json.Unmarshal(adminByte.([]byte), &admin)
		return admin.Id > 0
	}
	return false
}
