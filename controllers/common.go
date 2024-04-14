package controllers

import (
	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

type JsonErrStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

// 统一返回格式：成功
func ResponseSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonStruct{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}
	c.JSON(200, json)
}

// 统一返回格式：失败
func ResponseError(c *gin.Context, code int, msg interface{}) {
	json := &JsonErrStruct{
		Code: code,
		Msg:  msg,
	}
	c.JSON(200, json)
}
