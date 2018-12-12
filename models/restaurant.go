package models

import (
	"github.com/shoshanav/dragontail-test/db"
)

type Restaurant struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Phone	   string 	  `json:"phone"`
	Location   string     `json:"location"`
	Address    string     `json:"address" gorm:"-"`
}

func GetAllRestaurants() ([]Restaurant, error) {
	var restaurants []Restaurant
	err := db.DBCon.Order("id asc").Find(&restaurants).Error
	return restaurants, err
}

func UpdateRestaurant(id int, res Restaurant) (Restaurant, error) {
	res.ID = id
	err := db.DBCon.Save(&res).Error
	return res, err
}

func DeleteRestaurant(id string) error {
	var res Restaurant
	err := db.DBCon.Where("id=?", id).Delete(&res).Error
	return err
}

