package model

type User struct {
	Id       int64  `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Name     string `gorm:"column:name"`
	Avatar   string `gorm:"column:avatar"`
	Role     string `gorm:"column:role"`
	Sex      string `gorm:"column:sex"`
	Phone    string `gorm:"column:phone"`
	Email    string `gorm:"column:email"`
	Info     string `gorm:"column:info"`
	Birth    string `gorm:"column:birth"`
}

func InsertUser(u *User) error {
	return db.Create(&u).Error
}

func GetUserByName(username string) (user *User, err error) {
	user = &User{}
	err = db.Where("username = ?", username).First(user).Error
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

func GetUserList() (list []User, err error) {
	err = db.Find(&list).Error
	return
}
