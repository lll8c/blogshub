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

func AddAdminHandler(c *gin.Context) {
	var admin model.Admin
	if err := c.ShouldBind(&admin); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.AddAdmin(&admin); err != nil {
		fmt.Println(err)
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func DeleteAdminHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	if err := service.DeleteAdmin(id); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func BatchDeleteAdminHandler(c *gin.Context) {
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
	if err := service.BatchDeleteAdmin(ids); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func UpdateAdminHandler(c *gin.Context) {
	var admin model.Admin
	if err := c.ShouldBind(&admin); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.UpdateAdmin(&admin); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func GetAdminHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	user, err := service.GetAdmin(id)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, user)
}

func GetAdminListHandler(c *gin.Context) {
	var user model.Admin
	if err := c.ShouldBind(&user); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	users, err := service.GetAdminList(&user)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, users)
}

// GetAdminByPageHandler 分页获取用户列表
func GetAdminByPageHandler(c *gin.Context) {
	var user model.Admin
	if err := c.ShouldBind(&user); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	pageNumStr := c.DefaultQuery("pageNum", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	pageNumInt, _ := strconv.Atoi(pageNumStr)
	pageSizeInt, _ := strconv.Atoi(pageSizeStr)
	users, err := service.GetAdminByPage(&user, pageNumInt, pageSizeInt)
	if err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, users)
}
