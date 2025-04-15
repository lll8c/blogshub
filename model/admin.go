package model

type Admin struct {
	Id       int64  `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Name     string `gorm:"column:name"`
	Avatar   string `gorm:"column:avatar"`
	Role     string `gorm:"column:role"`
	Phone    string `gorm:"column:phone"`
	Email    string `gorm:"column:email"`
}

func (*Admin) TableName() string {
	return "admin"
}

func InsertAdmin(a *Admin) error {
	return db.Create(&a).Error
}

func GetAdminByID(id int64) (admin *Admin, err error) {
	err = db.Where("id = ?", id).First(admin).Error
	return
}

func GetAdminByName(username string) (admin *Admin, err error) {
	err = db.Where("username = ?", username).First(admin).Error
	return
}

func GetAdminList() (list []*Admin, err error) {
	err = db.Find(&list).Error
	return
}

func DeleteAdminById(id int64) error {
	return db.Where("id = ?", id).Delete(&Admin{}).Error
}

func UpdateAdminById(admin *Admin) error {
	return db.Model(&Admin{}).Where("id = ?", admin.Id).Updates(admin).Error
}
