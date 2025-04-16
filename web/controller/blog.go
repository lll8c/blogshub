package controller

import (
	"bloghub/domain"
	"bloghub/model"
	"bloghub/service"
	"bloghub/utils/ginx"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddBlogHandler(c *gin.Context) {
	var blog model.Blog
	if err := c.ShouldBind(&blog); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	account, err := ginx.GetCurrentUser(c)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	if account.Role == domain.USER {
		blog.UserId = account.Id
	}
	if err := service.AddBlog(&blog); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func DeleteBlogHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	if err := service.DeleteBlog(id); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)

}

func BatchDeleteBlogHandler(c *gin.Context) {
	rawData, err := c.GetRawData() // 获取原始请求体数据
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	var ids []int64
	if err := json.Unmarshal(rawData, &ids); err != nil { // 手动解析到切片
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.BatchDeleteBlog(ids); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func UpdateBlogHandler(c *gin.Context) {
	var blog model.Blog
	if err := c.ShouldBind(&blog); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.UpdateBlog(&blog); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)

}

// ReadBlogHandler 阅读量加1
func ReadBlogHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	if err := service.UpdateReadCount(id); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

// todo 获取文章详情
// GetBlogHandler 获取文章详情
func GetBlogHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	user, err := service.GetBlog(id)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, user)
}

func GetBlogListHandler(c *gin.Context) {
	var blog model.Blog
	if err := c.ShouldBind(&blog); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	list, err := service.GetAllBlog(&blog)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, list)
}

func GetBlogByPageHandler(c *gin.Context) {
	var query model.Blog
	if err := c.ShouldBind(&query); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	pageNumStr := c.DefaultQuery("pageNum", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageNumInt, _ := strconv.Atoi(pageNumStr)
	pageSizeInt, _ := strconv.Atoi(pageSizeStr)
	list, err := service.GetBlogByPage(&query, pageNumInt, pageSizeInt)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, list)
}

// GetBlogByUserHandler 分页查询当前用户的博客列表
func GetBlogByUserHandler(c *gin.Context) {

}

// GetBlogByLikeHandler 分页查询当前用户点赞的博客列表
func GetBlogByLikeHandler(c *gin.Context) {

}

// GetBlogByCollectHandler 分页查询当前用户收藏的博客列表
func GetBlogByCollectHandler(c *gin.Context) {

}

// GetBlogByCommentHandler 分页查询当前用户评论的博客列表
func GetBlogByCommentHandler(c *gin.Context) {

}

func GetTopBlogHandler(c *gin.Context) {

}
