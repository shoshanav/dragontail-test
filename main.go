package main

import (
	"github.com/astaxie/beego"
	"github.com/shoshanav/dragontail-test/db"
	_ "github.com/shoshanav/dragontail-test/db"
	_ "github.com/shoshanav/dragontail-test/routers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/saturn4er/beego-assets"
	"os"
)

const (
	CSVFILE  = "restaurants.csv"
)

func main() {
	var err error
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgres://localhost:5432/dragontail?sslmode=disable"
	}
	db.DBCon, err = gorm.Open("postgres", dbUrl)
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
	}


