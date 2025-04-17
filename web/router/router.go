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
	r.PUT("/updatePassword", controller.UpdatePassword)

	//用户管理
	user := r.Group("/user")
	user.POST("/add", controller.AddUserHandler)
	user.DELETE("/delete/:id", controller.DeleteUserHandler)
	user.DELETE("/delete/batch", controller.BatchDeleteUserHandler)
	user.PUT("/update", controller.UpdateUserHandler)
	user.GET("/selectById/:id", controller.GetUserHandler)
	user.GET("/selectAll", controller.GetUserListHandler)
	user.GET("/selectPage", controller.GetUserByPageHandler)

	//admin管理
	admin := r.Group("/admin")
	admin.POST("/add", controller.AddAdminHandler)
	admin.DELETE("/delete/:id", controller.DeleteAdminHandler)
	admin.DELETE("/delete/batch", controller.BatchDeleteAdminHandler)
	admin.PUT("/update", controller.UpdateAdminHandler)
	admin.GET("/selectById/:id", controller.GetAdminHandler)
	admin.GET("/selectAll", controller.GetAdminListHandler)
	admin.GET("/selectPage", controller.GetAdminByPageHandler)

	//分类管理
	category := r.Group("/category")
	category.POST("/add", controller.AddCategoryHandler)
	category.DELETE("/delete/:id", controller.DeleteCategoryHandler)
	category.DELETE("/delete/batch", controller.BatchDeleteCategoryHandler)
	category.PUT("/update", controller.UpdateCategoryHandler)
	category.GET("/selectById/:id", controller.GetCategoryHandler)
	category.GET("/selectAll", controller.GetCategoryListHandler)
	category.GET("/selectPage", controller.GetCategoryByPageHandler)

	//博客管理
	blog := r.Group("/blog")
	//后台管理
	blog.POST("/add", controller.AddBlogHandler)
	blog.DELETE("/delete/:id", controller.DeleteBlogHandler)
	blog.DELETE("/delete/batch", controller.BatchDeleteBlogHandler)
	blog.PUT("/update", controller.UpdateBlogHandler)
	blog.GET("/selectById/:id", controller.GetBlogHandler)
	blog.GET("/selectAll", controller.GetBlogListHandler)
	blog.GET("/selectPage", controller.GetBlogByPageHandler)
	//用户相关功能
	blog.POST("/updateReadCount/:id", controller.ReadBlogHandler)
	blog.GET("/selectUser", controller.GetBlogByUserHandler)
	blog.GET("/selectLike", controller.GetBlogByLikeHandler)
	blog.GET("/selectCollect", controller.GetBlogByCollectHandler)
	blog.GET("/selectComment", controller.GetBlogByCommentHandler)
	blog.GET("/selectTop", controller.GetTopBlogsHandler)
	blog.GET("/selectRecommend/:blogId", controller.GetRecommendBlogHandler)

	//活动管理
	activity := r.Group("/activity")
	//后台管理
	activity.POST("/add", controller.AddActivityHandler)
	activity.DELETE("/delete/:id", controller.DeleteActivityHandler)
	activity.DELETE("/delete/batch", controller.BatchDeleteActivityHandler)
	activity.PUT("/update", controller.UpdateActivityHandler)
	activity.GET("/selectById/:id", controller.GetActivityHandler)
	activity.GET("/selectAll", controller.GetActivitiesHandler)
	activity.GET("/selectPage", controller.GetActivitiesByPageHandler)
	//用户相关功能
	activity.POST("/updateReadCount/:id", controller.ReadActivityHandler)
	activity.GET("/selectUser", controller.GetUserSignActivityHandler)
	activity.GET("/selectLike", controller.GetUserLikeActivityHandler)
	activity.GET("/selectCollect", controller.GetUserCollectActivityHandler)
	activity.GET("/selectComment", controller.GetUserCommentActivityHandler)
	activity.GET("/selectTop", controller.GetTopActivitiesHandler)

	//活动报名管理
	activitySign := r.Group("/activitySign")
	activitySign.POST("/add", controller.AddActivitySignHandler)
	activitySign.DELETE("/delete/:id", controller.DeleteActivitySignHandler)
	activitySign.DELETE("/delete/batch", controller.BatchDeleteActivitySignHandler)
	activitySign.DELETE("/delete/user/:activityId/:userId", controller.DeleteUserActivitySignHandler)
	activitySign.GET("/selectPage", controller.GetActivitySignByPageHandler)

	//评论管理
	comment := r.Group("/comment")
	comment.POST("/add", controller.AddCommentHandler)
	comment.DELETE("/delete/:id", controller.DeleteCommentHandler)
	comment.DELETE("/delete/batch", controller.BatchDeleteCommentHandler)
	comment.PUT("/update", controller.UpdateCommentHandler)
	comment.GET("/selectById/:id", controller.GetCommentHandler)
	comment.GET("/selectAll", controller.GetCommentListHandler)
	comment.GET("/selectPage", controller.GetCommentByPageHandler)
	comment.GET("/selectForUser", controller.GetUserCommentListHandler)
	comment.GET("/selectCount", controller.CountCommentHandler)

	//点赞收藏
	r.POST("likes/set", controller.SetLikeHandler)
	r.POST("collect/set", controller.SetCollectHandler)

	//文件管理
	//files := r.Group("/files")
	//files.POST("upload", controller.UploadHandler)
	//files.POST("editor/upload", controller.EditorUploadHandler)
	//files.GET("likes/:flag", controller.AvatarPathHander)
	//files.DELETE("/:flag", controller.DelFileHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}
