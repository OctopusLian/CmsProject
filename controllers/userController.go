package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"CmsProject/util"
	"encoding/json"
	"CmsProject/models"
	"math/rand"
	"strings"
	"CmsProject/entity"
	"time"
)

type UserController struct {
	beego.Controller
}

const (
	USERTABLENAME = "user"
)

/**
 * 获取某日新增用户
 */
func (this *UserController) UserStatisDaily() {

	util.LogInfo("获取用户某日增长量统计")

	resp := make(map[string]interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	//判断是否已经登录，是否有查询权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"] = util.EEROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.EEROR_UNLOGIN)
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
		resp["status"] = util.RECODE_FAIL
		resp["count"] = 0
		return
	}

	//获取参数
	paths := strings.Split(this.Ctx.Input.URL(), "/")

	//从redis中根据key值获取对应的数据缓存
	statis := redisConn.Get("userStatis" + paths[3])

	if statis != nil {
		var statisCount entity.StatisEntity
		json.Unmarshal(statis.([]byte), &statisCount)
		resp["status"] = util.RECODE_OK
		resp["count"] = statisCount.StatisCount
		return
	}

	om := orm.NewOrm()

	//获取参数
	//paths := strings.Split(this.Ctx.Input.URL(), "/")
	//得到要进行统计的日期
	//datetime := paths[3]
	//采用字符串来进行比较查询某一日的用户增长量
	// select * from user where registe_time like '2018-11-28%'
	//select count(*) from user_order ...  Count()方法对应的sql语句
	//userCount, err := om.QueryTable(USERTABLENAME).Filter("registe_time__istartswith", datetime).Count()

	//仅做测试效果使用，正确的代码应该是上面的注释代码
	userCount, err := om.QueryTable(USERTABLENAME).Count()
	if err != nil {
		beego.Info(userCount)
		resp["status"] = util.RECODE_FAIL
		resp["count"] = 0
		return
	}

	/**
	 * 2018-11-29新增逻辑:
	 *  因为每日增长数据这个数据要求优先级并不是特别高，因此把查询到的数据存放于redis中，每隔1分钟更新一次
	 */

	//分两种情况进行redis缓存
	// statis/user/NaN-NaN-NaN/count
	// statis/user/2018-11-23/count
	// statis/user/2018-11-24/count
	// statis/user/2018-11-25/count
	// statis/user/2018-11-26/count
	// statis/user/2018-11-27/count
	// statis/user/2018-11-28/count

	// statis/user/2018-11-29/count

	//日期必须是2006-01-02
	todayStr := time.Now().Format("2006-01-02")
	statisCount := &entity.StatisEntity{
		StatisDate:  "userStatis" + paths[3],
		StatisCount: rand.Intn(150),
	}
	bytes, _ := json.Marshal(statisCount)

	if todayStr == paths[3] {
		redisConn.Put("userStatis"+paths[3], bytes, 60*time.Second)
	} else {
		redisConn.Put("userStatis"+paths[3], bytes, 60*60*24*time.Second)
	}

	resp["status"] = util.RECODE_OK
	//仅做测试效果使用，正确的代码应该是下面的注释代码
	resp["count"] = statisCount.StatisCount
	//resp["count"] = userCount
}

/**
 * 获取总用户数
 */
func (this *UserController) GetUserCount() {

	util.LogInfo("获取用户总数")

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
	userCount, err := om.QueryTable(USERTABLENAME).Count()
	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["count"] = 0
	} else {
		resp["status"] = util.RECODE_OK
		resp["count"] = userCount
	}
}

//获取用户信息列表
func (this *UserController) UserList() {

	util.LogInfo("获取用户列表")

	resp := make(map[string]interface{})
	this.Data["jsoon"] = resp
	defer this.ServeJSON()

	//判断用户是否已经登录 未登录，返回没有权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"] = util.EEROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.EEROR_UNLOGIN)
		return
	}

	//2.查询数据并放入[]中
	var userList []*models.User
	offset, _ := this.GetInt("offset")
	limit, _ := this.GetInt("limit")

	om := orm.NewOrm()
	om.QueryTable(USERTABLENAME).Filter("del_flag", 0).Limit(limit, offset).All(&userList)

	//3.使用loadRelated方法进行关联查询，并进行json格式组装
	var respList []interface{}
	for _, user := range userList {
		om.LoadRelated(user, "City")
		respList = append(respList, user.UserToRespDesc())
	}

	//4.返回查询数据
	if len(userList) > 0 { //查询到了用户数据
		this.Data["json"] = &respList
	} else { //未查询到用户数据
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_USERLIST
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_USERLIST)
	}
}

/**
 * 通过用户名查询用户信息
 */
func (this *UserController) GetUserInfoByUserName() {

	util.LogInfo("通过用户名查询用户信息")

	resp := make(map[string]interface{})

	defer this.ServeJSON()

	//判断用户是否已经登录 未登录，返回没有权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"] = util.EEROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.EEROR_UNLOGIN)
		return
	}

	om := orm.NewOrm()
	var user models.User
	om.QueryTable(USERTABLENAME).Filter("user_name", this.Ctx.Input.Param(":username")).One(&user)

	if user.Id <= 0 {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_USERINFO
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_USERINFO)
		this.Data["json"] = resp
		return
	}
	this.Data["json"] = user.UserToRespDesc()
}

/**
 * 判断用户是否已经登录
 */
func (this *UserController) IsLogin() bool {
	adminByte := this.GetSession(ADMIN)
	if adminByte != nil {
		var admin models.Admin
		json.Unmarshal(adminByte.([]byte), &admin)
		return admin.Id > 0
	}
	return false
}
