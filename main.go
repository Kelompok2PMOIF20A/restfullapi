package main

import (
	"fmt"
	Config "tes/config"
	"tes/model"
	"tes/route"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&model.User{})
	r := route.SetupRouter()
	//running
	r.Run()
}
