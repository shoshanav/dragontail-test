package models

import (
	"fmt"
	"github.com/dragontail/db"
)

type Restaurant struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Phone	   string 	  `json:"phone"`
	Location   string     `json:"location"`
	Address    string     `json:"address"`
}

func GetAllRestaurants() ([]Restaurant, error) {
	var restaurants []Restaurant
	err := db.DBCon.Order("id asc").Find(&restaurants).Error
	return restaurants, err
}

func SearchRestaurantsByName(name string) ([]Restaurant, error) {
	var restaurants []Restaurant
	err := db.DBCon.Where(fmt.Sprintf("name LIKE '%%%s%%'", name)).Find(&restaurants).Error
	return restaurants, err
}

func UpdateRestaurant(id int, res Restaurant) error {
	err := db.DBCon.Save(&res).Error
	return err
}

func DeleteRestaurant(id string) error {
	var res Restaurant
	err := db.DBCon.Where("id=?", id).Delete(&res).Error
	return err
}

