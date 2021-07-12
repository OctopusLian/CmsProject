package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"CmsProject/util"
	"CmsProject/models"
	"encoding/json"
	"time"
)

type CityController struct {
	beego.Controller
}

const (
	CITYTABLENAME = "city"
	CITYREDIS     = "cities"
)

/**
 * 查询并返回城市数据列表
 */
func (this *CityController) GetCities() {

	util.LogInfo("获取城市信息")

	resp := make(map[string]interface{})
	defer this.ServeJSON()

	//判断用户是否已经登录 未登录，返回没有权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"] = util.EEROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.EEROR_UNLOGIN)
		return
	}

	//获取redis实例
	redisConn, _ := util.GetRedis()
	//逻辑严谨
	var citiesByte interface{}
	if redisConn != nil {
		citiesByte = redisConn.Get(CITYREDIS)
	}

	var cities []*models.City
	if citiesByte != nil { //缓存中存在
		json.Unmarshal(citiesByte.([]byte), &cities)
	} else { //缓存不存在，查询，并存放到缓存中
		//1.数据库查询
		om := orm.NewOrm()
		om.QueryTable(CITYTABLENAME).All(&cities)
		//2.存放于redis当中
		citiesByte, _ = json.Marshal(&cities)
		redisConn.Put(CITYREDIS, citiesByte, 60*60*24*time.Second)
	}

	var cityList []interface{}
	for _, city := range cities {
		cityList = append(cityList, city.CityToRespDesc())
	}

	if len(cityList) > 0 {
		this.Data["json"] = cityList
	} else {
		resp["status"] = util.RESPMSG_FAIL
		resp["type"] = util.RESPMSG_ERROR_CITYLIST
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_CITYLIST)
		this.Data["json"] = resp
	}
}

/**
 * 判断用户是否已经登录
 */
func (this *CityController) IsLogin() bool {
	adminByte := this.GetSession(ADMIN)
	if adminByte != nil {
		var admin models.Admin
		json.Unmarshal(adminByte.([]byte), &admin)
		return admin.Id > 0
	}
	return false
}
