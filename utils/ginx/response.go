package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

func ResponseError(ctx *gin.Context, err error) {
	var rd *ResponseData
	myErr, ok := err.(*MyError)
	fmt.Println(ok, myErr)
	if !ok {
		rd = &ResponseData{
			Code: SystemErr.Code,
			Msg:  SystemErr.Msg,
		}
	} else {
		rd = &ResponseData{
			Code: myErr.Code,
			Msg:  myErr.Msg,
		}
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: Success.Code,
		Msg:  Success.Msg,
		Data: data,
	}
	ctx.JSON(http.StatusOK, rd)
}
