package controller

import (
	"bloghub/model"
	"bloghub/service"
	"bloghub/utils/ginx"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddActivityHandler(c *gin.Context) {
	var activity model.Activity
	if err := c.ShouldBind(&activity); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.AddActivity(&activity); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func DeleteActivityHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	if err := service.DeleteActivity(id); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func BatchDeleteActivityHandler(c *gin.Context) {
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
	if err := service.BatchDeleteActivity(ids); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func UpdateActivityHandler(c *gin.Context) {
	var activity model.Activity
	if err := c.ShouldBind(&activity); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.UpdateActivity(&activity); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

// GetActivityHandler 获取活动详情
func GetActivityHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	user, err := ginx.GetCurrentUser(c)
	if err != nil {
		ginx.ResponseError(c, ginx.UserNotExistErr)
		return
	}
	activity, err := service.GetActivity(id, user)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, activity)
}

func GetActivitiesHandler(c *gin.Context) {
	var query model.Activity
	if err := c.ShouldBind(&query); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	list, err := service.GetActivities(&query)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, list)
}

func GetActivitiesByPageHandler(c *gin.Context) {
	var query model.Activity
	if err := c.ShouldBind(&query); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	pageNumStr := c.DefaultQuery("pageNum", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageNumInt, _ := strconv.Atoi(pageNumStr)
	pageSizeInt, _ := strconv.Atoi(pageSizeStr)
	user, err := ginx.GetCurrentUser(c)
	if err != nil {
		ginx.ResponseError(c, ginx.UserAccountErr)
		return
	}
	list, err := service.GetActivityByPage(&query, pageNumInt, pageSizeInt, user)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, list)
}

func ReadActivityHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	if err := service.ReadActivityCount(id); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

// GetUserSignActivityHandler 分页查询当前用户报名的活动
func GetUserSignActivityHandler(c *gin.Context) {
	var query model.Activity
	if err := c.ShouldBind(&query); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	pageNumStr := c.DefaultQuery("pageNum", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageNumInt, _ := strconv.Atoi(pageNumStr)
	pageSizeInt, _ := strconv.Atoi(pageSizeStr)
	user, err := ginx.GetCurrentUser(c)
	if err != nil {
		ginx.ResponseError(c, ginx.UserNotExistErr)
		return
	}
	query.UserId = user.Id
	list, err := service.GetUserSignActivity(&query, pageNumInt, pageSizeInt)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, list)
}

// GetUserLikeActivityHandler 分页查询当前用户点赞的博客列表
func GetUserLikeActivityHandler(c *gin.Context) {
	var query model.Activity
	if err := c.ShouldBind(&query); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	pageNumStr := c.DefaultQuery("pageNum", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageNumInt, _ := strconv.Atoi(pageNumStr)
	pageSizeInt, _ := strconv.Atoi(pageSizeStr)
	user, err := ginx.GetCurrentUser(c)
	if err != nil {
		ginx.ResponseError(c, ginx.UserNotExistErr)
		return
	}
	query.UserId = user.Id
	list, err := service.GetUserLikeActivity(&query, pageNumInt, pageSizeInt)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, list)

}

// GetUserCollectActivityHandler 分页查询当前用户收藏的博客列表
func GetUserCollectActivityHandler(c *gin.Context) {
	var query model.Activity
	if err := c.ShouldBind(&query); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	pageNumStr := c.DefaultQuery("pageNum", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageNumInt, _ := strconv.Atoi(pageNumStr)
	pageSizeInt, _ := strconv.Atoi(pageSizeStr)
	user, err := ginx.GetCurrentUser(c)
	if err != nil {
		ginx.ResponseError(c, ginx.UserNotExistErr)
		return
	}
	query.UserId = user.Id
	list, err := service.GetUserCollectActivity(&query, pageNumInt, pageSizeInt)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, list)
}

// GetUserCommentActivityHandler 分页查询当前用户评论的博客列表
func GetUserCommentActivityHandler(c *gin.Context) {
	var query model.Activity
	if err := c.ShouldBind(&query); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	pageNumStr := c.DefaultQuery("pageNum", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageNumInt, _ := strconv.Atoi(pageNumStr)
	pageSizeInt, _ := strconv.Atoi(pageSizeStr)
	user, err := ginx.GetCurrentUser(c)
	if err != nil {
		ginx.ResponseError(c, ginx.UserNotExistErr)
		return
	}
	query.UserId = user.Id
	list, err := service.GetUserCommentActivity(&query, pageNumInt, pageSizeInt)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, list)
}

// GetTopActivitiesHandler 热门活动榜单
func GetTopActivitiesHandler(c *gin.Context) {
	list, err := service.GetTopActivities()
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, list)
}
