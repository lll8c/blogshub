package model

// 点赞或收藏的模块
const (
	BLOG     = "博客"
	ACTIVITY = "活动"
)

type Likes struct {
	Id     int64  `gorm:"column:id" json:"id,omitempty"`
	Fid    int64  `gorm:"column:fid" json:"fid,omitempty"`
	UserId int64  `gorm:"column:user_id" json:"user_id,omitempty"`
	Module string `gorm:"column:module" json:"module,omitempty"`
}

func (*Likes) TableName() string {
	return "likes"
}

func InsertLikes(a *Likes) error {
	return db.Create(&a).Error
}

func DeleteLikesById(id int64) error {
	return db.Where("id = ?", id).Delete(&Likes{}).Error
}

func GetUserLikes(likes *Likes) (like *Likes, err error) {
	err = db.Where("fid = ? and user_id = ? and module = ?", likes.Fid, likes.UserId, likes.Module).First(&like).Error
	return
}

// CountLikesByFidAndModule 查询点赞数量
func CountLikesByFidAndModule(fid int64, module string) (int64, error) {
	var count int64
	err := db.Model(&Likes{}).Where("fid = ? and module = ?", fid, module).Count(&count).Error
	return count, err
}
