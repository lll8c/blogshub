package model

type Category struct {
	Id   int64  `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (*Category) TableName() string {
	return "category"
}

func InsertCategory(d *Category) error {
	return db.Create(&d).Error
}

func GetCategoryList() ([]*Category, error) {
	var list []*Category
	err := db.Find(&list).Error
	return list, err
}

func GetCategoryById(id int64) (*Category, error) {
	var category Category
	err := db.Where("id = ?", id).First(&category).Error
	return &category, err
}

func GetCategoryByName(name string) (*Category, error) {
	var category Category
	err := db.Where("name = ?", name).First(&category).Error
	return &category, err
}

func DeleteCategoryById(id int64) error {
	return db.Where("id = ?", id).Delete(&Category{}).Error
}

func BatchDeleteCategoryByIds(ids []int64) error {
	return db.Where("id in ?", ids).Delete(&Category{}).Error
}

func UpdateCategory(d *Category) error {
	return db.Save(&d).Error
}

func GetCategoryByPage(page int, pageSize int) (list []*Category, err error) {
	err = db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&list).Error
	return
}
