package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"CmsProject/util"
	"CmsProject/models"
	"strconv"
	"encoding/json"
	"math/rand"
	"CmsProject/entity"
)

/**
 * 食品控制器
 */
type FoodController struct {
	beego.Controller
}

const (
	FOODTABLENAME = "food"
)

/**
 * 获取食品总数
 */
func (this *FoodController) GetFoodCount() {

	util.LogInfo("获取食品总数")

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
	foodCount, err := om.QueryTable(FOODTABLENAME).Filter("del_flag", 0).Count()
	if err != nil {
		resp["status"] = util.RESPMSG_FAIL
		resp["count"] = 0
	} else {
		resp["status"] = util.RESPMSG_OK
		resp["count"] = foodCount
	}
}

/**
 *  获取食品列表
 */
func (this *FoodController) GetFoodList() {

	util.LogInfo("获取食品列表")

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

	var foods []*models.Food
	om := orm.NewOrm()
	offset, _ := this.GetInt("offset")
	limit, _ := this.GetInt("limit")

	foodCount, err := om.QueryTable(FOODTABLENAME).Filter("del_flag", 0).Limit(limit, offset).All(&foods)
	//关联查询商铺的信息，食品种类信息
	var foodList []interface{}
	for _, food := range foods {
		//这里的关联第二个字段 是 字段名字 而非 关联的表名！！切记！！
		om.LoadRelated(food, "Restaurant") //关联商铺表
		om.LoadRelated(food, "Category")   //关联食品种类表
		foodList = append(foodList, food.FoodToRespDesc())
	}
	if foodCount <= 0 || err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_FOODLIST
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_FOODLIST)
	} else {
		this.Data["json"] = &foodList
	}
}

/**
 * 删除食品记录
 */
func (this *FoodController) DeleteFood() {

	util.LogInfo("删除食品")

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
	om.LoadRelated(&admin, "Permission")
	beego.Warn("关联查询后的管理员权限：", len(admin.Permission))

	var authority bool
	var permissions []*models.Permission
	permissions = admin.Permission
	for _, permission := range permissions {
		if permission.Level == "DELETE" {
			authority = true
		}
	}

	//判断是否有权限
	if !authority { //没有权限
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_HASNOACCESS
		resp["message"] = util.Recode2Text(util.RESPMSG_HASNOACCESS)
		return
	}

	//获取要删除的记录id
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	//更新数据库字段
	food := models.Food{Id: id, DelFlag: 1}
	_, err := om.Update(&food, "del_flag")
	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_FOODDELE
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_FOODDELE)
		return
	}

	//删除成功再把删除后的数据返回给前台
	var foodEntity models.Food
	om.QueryTable(FOODTABLENAME).Filter("id", id).Filter("del_flag", 1).One(&foodEntity)
	if foodEntity.Id > 0 {
		resp["status"] = util.RECODE_OK
	} else {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_FOODDELE
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_FOODDELE)
	}
}

/**
 * 添加食品记录
 */
func (this *FoodController) AddFood() {
	util.LogInfo("添加食品")

	resp := make(map[string]interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	var addFood entity.AddFoodEntity
	//json解析错误：json: cannot unmarshal string into Go struct field AddFoodEntity.restaurant_id of type int
	//json解析错误：  无法解析 string 类型 到 结构体类型AddFoodEnitity 中的 int 类型的变量 restaurant_id
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &addFood)
	if err != nil {
		beego.Info(err.Error())
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_FOODADD
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_FOODADD)
		return
	}

	categoryId := addFood.CategoryId
	restaurantId := addFood.RestaurantId

	var category models.FoodCategory
	var restaurant models.Shop

	om := orm.NewOrm()

	//查询category记录
	om.QueryTable("food_category").Filter("id", categoryId).One(&category)
	//查询restaurant记录
	om.QueryTable(SHOPTABLENAME).Filter("id", restaurantId).One(&restaurant)

	var food models.Food
	food.Name = addFood.Name
	food.Description = addFood.Description
	food.ImagePath = addFood.ImagePath
	food.Activity = addFood.Activity
	food.Category = &category
	food.Restaurant = &restaurant
	food.DelFlag = 0
	food.Rating = rand.Intn(10)

	num, err := om.Insert(&food)
	beego.Info(num)
	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_FOODADD
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_FOODADD)
	} else {
		resp["status"] = util.RECODE_OK
		resp["success"] = util.Recode2Text(util.RESPMSG_SUCCESS_FOODADD)
	}
}

/**
 * 判断用户是否已经登录
 */
func (this *FoodController) IsLogin() bool {
	adminByte := this.GetSession(ADMIN)
	if adminByte != nil {
		var admin models.Admin
		json.Unmarshal(adminByte.([]byte), &admin)
		return admin.Id > 0
	}
	return false
}
