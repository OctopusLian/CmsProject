package main

import (
	_ "CmsProject/routers"
	"github.com/astaxie/beego"
	//连接数据操作测试
	//_ "CmsProject/models"
)

/**
 * 前端资源导入：
 * 1.index.html作为项目默认的首页导入到views目录下
 * 2.js,css,fonts,img等静态资源导入到static目录下
 * 3.修改index.html中css,js文件的引用
 * 4.浏览器开发者工具调试，设置静态资源路径
 */

func main() {

	beego.AddAPPStartHook(func() error {
		beego.Info(" 自定义配置 ")
		return nil
	})

	//映射前：http://localhost:8080/manage/static/js/16.7b80c57163637f4aa1ae.js
	//映射后：http://localhost:8080/static/js16.7b80c57163637f4aa1ae.js
	beego.SetStaticPath("manage/static", "static")

	//上传图片的映射
	beego.SetStaticPath("img", "img")

	//监听
	beego.Run()
}

//示例

//1对1的关系
//table person
type Person struct {
	Id      int
	Name    string       `json:"name"`                    //姓名
	Age     int          `json:"age"`                     //年龄
	Address *HomeAddress `json:"address"  orm:"rel(one)"` //家庭地址 //1对1的关系映射设置
}

//table homeaddress
type HomeAddress struct {
	Id            int
	AddressName   string  `json:"address_name"`               //地址名称
	AddressDetail string  `json:"address_detail"`             //详细地址
	AddressPort   int     `json:"address_port"`               //门牌号
	Person        *Person `json:"person"  orm:"reverse(one)"` //地址所属的主人 //1对1的关系反向映射
}
