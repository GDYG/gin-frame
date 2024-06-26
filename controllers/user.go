package controllers

import (
	"fmt"
	. "gin-frame/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 使用结构体来模拟类的行为，在go中没有类的概念，这样可以模块化代码，不会在同一个controllers下写同名的方法时报错
type UserController struct{}

func (u UserController) CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		ResponseError(c, 4001, "name is empty")
		return
	}
	data := CreateUserInfos(c, name)
	if data.Error != nil {
		ResponseError(c, 4001, gin.H{"err": data.Error})
		return
	}
	ResponseSuccess(c, 200, "create user", gin.H{
		"name": name,
	}, 1)
}

func (u UserController) GetList(c *gin.Context) {
	// defer recover panic 异常处理
	// Call the handlePanic function
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		ResponseError(c, 500, "系统异常")
	// 	}
	// }()
	// num1 := 10
	// num2 := 0
	// num3 := num1 / num2
	// ResponseError(c, 4004, num3)
	// defer recover panic 异常处理

	data, result := GetUserList(c)

	//打印result
	fmt.Println(result)

	if data.Error != nil {
		ResponseError(c, 4004, "获取列表失败")
		return
	}

	ResponseSuccess(c, 200, "success", result, data.RowsAffected)
}

func (u UserController) GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, result := GetUserInfos(c, id)

	if data.Error != nil {
		ResponseError(c, 4004, "获取用户信息失败")
		return
	}

	ResponseSuccess(c, 200, "success", result, data.RowsAffected)
}

func (u UserController) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.PostForm("name")
	data, user := UpdateUserInfos(c, id, name)
	if data.Error != nil {
		ResponseError(c, 4004, "更新用户信息失败")
		return
	}
	ResponseSuccess(c, 200, "success", user, data.RowsAffected)
}

func (u UserController) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, user := DeleteUserInfos(c, id)
	// 输出id
	fmt.Println(id)
	if data.Error != nil {
		ResponseError(c, 4004, "删除用户信息失败")
		return
	}
	ResponseSuccess(c, 200, "success", user, data.RowsAffected)
}
