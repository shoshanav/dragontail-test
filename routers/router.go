package routers

import (
	"github.com/dragontail/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/restaurant",
			beego.NSInclude(
				&controllers.RestaurantController{},
		),
	)
	beego.AddNamespace(ns)
	beego.SetStaticPath("/assets", "assets")
	beego.SetStaticPath("/public/assets", "public/assets")
}