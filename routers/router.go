package routers

import (
	"beego_cms/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/user/login", &controllers.UserController{}, "get:Login")
	beego.Router("/user/dologin", &controllers.UserController{}, "post:DoLogin")
	beego.Router("/user/reg", &controllers.UserController{}, "get:Reg")
	beego.Router("/user/doreg", &controllers.UserController{}, "post:DoReg")
	beego.Router("/user/logout", &controllers.UserController{}, "get:Logout")

	beego.Router("/articles", &controllers.ArticlesController{}, "get:Index")
	beego.Router("/articles/add/?:id", &controllers.ArticlesController{}, "get:Add")
	beego.Router("/articles/doadd", &controllers.ArticlesController{}, "post:DoAdd")
	beego.Router("/articles/list", &controllers.ArticlesController{}, "get:List")
	beego.Router("/articles/del/:id", &controllers.ArticlesController{}, "get:Del")

	beego.Router("/cate", &controllers.CateController{}, "get:Index")
	beego.Router("/cate/add", &controllers.CateController{}, "get:Add")
	beego.Router("/cate/doadd", &controllers.CateController{}, "post:DoAdd")
	beego.Router("/cate/list", &controllers.CateController{}, "get:List")
}
