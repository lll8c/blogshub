package midderware

import (
	"bloghub/utils/ginx"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const TOKEN = "token"

func JwtMidderware(c *gin.Context) {
	var tokenStr string
	tokenStr = c.GetHeader(TOKEN)
	if tokenStr == "" {
		any, ok := c.Get(TOKEN)
		if !ok {
			return
		}
		tokenStr = any.(string)
	}
	//解析时传指针
	claims := &ginx.UserClaims{}
	//解析token
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		ginx.ResponseError(c, ginx.TokenInvalidErr)
		return
	}
	if token == nil || !token.Valid {
		ginx.ResponseError(c, ginx.TokenInvalidErr)
		return
	}
}
