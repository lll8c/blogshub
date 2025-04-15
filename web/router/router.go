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
	user.DELETE("/delete/:id", controller.DeleteUserHandler)
	user.DELETE("/delete/batch", controller.BatchDeleteUserHandler)
	user.POST("/update", controller.UpdateUserHandler)
	user.GET("/selectById/:id", controller.GetUserHandler)
	user.GET("/selectAll", controller.GetUserListHandler)
	user.GET("/selectPage", controller.GetUserByPageHandler)

	//分类管理
	category := r.Group("/category")
	category.POST("/add", controller.AddCategoryHandler)
	category.DELETE("/delete/:id", controller.DeleteCategoryHandler)
	category.DELETE("/delete/batch", controller.BatchDeleteCategoryHandler)
	category.POST("/update", controller.UpdateCategoryHandler)
	category.GET("/selectById/:id", controller.GetCategoryHandler)
	category.GET("/selectAll", controller.GetCategoryListHandler)
	category.GET("/selectPage", controller.GetCategoryByPageHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}
