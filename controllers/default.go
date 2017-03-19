package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	uid := c.GetSession("uid")
	if uid == nil {
		c.Redirect("/user/login", 302)
	}
	username := c.GetSession("username")
	c.Data["Username"] = username
	c.TplName = "index.tpl"
}
