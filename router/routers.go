package router

import (
	"gin-frame/controllers"
	"gin-frame/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
	// 	SkipPaths: []string{},
	// }))

	// 添加中间件来记录请求日志
	r.Use(logger.RequestLogger())
	r.Use(logger.RecoveryMiddleware())

	user := r.Group("/user")
	{
		user.GET("/:name", controllers.UserController{}.CreateUser)

		user.POST("/list", controllers.UserController{}.GetList)

		user.PUT("/update", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "success",
				"data": "update success",
			})
		})

		user.DELETE("/delete", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "success",
				"data": "delete success",
			})
		})
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}

	return r
}
