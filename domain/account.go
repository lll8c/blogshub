package domain

const (
	ADMIN = "ADMIN"
	USER  = "USER"
)

type Account struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	NewPassword string `json:"new_password"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	token       string `json:"token"`
}
