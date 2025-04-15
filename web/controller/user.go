package controller

import (
	"bloghub/domain"
	"bloghub/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddUserHandler(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		ResponseError(c, ParamErr)
		return
	}
	if err := service.AddUser(&user); err != nil {
		ResponseError(c, SystemErr)
	}
	ResponseSuccess(c)
}

func DeleteUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, ParamErr)
	}
	if err := service.DeleteUser(id); err != nil {
		ResponseError(c, SystemErr)
	}
	ResponseSuccess(c)
}

func BatchDeleteUserHandler(c *gin.Context) {
	rawData, err := c.GetRawData() // 获取原始请求体数据
	if err != nil {
		ResponseError(c, ParamErr)
		return
	}
	var ids []int64
	if err := json.Unmarshal(rawData, &ids); err != nil { // 手动解析到切片
		ResponseError(c, ParamErr)
		return
	}
	if err := service.BatchDeleteUser(ids); err != nil {
		ResponseError(c, SystemErr)
	}
	ResponseSuccess(c)
}

func UpdateUserHandler(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		ResponseError(c, ParamErr)
		return
	}
	if err := service.UpdateUser(&user); err != nil {
		ResponseError(c, SystemErr)
	}
	ResponseSuccess(c)
}

func GetUserHandler(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		ResponseError(c, ParamErr)
		return
	}
	if err := service.AddUser(&user); err != nil {
		ResponseError(c, SystemErr)
	}
	ResponseSuccess(c)
}

func GetUserListHandler(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		ResponseError(c, ParamErr)
		return
	}
	if err := service.AddUser(&user); err != nil {
		ResponseError(c, SystemErr)
	}
	ResponseSuccess(c)
}

func GetUserByPageHandler(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		ResponseError(c, ParamErr)
		return
	}
	if err := service.AddUser(&user); err != nil {
		ResponseError(c, SystemErr)
	}
	ResponseSuccess(c)
}
