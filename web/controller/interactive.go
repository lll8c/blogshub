package controller

import (
	"bloghub/model"
	"bloghub/service"
	"bloghub/utils/ginx"
	"github.com/gin-gonic/gin"
)

// SetLikeHandler 点赞或取消点赞
func SetLikeHandler(c *gin.Context) {
	var like model.Likes
	if err := c.ShouldBind(&like); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	user, err := ginx.GetCurrentUser(c)
	if err != nil {
		ginx.ResponseError(c, ginx.UserNotExistErr)
		return
	}
	like.UserId = user.Id
	if err := service.SetLikes(&like); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

// SetCollectHandler 收藏或取消收藏
func SetCollectHandler(c *gin.Context) {
	var collect model.Collect
	if err := c.ShouldBind(&collect); err != nil {
		ginx.ResponseError(c, ginx.ParamErr)
		return
	}
	user, err := ginx.GetCurrentUser(c)
	if err != nil {
		ginx.ResponseError(c, ginx.UserNotExistErr)
		return
	}
	collect.UserId = user.Id
	if err := service.SetCollect(&collect); err != nil {
		ginx.ResponseError(c, err)
		return
	}
	ginx.ResponseSuccess(c, nil)
}
