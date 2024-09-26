package models

import (
	"encoding/json"
	dao "gin-frame/dao"
	_redis "gin-frame/pkg/redis"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	// gorm.Model
	Id   int `gorm:"primary_key"`
	Name string
}

func (User) TableName() string {
	return "users"
}

// 添加用户
func CreateUserInfos(c *gin.Context, name string) *gorm.DB {
	data := dao.DB.Create(&User{Name: name})
	return data
}

// 通过id获取用户信息
func GetUserInfos(c *gin.Context, id int) (*gorm.DB, User) {
	result := User{}
	rdb := _redis.InitRedis()
	defer rdb.Close()
	key := "user_" + strconv.Itoa(id)
	jsonStr, err := rdb.Get(c, key).Bytes()

	if err == redis.Nil {
		data := dao.DB.Model(&User{}).First(&result, id)
		value, _ := json.Marshal(result)
		rdb.Set(c, key, string(value), 0)

		return data, result
	}

	errJson := json.Unmarshal(jsonStr, &result)
	if errJson != nil {
		panic(errJson)
	}
	return dao.DB, result
}

// 获取用户列表
func GetUserList(c *gin.Context) (*gorm.DB, []User) {
	result := make([]User, 0)
	data := dao.DB.Model(&User{}).Find(&result)
	return data, result
}

// 更新用户信息
func UpdateUserInfos(c *gin.Context, id int, name string) (*gorm.DB, []User) {
	var users []User
	data := dao.DB.Model(&users).Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}}}).Where("id = ?", id).Updates(map[string]interface{}{"name": name})
	return data, users
}

// 删除用户信息
func DeleteUserInfos(c *gin.Context, id int) (*gorm.DB, []User) {
	var users []User
	data := dao.DB.Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}}}).Delete(&users, id)
	return data, users
}
