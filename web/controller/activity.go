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

func GetActivitiesByUserHandler(context *gin.Context) {
	
}

func GetActivitiesByLikeHandler(context *gin.Context) {
	
}

func GetActivitiesByCollectHandler(context *gin.Context) {
	
}

func GetActivitiesByCommentHandler(context *gin.Context) {
	
}

func GetTopActivitiesHandler(context *gin.Context) {
	
}