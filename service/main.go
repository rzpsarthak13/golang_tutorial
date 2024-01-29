package main

import (
	"fmt"
	"service/Config"
	"service/Models"
	"service/Routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL((Config.BuildDBConfig())))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Customer{}, &Models.Order{})
	r := Routes.SetUpRouter()
	r.Run()
}
