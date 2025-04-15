package domain

const (
	ADMIN = "ADMIN"
	USER  = "USER"
)

// Account 账号信息
type Account struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	NewPassword string `json:"newPassword"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Token       string `json:"token"`
}
