package models

import (
	dao "gin-frame/dao"

	// "gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	Id   int `gorm:"primary_key"`
	Name string
}

func (User) TableName() string {
	return "user"
}

func GetUserInfos(c *gin.Context, name string) *gorm.DB {
	// dao.DB.AutoMigrate(&User{})
	data := dao.DB.Create(&User{Name: name})
	return data
}
