package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"CmsProject/models"
	"CmsProject/entity"
	"CmsProject/util"
	"strings"
	"math/rand"
	"time"
)

/**
 * 管理员控制器
 */
type AdminController struct {
	beego.Controller
}

const (
	ADMINTABLENAME = "admin"
	ADMIN          = "admin"
)

/**
 * 管理员登陆
 */
func (this *AdminController) AdminLogin() {

	util.LogInfo("管理员登陆")

	reJson := make(map[string]interface{})
	this.Data["json"] = reJson
	defer this.ServeJSON()

	//1.获取请求数据
	var loginEntity entity.AdminLoginEntity
	util.JsonToEntity(this.Ctx.Input.RequestBody, &loginEntity)

	//查询结果变量
	var admin models.Admin

	//示例化orm对象
	om := orm.NewOrm()

	//select * from admin where user_anme = ? and pwd = ?  values root, 1234
	om.QueryTable(ADMINTABLENAME).Filter("user_name", loginEntity.User_name).Filter("pwd", loginEntity.Password).One(&admin)

	if (admin.Id > 0) { //管理员登录成功

		userByte, _ := json.Marshal(admin)

		//设置session
		this.SetSession(ADMIN, userByte)

		reJson["status"] = util.RECODE_OK
		reJson["success"] = util.Recode2Text(util.RESPMSG_SUCCESSLOGIN)
		return
	}

	reJson["status"] = util.RECODE_FAIL
	reJson["message"] = util.Recode2Text(util.RESPMSG_FAILURELOGIN)
}

/**
 *  获取管理员信息
 */
func (this *AdminController) GetAdminInfo() {

	util.LogInfo("获取管理员信息")

	reJson := make(map[string]interface{})
	this.Data["json"] = reJson
	defer this.ServeJSON()

	//如果是xml
	//reXml := make(map[string]interface{})
	//this.Data["xml"] = reXml
	//defer this.ServeXML()

	//从session中获取信息
	userByte := this.GetSession(ADMIN)

	//session为空
	if userByte == nil {
		reJson["status"] = util.RECODE_UNLOGIN
		reJson["type"] = util.EEROR_UNLOGIN
		reJson["message"] = util.Recode2Text(util.EEROR_UNLOGIN)
		return
	}

	var admin models.Admin
	//安马歇尔
	err := json.Unmarshal(userByte.([]byte), &admin)

	if err != nil {
		//失败
		util.LogInfo("获取管理员信息失败")
		reJson["status"] = util.RECODE_FAIL
		reJson["type"] = util.RESPMSG_ERRORSESSION
		reJson["message"] = util.Recode2Text(util.RESPMSG_ERRORSESSION)
		return
	}

	//成功
	if (admin.Id > 0) {
		util.LogInfo("获取管理员信息成功")
		reJson["status"] = util.RECODE_OK
		reJson["data"] = admin.AdminToRespDesc()
		return
	}
}

/**
 *  管理员退出当前账号
 */
