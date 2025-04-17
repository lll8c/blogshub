package model

type Comment struct {
	Id        int64      `gorm:"column:id" json:"id,omitempty"`
	Content   string     `gorm:"column:content" json:"content,omitempty"`
	UserId    int64      `gorm:"column:user_id" json:"user_id,omitempty"`
	Pid       int64      `gorm:"column:pid" json:"pid,omitempty"`
	RootId    int64      `gorm:"column:root_id" json:"root_id,omitempty"`
	Time      string     `gorm:"column:time" json:"time,omitempty"`
	Fid       int64      `gorm:"column:fid" json:"fid,omitempty"`
	Module    string     `gorm:"module" json:"module,omitempty"`
	UserName  string     `gorm:"-" json:"username,omitempty"`
	Avatar    string     `gorm:"-" json:"avatar,omitempty"`
	replyUser string     `gorm:"-" json:"reply_user,omitempty"`
	Children  []*Comment `gorm:"-" json:"children,omitempty"`
}

func (*Comment) TableName() string {
	return "comment"
}

func InsertComment(c *Comment) error {
	return db.Create(&c).Error
}

func DeleteCommentById(id int64) error {
	return db.Where("id = ?", id).Delete(&Comment{}).Error
}

func BatchDeleteComment(ids []int64) error {
	return db.Where("id IN ?", ids).Delete(&Comment{}).Error
}

func UpdateCommentById(c *Comment) error {
	return db.Save(&c).Error
}

func GetCommentById(id int64) (*Comment, error) {
	var comment Comment
	err := db.Where("id = ?", id).First(&comment).Error
	return &comment, err
}

// GetCommentList 难点
// 需要获取到发表评论的用户名以及回复的用户名
func GetCommentList(c *Comment) ([]*Comment, error) {
	var comments []*Comment
	query := db.Model(&c).Select("comment.*, user.name as userName, user.avatar as avatar, u2.name as replyUser")
	query = db.Joins("left join user on comment.user_id = user.Id")
	query = db.Joins("left join comment c2 on c2.id = comment.pid")
	query = db.Joins("left join user u2 on c2.user_id = u2.Id")
	if c.UserName != "" {
		query = db.Where("user.name = ?", "%"+c.UserName+"%")
	}
	if c.Fid != 0 {
		query = query.Where("comment.fid = ?", c.Fid)
	}
	if c.Module != "" {
		query = query.Where("comment.module = ?", c.Module)
	}
	if c.RootId != 0 {
		query = query.Where("comment.root_id = ?", c.RootId)
	}
	err := query.Find(&comments).Error
	return comments, err
}

// GetRootCommentList 获取所有根评论
func GetRootCommentList(c *Comment) ([]*Comment, error) {
	var comments []*Comment
	query := db.Model(&c).Select("comment.*, user.name as userName, user.avatar as avatar")
	query = db.Joins("left join user on comment.user_id = user.Id")
	if c.UserName != "" {
		query = db.Where("user.name = ?", "%"+c.UserName+"%")
	}
	if c.Fid != 0 {
		query = query.Where("comment.fid = ?", c.Fid)
	}
	if c.Module != "" {
		query = query.Where("comment.module = ?", c.Module)
	}
	err := query.Where("pid = 0").Find(&comments).Error
	return comments, err
}

func GetCommentListByPage(c *Comment, num int, size int) ([]*Comment, error) {
	var comments []*Comment
	query := db.Model(&c).Select("comment.*, user.name as userName, user.avatar as avatar, u2.name as replyUser")
	query = db.Joins("left join user on comment.user_id = user.Id")
	query = db.Joins("left join comment c2 on c2.id = comment.pid")
	query = db.Joins("left join user u2 on c2.user_id = u2.Id")
	if c.UserName != "" {
		query = db.Where("user.name = ?", "%"+c.UserName+"%")
	}
	if c.Fid != 0 {
		query = query.Where("comment.fid = ?", c.Fid)
	}
	if c.Module != "" {
		query = query.Where("comment.module = ?", c.Module)
	}
	if c.RootId != 0 {
		query = query.Where("comment.root_id = ?", c.RootId)
	}
	err := query.Offset((num - 1) * size).Limit(size).Find(&comments).Error
	return comments, err
}

// CountCommentByFid 统计文章或活动评论数量
func CountCommentByFid(fid int64, module string) (int64, error) {
	var count int64 = 0
	err := db.Where("fid = ? and module = ?", fid, module).Count(&count).Error
	return count, err
}
