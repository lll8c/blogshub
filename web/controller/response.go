package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code MyCode      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

func ResponseError(ctx *gin.Context, code MyCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code MyCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context) {
	rd := &ResponseData{
		Code: Success,
		Msg:  Success.Msg(),
	}
	ctx.JSON(http.StatusOK, rd)
}
