package ginx

import (
	"bloghub/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 100000)),
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

func GetCurrentUser(c *gin.Context) (*model.Account, error) {
	value, ok := c.Get("user")
	if !ok {
		return nil, ParamLostErr
	}
	userClaim, _ := value.(*UserClaims)
	var account = &model.Account{}
	if userClaim.Role == "ADMIN" {
		admin, err := model.GetAdminByID(userClaim.ID)
		if err != nil {
			return nil, SystemErr
		}
		copier.Copy(account, admin)
	}
	if userClaim.Role == "USER" {
		user, err := model.GetUserById(userClaim.ID)
		if err != nil {
			return nil, SystemErr
		}
		copier.Copy(account, user)
	}
	return account, nil
}
