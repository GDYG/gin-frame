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

// 添加用户
func CreateUserInfos(c *gin.Context, name string) *gorm.DB {
	data := dao.DB.Create(&User{Name: name})
	return data
}

// 通过id获取用户信息
func GetUserInfos(c *gin.Context, id int) (*gorm.DB, User) {
	result := User{}
	data := dao.DB.Model(&User{}).First(&result, id)
	return data, result
}

// 获取用户列表
func GetUserList(c *gin.Context) (*gorm.DB, []User) {
	result := make([]User, 0)
	data := dao.DB.Model(&User{}).Find(&result)
	return data, result
}

// 删除用户信息
func DeleteUserInfos(c *gin.Context, id int) *gorm.DB {
	data := dao.DB.Delete(&User{}, id)
	return data
}
