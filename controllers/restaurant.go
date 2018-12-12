package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/shoshanav/dragontail-test/lib"
	"github.com/shoshanav/dragontail-test/models"
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
		err := geoCodeMutlti(restaurants)
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
	restaurant, err := models.UpdateRestaurant(resId, restaurant)
	if err != nil {
		r.Data["json"] = map[string]string {"error": err.Error()}
	} else {
		geoCode(&restaurant)
		r.Data["json"] = map[string]interface{} {"restaurant": restaurant}
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
		r.Data["json"] = map[string]string {"success": "Delete Succeeded"}
	}
	r.ServeJSON()
}

func geoCodeMutlti(restaurants []models.Restaurant) error {
	var err error
	for i,_ := range restaurants {
		err = geoCode(&restaurants[i])
	}
	return err
}

func geoCode(res *models.Restaurant) error {
	locations:= strings.Split(res.Location, "/")
	if len(locations) < 2 {
		return nil
	}
	address,err := lib.GetAddress(locations[0], locations[1])
	if err != nil {
		return err
	}
	res.Address = address
	return nil
}


