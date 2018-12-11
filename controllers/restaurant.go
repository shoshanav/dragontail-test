package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/dragontail/lib"
	"github.com/dragontail/models"
	"strconv"
	"strings"
)

type RestaurantController struct {
	beego.Controller
}

// @router / [get]
func (r *RestaurantController) GetAll() {
	restaurants, err := models.GetAllRestaurants()
	if err != nil {
		r.Data["json"] = map[string]string {"error": err.Error()}
	} else {
		err := geoCode(restaurants)
		if err != nil {
			r.Data["json"] = map[string]string {"error": err.Error()}
		}
		r.Data["json"] = map[string]interface{} {"restaurants": restaurants}
	}
	r.ServeJSON()
}

// @router /:restaurantId [put]
func (r *RestaurantController) Put() {
	resId,_ := strconv.Atoi(r.Ctx.Input.Param(":restaurantId"))
	var restaurant models.Restaurant
	json.Unmarshal(r.Ctx.Input.RequestBody, &restaurant)
	err := models.UpdateRestaurant(resId, restaurant)
	if err != nil {
		r.Data["json"] = map[string]string {"error": err.Error()}
	} else {
		r.Data["json"] = map[string]bool {"success": true}
	}
	r.ServeJSON()
}

// @router /:restaurantId [delete]
func (r *RestaurantController) Delete() {
	restaurantId := r.Ctx.Input.Param(":restaurantId")
	err := models.DeleteRestaurant(restaurantId)
	if err != nil {
		r.Data["json"] = map[string]string {"error": err.Error()}
	} else {
		r.Data["json"] = map[string]bool {"success": true}
	}
	r.ServeJSON()
}

func geoCode(restaurants []models.Restaurant) error {
	for i, res := range restaurants {
		locations:= strings.Split(res.Location, "/")
		if len(locations) < 2 {
			continue
		}
		address,err := lib.GetAddress(locations[0], locations[1])
		if err != nil {
			continue
		}
		restaurants[i].Address = address
	}
	return nil
}

//// @router /search [get]
//func (r *RestaurantController) Search() {
//	restaurantName := r.GetString("restaurantName")
//	restaurants,err := models.SearchRestaurantsByName(restaurantName)
//	if err != nil {
//		r.Data["json"] = map[string]string {"error": err.Error()}
//	} else {
//		r.Data["json"] = map[string]interface{} {"restaurants": restaurants}
//	}
//	r.ServeJSON()
//}



