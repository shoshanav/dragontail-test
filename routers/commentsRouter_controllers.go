package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/shoshanav/dragontail-test/controllers:RestaurantController"] = append(beego.GlobalControllerRouter["github.com/shoshanav/dragontail-test/controllers:RestaurantController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/shoshanav/dragontail-test/controllers:RestaurantController"] = append(beego.GlobalControllerRouter["github.com/shoshanav/dragontail-test/controllers:RestaurantController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:restaurantId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/shoshanav/dragontail-test/controllers:RestaurantController"] = append(beego.GlobalControllerRouter["github.com/shoshanav/dragontail-test/controllers:RestaurantController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:restaurantId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
