package main

import (
	"github.com/astaxie/beego"
	"github.com/dragontail/db"
	_ "github.com/dragontail/db"
	_ "github.com/dragontail/routers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/saturn4er/beego-assets"
)

const (
	CSVFILE  = "restaurants.csv"
)

func main() {
	var err error
	db.DBCon, err = gorm.Open("postgres", "postgres://localhost:5432/dragontail?sslmode=disable")
	db.DBCon.LogMode(true)
	if err != nil {
		panic(err)
	}
	defer db.DBCon.Close()

	err = db.InitDB(CSVFILE)
	if err != nil {
		panic(err)
	}

	beego.Run()

	//db.AutoMigrate(&models.Restaurant{})
}


