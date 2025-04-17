package model

type Collect struct {
	Id     int64  `gorm:"column:id" json:"id,omitempty"`
	Fid    int64  `gorm:"column:fid" json:"fid,omitempty"`
	UserId int64  `gorm:"column:user_id" json:"user_id,omitempty"`
	Module string `gorm:"column:module" json:"module,omitempty"`
}

func (*Collect) TableName() string {
	return "collect"
}

func InsertCollect(c *Collect) error {
	return db.Create(&c).Error
}

func DeleteCollectById(id int64) error {
	return db.Where("id = ?", id).Delete(&Collect{}).Error
}

// GetUserCollect 查询用户收藏数据
func GetUserCollect(collect *Collect) (c *Collect, err error) {
	err = db.Where("fid = ? and user_id = ? and module = ?", collect.Fid, collect.UserId, collect.Module).First(&c).Error
	return
}

func CountCollectByFidAndModule(fid int64, module string) (int64, error) {
	var count int64
	err := db.Model(&Collect{}).Where("fid = ? and module = ?", fid, module).Count(&count).Error
	return count, err
}