func (this *AdminController) SignOut() {

	util.LogInfo("管理员退出当前账号")

	resp := make(map[string]interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	//删除session
	this.DelSession(ADMIN)

	resp["status"] = util.RECODE_OK
	resp["success"] = util.Recode2Text(util.RESPMSG_SIGNOUT)
}

/**
 * 获取管理员总数
 */
func (this *AdminController) GetAdminCount() {

	util.LogInfo("获取管理员总数")

	reJson := make(map[string]interface{})
	this.Data["json"] = reJson
	defer this.ServeJSON()

	//判断是否有权限的优化
	if !this.IsLogin() {
		reJson["status"] = util.RECODE_UNLOGIN
		reJson["type"] = util.EEROR_UNLOGIN
		reJson["message"] = util.Recode2Text(util.EEROR_UNLOGIN)
		return
	}

	om := orm.NewOrm()
	adminCount, err := om.QueryTable(ADMINTABLENAME).Filter("status", 0).Count()
	if err != nil {
		reJson["status"] = util.RECODE_FAIL
		reJson["message"] = util.Recode2Text(util.RESPMSG_ERRORADMINCOUNT)
		reJson["count"] = 0
	} else {
		reJson["status"] = util.RECODE_OK
		reJson["count"] = adminCount
	}
}

/**
 * 返回管理员当日统计结果
 */
func (this *AdminController) GetAdminStatis() {

	util.LogInfo("获取管理员某个日期统计结果")

	reJson := make(map[string]interface{})
	this.Data["json"] = reJson
	defer this.ServeJSON()

	//beego.Info("请求path:", this.Ctx.Request.URL.Path)
	//beego.Info(this.Ctx.Input.URL())

	//判断是否有权限的优化
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
	//    statis/admin/2018-11-23/count
	paths := strings.Split(this.Ctx.Input.URL(), "/")

	//从redis中根据key值获取对应的数据缓存
	statis := redisConn.Get("adminStatis" + paths[3])

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
	// select * from admin where create_time like '2018-11-22%'
	//select count(*) from user_order ...  Count()方法对应的sql语句
	//adminCount, err := om.QueryTable(ADMINTABLENAME).Filter("create_time__istartswith", datetime).Count()

	//仅做测试效果使用，正确的代码应该是上面的注释代码
	adminCount, err := om.QueryTable(ADMINTABLENAME).Count()
	if err != nil {
		beego.Info(adminCount)
		reJson["status"] = util.RECODE_FAIL
		reJson["count"] = 0
		return
	}

	/**
	 * 2018-11-29新增逻辑:
	 *  因为每日增长数据这个数据要求优先级并不是特别高，因此把查询到的数据存放于redis中，每隔1分钟更新一次
	 */

	//分两种情况进行redis缓存
	// statis/admin/NaN-NaN-NaN/count
	//24小时更新一次 时效性24小时
	// statis/admin/2018-11-23/count    11月23日增长25人
	// statis/admin/2018-11-24/count    11月24日增长56人
	// statis/admin/2018-11-25/count
	// statis/admin/2018-11-26/count
	// statis/admin/2018-11-27/count
	// statis/admin/2018-11-28/count

	//1分钟更新一次 时效性1分钟
	// statis/admin/2018-11-29/count    11月29日 当天

	//日期必须是2006-01-02
	todayStr := time.Now().Format("2006-01-02") //2018-11-29
	statisCount := &entity.StatisEntity{
		//  { adminStatis2018-11-23   25 }
		//  { adminStatis2018-11-24   56 }
		StatisDate:  "adminStatis" + paths[3],
		StatisCount: rand.Intn(50),
	}
	bytes, _ := json.Marshal(statisCount)

	//分类存储到redis当中   2018-11-23
	if todayStr == paths[3] { //2018-11-29 今天当天
		redisConn.Put("adminStatis"+paths[3], bytes, 60*time.Second)
	} else {
		redisConn.Put("adminStatis"+paths[3], bytes, 60*60*24*time.Second)
	}

	reJson["status"] = util.RECODE_OK
	//仅做测试效果使用，正确的代码应该是下面的注释代码
	reJson["count"] = statisCount.StatisCount
	//reJson["count"] = adminCount
}

/**
 * 获取管理员列表
 */
func (this *AdminController) GetAdminList() {

	util.LogInfo("管理员列表")

	reJson := make(map[string]interface{})
	this.Data["json"] = reJson
	defer this.ServeJSON()

	//判断是否登录的权限
	if !this.IsLogin() {
		reJson["status"] = util.RECODE_UNLOGIN
		reJson["type"] = util.EEROR_UNLOGIN
		reJson["message"] = util.Recode2Text(util.EEROR_UNLOGIN)
		return
	}

	var adminList []*models.Admin
	om := orm.NewOrm()
	offset, _ := this.GetInt("offset")
	limit, _ := this.GetInt("limit")
	_, err := om.QueryTable(ADMINTABLENAME).Filter("status", 0).Limit(limit, offset).All(&adminList)

	if err != nil {
		reJson["status"] = util.RECODE_FAIL
		reJson["type"] = util.RESPMSG_ERROR_FOODLIST
		reJson["message"] = util.Recode2Text(util.RESPMSG_ERROR_FOODLIST)
		return
	}

	var respList []interface{}
	for _, admin := range adminList {
		//重点
		om.LoadRelated(admin, "City")
		respList = append(respList, admin.AdminToRespDesc())
	}

	reJson["status"] = util.RECODE_OK
	reJson["data"] = respList

}

/**
 * 判断用户是否已经登陆过：通过session进行判断
 */
func (this *AdminController) IsLogin() bool {
	adminByte := this.GetSession(ADMIN)
	if adminByte != nil {
		var admin models.Admin
		json.Unmarshal(adminByte.([]byte), &admin)
		return admin.Id > 0
	}
	return false
}
