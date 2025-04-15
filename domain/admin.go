package domain

type Admin struct {
	Account
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
