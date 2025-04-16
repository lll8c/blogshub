package controller

import (
	"bloghub/model"
	"bloghub/service"
	"bloghub/utils/ginx"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddCategoryHandler(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBind(&category); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.AddCategory(&category); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return

	}
	ginx.ResponseSuccess(c, nil)
}

func DeleteCategoryHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	if err := service.DeleteCategoryById(id); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func BatchDeleteCategoryHandler(c *gin.Context) {
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
	if err := service.BatchDeleteCategory(ids); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func UpdateCategoryHandler(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBind(&category); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.UpdateCategory(&category); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func GetCategoryHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	user, err := service.GetCategory(id)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, user)
}

func GetCategoryListHandler(c *gin.Context) {
	users, err := service.GetAllCategory()
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, users)
}

func GetCategoryByPageHandler(c *gin.Context) {
	pageNumStr := c.DefaultQuery("pageNum", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	pageNumInt, _ := strconv.Atoi(pageNumStr)
	pageSizeInt, _ := strconv.Atoi(pageSizeStr)
	users, err := service.GetCategoryByPage(pageNumInt, pageSizeInt)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, users)
}
