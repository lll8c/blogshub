package controller

import (
	"bloghub/domain"
	"bloghub/service"
	"bloghub/utils/ginx"
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var account domain.Account
	if err := c.ShouldBind(&account); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if account.Username == "" || account.Password == "" || account.Role == "" {
		ginx.ResponseError(c, ginx.ParamLostErr)
		return
	}
	if account.Role == "ADMIN" {
		admin, err := service.LoginAdmin(&account)
		if err != nil {
			fmt.Println(err)
			ginx.ResponseError(c, err)
			return
		}
		ginx.ResponseSuccess(c, admin)
	}
	if account.Role == "USER" {
		user, err := service.LoginUser(&account)
		if err != nil {
			ginx.ResponseError(c, err)
			return
		}
		ginx.ResponseSuccess(c, user)
	}
}

func RegisterController(c *gin.Context) {
	var account domain.Account
	if err := c.ShouldBind(&account); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	if account.Username == "" || account.Password == "" || account.Role == "" {
		ginx.ResponseError(c, ginx.ParamLostErr)
		return
	}
	var err error
	//只能注册普通用户
	if account.Role != "USER" {
		ginx.ResponseError(c, ginx.ParamErr)
	}
	err = service.RegisterUser(&account)
	if err != nil {
		fmt.Println(err)
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}
