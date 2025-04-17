package controller

import (
	"bloghub/model"
	"bloghub/service"
	"bloghub/utils/ginx"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddCommentHandler(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBind(&comment); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	user, err := ginx.GetCurrentUser(c)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamLostErr)
		return

	}
	comment.UserId = user.Id
	if err := service.AddComment(&comment); err != nil {
		fmt.Println(err)
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func DeleteCommentHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	if err := service.DeleteComment(id); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func BatchDeleteCommentHandler(c *gin.Context) {
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
	if err := service.BatchDeleteComment(ids); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func UpdateCommentHandler(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBind(&comment); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.UpdateComment(&comment); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func GetCommentHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	comment, err := service.GetComment(id)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, comment)
}

func GetCommentListHandler(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBind(&comment); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	list, err := service.GetCommentList(&comment)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, list)
}

// GetCommentByPageHandler 分页获取列表
func GetCommentByPageHandler(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBind(&comment); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	pageNumStr := c.DefaultQuery("pageNum", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	pageNumInt, _ := strconv.Atoi(pageNumStr)
	pageSizeInt, _ := strconv.Atoi(pageSizeStr)
	list, err := service.GetCommentListByPage(&comment, pageNumInt, pageSizeInt)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, list)
}

func GetUserCommentListHandler(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBind(&comment); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	list, err := service.GetUserCommentList(&comment)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, list)
}

func CountCommentHandler(c *gin.Context) {
	fidStr := c.Query("fid")
	module := c.Query("module")
	fid, err := strconv.ParseInt(fidStr, 10, 64)
	if err != nil || module == "" {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	count, err := service.CountCommentByFid(fid, module)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, count)
}
