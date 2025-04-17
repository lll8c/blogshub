package model

type Admin struct {
	Id       int64  `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Name     string `gorm:"column:name" json:"name"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Role     string `gorm:"column:role" json:"role"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Email    string `gorm:"column:email" json:"email"`

	Account `gorm:"-"`
}

func (*Admin) TableName() string {
	return "admin"
}

func InsertAdmin(a *Admin) error {
	return db.Create(&a).Error
}

func DeleteAdminById(id int64) error {
	return db.Where("id = ?", id).Delete(&Admin{}).Error
}

func BatchDeleteAdminByIds(ids []int64) error {
	return db.Where("id in ? ", ids).Delete(&Admin{}).Error
}

func UpdateAdminById(admin *Admin) error {
	return db.Model(&Admin{}).Where("id = ?", admin.Id).Updates(&admin).Error
}

func GetAdminByID(id int64) (admin *Admin, err error) {
	err = db.Where("id = ?", id).First(&admin).Error
	return
}

func GetAdminByName(username string) (admin *Admin, err error) {
	err = db.Where("username = ?", username).First(&admin).Error
	return
}

func GetAdminList(admin *Admin) (list []*Admin, err error) {
	query := db.Model(&admin)
	if admin.Username != "" {
		query = query.Where("username = ?", admin.Username)
	}
	if admin.Name != "" {
		query = query.Where("name = ?", admin.Name)
	}
	err = query.Find(&list).Error
	return
}

func GetAdminByPage(admin *Admin, page int, pageSize int) (list []*Admin, err error) {
	query := db.Model(&admin)
	if admin.Username != "" {
		query = query.Where("username = ?", admin.Username)
	}
	if admin.Name != "" {
		query = query.Where("name = ?", admin.Name)
	}
	err = query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&list).Error
	return
}
