package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"CmsProject/util"
	"CmsProject/models"
)

type CategoryController struct {
	beego.Controller
}

//表名
const CATEGORYTABLENAME = "food_category"

/**
 * 添加食品种类记录
 */
func (this *CategoryController) AddCategory() {

	util.LogInfo("添加食品种类")

	resp := make(map[string]interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	var category models.FoodCategory
	json.Unmarshal(this.Ctx.Input.RequestBody, &category)

	om := orm.NewOrm()
	_, err := om.Insert(&category)
	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_CATEGORYADD)
	} else {
		resp["status"] = util.RECODE_OK
		resp["success"] = util.Recode2Text(util.RESPMSG_SUCCESS_CATEGORYADD)
	}
}

/**
 * 查询某个食品店铺所拥有的食品种类
 */

func (this *CategoryController) GetShopCategory() {

	util.LogInfo("获取食品店铺拥有的食品种类")

	resp := make(map[string]interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	shopId := this.Ctx.Input.Param(":id")
	var categories []*models.FoodCategory
	om := orm.NewOrm()
	_, err := om.QueryTable(CATEGORYTABLENAME).Filter("restaurant_id", shopId).All(&categories)
	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_CATEGORIES)
	} else {
		resp["status"] = util.RECODE_OK
		resp["category_list"] = &categories
	}
}

/**
 * 获取到所有的食品种类
 */
func (this *CategoryController) GetAllCategory() {

	util.LogInfo("获取所有的食品种类")

	resp := make(map[string]interface{})

	defer this.ServeJSON()

	//判断用户是否已经登录 未登录，返回没有权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"] = util.EEROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.EEROR_UNLOGIN)
		return
	}

	var categories []*models.FoodCategory
	om := orm.NewOrm()
	_, err := om.QueryTable(CATEGORYTABLENAME).Filter("parent_category_id", 0).All(&categories)
	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_CATEGORIES)
		this.Data["json"] = resp
	} else {
		var respList []interface{}
		for _, category := range categories {
			respList = append(respList, category.CategoryToStringResp())
		}
		this.Data["json"] = respList
	}
}

/**
 * 判断用户是否已经登录
 */
func (this *CategoryController) IsLogin() bool {
	adminByte := this.GetSession(ADMIN)
	if adminByte != nil {
		var admin models.Admin
		json.Unmarshal(adminByte.([]byte), &admin)
		return admin.Id > 0
	}
	return false
}
