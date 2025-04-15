package ginx

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserClaims struct {
	jwt.RegisteredClaims
	//声明你自己的要放进去token里面的数据
	ID   int64
	Role string
}

func CreateToken(id int64, role string) (string, error) {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			//设置过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
		Role: role,
		ID:   id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenStr, err
}
