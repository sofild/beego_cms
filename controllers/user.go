package controllers

import (
	"beego_cms/models"
	"fmt"
	"github.com/astaxie/beego"
	//"strconv"
	"time"
)

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {
	c.TplName = "login.tpl"
}

func (c *UserController) DoLogin() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		beego.Info(err)
	} else {
		uinfo := models.FindUser(u.Username, u.Password)
		fmt.Println(uinfo)
		if uinfo.Id > 0 {
			uinfo.Logintime = time.Now().Unix()
			models.UpUser(uinfo)

			c.SetSession("uid", uinfo.Id)
			c.SetSession("username", uinfo.Username)

			c.Redirect("/", 302)
			//c.Ctx.WriteString("UID:" + strconv.FormatInt(uinfo.Id, 10))
		} else {
			c.Ctx.WriteString("Login Failed.")
		}
	}
}

func (c *UserController) Reg() {

}

func (c *UserController) DoReg() {

}

func (c *UserController) Logout() {
	c.DelSession("uid")
	c.Redirect("/user/login", 302)
}
