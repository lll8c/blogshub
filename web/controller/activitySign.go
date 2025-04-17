package controller

import (
	"bloghub/model"
	"bloghub/service"
	"bloghub/utils/ginx"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func AddActivitySignHandler(c *gin.Context) {
	var actSign model.ActivitySign
	if err := c.ShouldBind(&actSign); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	user, err := ginx.GetCurrentUser(c)
	if err != nil {
		ginx.ResponseError(c, ginx.UserNotExistErr)
		return
	}
	actSign.UserId = user.Id
	actSign.Time = time.Now().String()
	if err := service.AddActivitySign(&actSign); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func DeleteActivitySignHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	if err := service.DeleteActivitySign(id); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func BatchDeleteActivitySignHandler(c *gin.Context) {
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
	if err := service.BatchDeleteActivitySign(ids); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func DeleteUserActivitySignHandler(c *gin.Context) {
	activityIdStr := c.Param("activityId")
	useIdStr := c.Param("useId")
	activityID, err := strconv.ParseInt(activityIdStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	useId, err := strconv.ParseInt(useIdStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	if err := service.DeleteUserActivitySign(activityID, useId); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func GetActivitySignByPageHandler(c *gin.Context) {
	var query model.ActivitySign
	if err := c.ShouldBind(&query); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	pageNumStr := c.DefaultQuery("pageNum", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageNumInt, _ := strconv.Atoi(pageNumStr)
	pageSizeInt, _ := strconv.Atoi(pageSizeStr)
	list, err := service.GetActivitySignByPage(&query, pageNumInt, pageSizeInt)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, list)
}
