package controllers

import (
	"github.com/gin-gonic/gin"
)

type OrderController struct{}

type Search struct {
	Name string `json:"name"`
	Cid  int64  `json:"cid"`
}

func (o OrderController) GetList(c *gin.Context) {
	// name := c.Param("name") // 获取url参数 get query

	// 获取post参数 x-www-form-urlencoded
	// cid := c.PostForm("cid")
	// name := c.DefaultPostForm("name", "gdyg") // 赋值默认值
	// ResponseSuccess(c, 200, "success", "order list: "+cid+" "+name)
	// 获取post参数 x-www-form-urlencoded

	// 获取post参数 json类型
	// 实现一： map实现
	// param := make(map[string]interface{})
	// err := c.BindJSON(&param)
	// 实现一： map实现

	// 实现二：结构体实现
	search := &Search{}
	err := c.BindJSON(&search)
	// 实现二：结构体实现

	if err != nil {
		ResponseError(c, 4001, gin.H{"err": err})
		return
	}

	ResponseSuccess(c, 200, "success", search, 1)
	// 获取post参数 json类型

	// ResponseSuccess(c, 200, "success", []string{"gdyg", "gdyg1", "gdyg2"})
}
