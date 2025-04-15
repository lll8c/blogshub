package domain

type User struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	Role         string `json:"role"`
	Sex          string `json:"sex"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Info         string `json:"info"`
	Birth        string `json:"birth"`
	BlogCount    int    `json:"blog_count"`
	LikesCount   int    `json:"likes_count"`
	CollectCount int    `json:"collect_count"`
}
