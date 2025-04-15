package router

import (
	"bloghub/web/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//登录注册
	r.POST("/login", controller.LoginController)
	r.POST("/register", controller.RegisterController)

	//r.Use(midderware.JwtMidderware)
	//用户管理
	user := r.Group("/user")
	user.POST("/add", controller.AddUserHandler)
	user.POST("/delete/:id", controller.DeleteUserHandler)
	user.DELETE("/delete/batch", controller.BatchDeleteUserHandler)
	user.POST("/update", controller.UpdateUserHandler)
	user.GET("/selectById/:id", controller.GetUserHandler)
	user.POST("/selectAll", controller.GetUserListHandler)
	user.POST("/selectPage", controller.UpdateUserHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}
