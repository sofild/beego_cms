package main

import (
	_ "beego_cms/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
