package controllers

import (
	"github.com/gin-gonic/gin"
)

// 使用结构体来模拟类的行为，在go中没有类的概念，这样可以模块化代码，不会在同一个controllers下写同名的方法时报错
type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	name := c.Param("name")
	ResponseSuccess(c, 200, "success", "hello "+name, 0)
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

	ResponseError(c, 4004, "相关信息不存在！")

	// ResponseSuccess(c, 200, "success", []string{"gdyg", "gdyg1", "gdyg2"})
}
