package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

/**
 * 默认页面
 */
func (c *MainController) Get() {
	c.TplName = "index.html"
}
