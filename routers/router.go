package routers

import (
	"github.com/astaxie/beego"
	"CmsProject/controllers"
)

/**
 * 路由
 */
func init() {

	beego.Router("/", &controllers.MainController{})

	//============================管理员相关操作==========================
	beego.Router("/admin/login", &controllers.AdminController{}, "POST:AdminLogin")             //登陆
	beego.Router("/admin/info", &controllers.AdminController{}, "GET:GetAdminInfo")             //获取管理员信息
	beego.Router("/admin/singout", &controllers.AdminController{}, "GET:SignOut")               //管理员退出
	beego.Router("/admin/count", &controllers.AdminController{}, "GET:GetAdminCount")           //获取管理员总数
	beego.Router("/statis/admin/*/count", &controllers.AdminController{}, "GET:GetAdminStatis") ////获取某一日的管理员增长统计数据
	beego.Router("/admin/all", &controllers.AdminController{}, "GET:GetAdminList")              //查询整所有的用户列表

	//============================用户模块相关操作========================
	beego.Router("/statis/user/*/count", &controllers.UserController{}, "GET:UserStatisDaily") //获取某个日期的用户的增长数据
	beego.Router("/v1/users/count", &controllers.UserController{}, "GET:GetUserCount")
	beego.Router("/v1/users/list", &controllers.UserController{}, "GET:UserList")
	//订单模块使用
	beego.Router("/v1/user/:username", &controllers.UserController{}, "Get:GetUserInfoByUserName")

	//============================商家店铺相关操作========================
	beego.Router("/v1/cities", &controllers.CityController{}, "GET:GetCities")
	beego.Router("/shopping/restaurants/count", &controllers.ShopController{}, "GET:GetRestaurantCount") //获取商家总数
	beego.Router("/shopping/restaurants", &controllers.ShopController{}, "GET:GetRestaurantList")        //获取商户列表数据
	beego.Router("/shopping/restaurant/:id", &controllers.ShopController{}, "GET:GetRestaurantInfo;DELETE:DeleteRestaurant")
	beego.Router("/shopping/addShop", &controllers.ShopController{}, "POST:AddRestaurant") //添加商铺
	beego.Router("/v1/pois?:username", &controllers.ShopController{}, "GET:PoiSearch")

	//============================食品种类模块操作=========================
	beego.Router("/shopping/addcategory", &controllers.CategoryController{}, "POST:AddCategory")
	beego.Router("/shopping/getcategory/:id", &controllers.CategoryController{}, "GET:GetShopCategory")
	beego.Router("/shopping/v2/restaurant/category", &controllers.CategoryController{}, "GET:GetAllCategory")
	beego.Router("/shopping/addfood", &controllers.FoodController{}, "POST:AddFood")

	//============================食品模块相关操作=========================
	beego.Router("/shopping/v2/foods/count", &controllers.FoodController{}, "GET:GetFoodCount") //获取食品总数
	beego.Router("/shopping/v2/foods", &controllers.FoodController{}, "GET:GetFoodList")        //获取食品列表记录
	beego.Router("/shopping/v2/food/:id", &controllers.FoodController{}, "DELETE:DeleteFood")   //删除食品列表

	//============================订单模块相关操作=========================
	beego.Router("/statis/order/*/count", &controllers.OrderController{}, "GET:GetOrderStatis")
	beego.Router("/bos/orders/count", &controllers.OrderController{}, "GET:GetOrderCount") //获取订单总数
	beego.Router("/bos/orders", &controllers.OrderController{}, "GET:GetOrderList")        //获取订单列表
	beego.Router("/v1/addresse/:id", &controllers.OrderController{}, "GET:GetOrderAddress")

	//==================文件相关操作==================//未测试
	//beego.Router("/v1/addimg/:username", &controllers.FileController{}, "POST:UploadImg")
	beego.Router("/admin/update/avatar/:adminId", &controllers.FileController{}, "POST:UpdateAdminAvatar")
	beego.Router("/v1/addimg/:username", &controllers.FileController{}, "POST:UploadImg")
}
