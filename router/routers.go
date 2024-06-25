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

	// 前缀
	user := r.Group("/user")
	{
		// 增
		user.POST("/create", controllers.UserController{}.CreateUser)

		// 查
		user.GET("/info/:id", controllers.UserController{}.GetUserInfo)
		user.GET("/list", controllers.UserController{}.GetList)

		// 改
		user.PUT("/update", controllers.UserController{}.UpdateUser)

		// 删
		user.DELETE("/delete/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "success",
				"data": "delete success",
			})
		})
	}

	// 其他前缀
	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}

	return r
}
