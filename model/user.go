package model

type User struct {
	Id       int64  `gorm:"column:id" json:"id,omitempty"`
	Username string `gorm:"column:username" json:"username,omitempty"`
	Password string `gorm:"column:password" json:"password,omitempty"`
	Name     string `gorm:"column:name" json:"name,omitempty"`
	Avatar   string `gorm:"column:avatar" json:"avatar,omitempty"`
	Role     string `gorm:"column:role" json:"role,omitempty"`
	Sex      string `gorm:"column:sex" json:"sex,omitempty"`
	Phone    string `gorm:"column:phone" json:"phone,omitempty"`
	Email    string `gorm:"column:email" json:"email,omitempty"`
	Info     string `gorm:"column:info" json:"info,omitempty"`
	Birth    string `gorm:"column:birth" json:"birth,omitempty"`

	BlogCount    int64 `gorm:"-"  json:"blog_count,omitempty"`
	LikesCount   int64 `gorm:"-"  json:"likes_count,omitempty"`
	CollectCount int64 `gorm:"-"  json:"collect_count,omitempty"`

	Account `gorm:"-" json:"account"`
}

func (*User) TableName() string {
	return "user"
}

func InsertUser(u *User) error {
	return db.Create(&u).Error
}

func GetUserByName(username string) (user *User, err error) {
	err = db.Where("username = ?", username).First(&user).Error
	return
}

func UpdateUserById(u *User) error {
	return db.Model(&User{}).Where("id = ?", u.Id).Updates(u).Error
}

func DeleteUserById(id int64) error {
	return db.Where("id = ?", id).Delete(&User{}).Error
}

func BatchDeleteUserByIds(ids []int64) error {
	return db.Where("id in ?", ids).Delete(&User{}).Error
}

func GetUserById(id int64) (user *User, err error) {
	user = &User{}
	err = db.Where("id = ?", id).First(user).Error
	return
}

func GetUserList(user *User) (list []*User, err error) {
	query := db.Model(&user)
	if user.Username != "" {
		query = query.Where("username = ?", user.Username)
	}
	if user.Name != "" {
		query = query.Where("name = ?", user.Name)
	}
	err = query.Find(&list).Error
	return
}

func GetUserByPage(user *User, page int, pageSize int) (list []*User, err error) {
	query := db.Model(&user)
	if user.Username != "" {
		query = query.Where("username = ?", user.Username)
	}
	if user.Name != "" {
		query = query.Where("name = ?", user.Name)
	}
	err = query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&list).Error
	return
}
