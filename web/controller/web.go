package controller

import (
	"bloghub/domain"
	"bloghub/service"
	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var account domain.Account
	if err := c.ShouldBind(&account); err != nil {
		ResponseError(c, ParamErr)
		return
	}
	if account.Username == "" || account.Password == "" || account.Role == "" {
		ResponseError(c, ParamLostErr)
		return
	}
	var err error
	if account.Role == "ADMIN" {
		err = service.LoginAdmin(&account)
	} else if account.Role == "USER" {
		err = service.LoginUser(&account)
	}
	if err != nil {
		ResponseError(c, SystemErr)
	}
	ResponseSuccess(c)
}

func SignUpController(c *gin.Context) {
	var account domain.Account
	if err := c.ShouldBind(&account); err != nil {
		ResponseError(c, ParamErr)
		return
	}
	if account.Username == "" || account.Password == "" || account.Role == "" {
		ResponseError(c, ParamLostErr)
		return
	}
	var err error
	if account.Role == "ADMIN" {
		err = service.LoginAdmin(&account)
	} else if account.Role == "USER" {
		err = service.LoginUser(&account)
	}
	if err != nil {
		ResponseError(c, SystemErr)
	}
	ResponseSuccess(c)
}
