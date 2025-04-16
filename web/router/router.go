package router

import (
	"bloghub/web/controller"
	"bloghub/web/midderware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//登录注册
	r.POST("/login", controller.LoginController)
	r.POST("/register", controller.RegisterController)

	r.Use(midderware.JwtMidderware)
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

	//博客管理
	blog := r.Group("/blog")
	//后台管理
	blog.POST("/add", controller.AddBlogHandler)
	blog.DELETE("/delete/:id", controller.DeleteBlogHandler)
	blog.DELETE("/delete/batch", controller.BatchDeleteBlogHandler)
	blog.POST("/update", controller.UpdateBlogHandler)
	blog.GET("/selectById/:id", controller.GetBlogHandler)
	blog.GET("/selectAll", controller.GetBlogListHandler)
	blog.GET("/selectPage", controller.GetBlogByPageHandler)
	//用户相关功能
	blog.POST("/updateReadCount/:id", controller.ReadBlogHandler)
	blog.GET("/selectUser", controller.GetBlogByUserHandler)
	blog.GET("/selectLike", controller.GetBlogByLikeHandler)
	blog.GET("/selectCollect", controller.GetBlogByCollectHandler)
	blog.GET("/selectComment", controller.GetBlogByCommentHandler)
	blog.GET("/selectTop", controller.GetTopBlogHandler)

	//活动管理
	activity := r.Group("/activity")
	//后台管理
	activity.POST("/add", controller.AddActivityHandler)
	activity.DELETE("/delete/:id", controller.DeleteActivityHandler)
	activity.DELETE("/delete/batch", controller.BatchDeleteActivityHandler)
	activity.POST("/update", controller.UpdateActivityHandler)
	activity.GET("/selectById/:id", controller.GetActivityHandler)
	activity.GET("/selectAll", controller.GetActivitiesHandler)
	activity.GET("/selectPage", controller.GetActivitiesByPageHandler)
	//用户相关功能
	activity.POST("/updateReadCount/:id", controller.ReadActivityHandler)
	activity.GET("/selectUser", controller.GetActivitiesByUserHandler)
	activity.GET("/selectLike", controller.GetActivitiesByLikeHandler)
	activity.GET("/selectCollect", controller.GetActivitiesByCollectHandler)
	activity.GET("/selectComment", controller.GetActivitiesByCommentHandler)
	activity.GET("/selectTop", controller.GetTopActivitiesHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}
