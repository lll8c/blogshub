package controller

import (
	"bloghub/domain"
	"bloghub/service"
	"bloghub/utils/ginx"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddUserHandler(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.AddUser(&user); err != nil {
		fmt.Println(err)
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func DeleteUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	if err := service.DeleteUser(id); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func BatchDeleteUserHandler(c *gin.Context) {
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
	if err := service.BatchDeleteUser(ids); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func UpdateUserHandler(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.UpdateUser(&user); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func GetUserHandler(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.AddUser(&user); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func GetUserListHandler(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.AddUser(&user); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func GetUserByPageHandler(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if err := service.AddUser(&user); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}
