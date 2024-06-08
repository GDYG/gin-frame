package main

import (
	"gin-frame/dao"
	"gin-frame/models"
	"gin-frame/router"
)

func main() {
	db := dao.InitDB()
	db.AutoMigrate(&models.User{})

	r := router.Router()
	r.Run(":9999")
}
