package model

import "gorm.io/gorm"

type Blog struct {
	Id           int64  `gorm:"column:id" json:"id"`
	Title        string `gorm:"column:title" json:"title"`
	Content      string `gorm:"column:content" json:"content"`
	Descr        string `gorm:"column:descr" json:"descr"`
	Cover        string `gorm:"column:cover" json:"cover"`
	Tags         string `gorm:"column:tags" json:"tags"`
	UserId       int64  `gorm:"column:user_id" json:"user_id"`
	Date         string `gorm:"column:date" json:"date,omitempty"`
	CategoryId   int    `gorm:"column:category_id" json:"category_id"`
	ReadCount    int    `gorm:"column:read_count" json:"read_count"`
	LikesCount   int    `gorm:"-" json:"likes_count"`
	CollectCount int    `gorm:"-" json:"collect_count"`
	CategoryName string `gorm:"-" json:"category_name"`
	UserName     string `gorm:"-" json:"user_name"`
	User         *User  `gorm:"-" json:"user"`
	//文章是否被当前用户点赞、收藏
	userLike    bool `gorm:"-" json:"user_like,omitempty"`
	userCollect bool `gorm:"-" json:"user_collect,omitempty"`
}

func (*Blog) TableName() string {
	return "blog"
}

func InsertBlog(b *Blog) error {
	return db.Create(&b).Error
}

func DeleteBlogById(id int64) error {
	return db.Where("id = ?", id).Delete(&Blog{}).Error
}

func BatchDeleteBlogByIds(ids []int64) error {
	return db.Where("id IN ?", ids).Delete(&Blog{}).Error
}

func UpdateBlogById(blog *Blog) error {
	return db.Model(&Blog{}).Where("id = ?", blog.Id).Updates(blog).Error
}

func GetBlogByID(id int64) (blog *Blog, err error) {
	query := db.Model(&Blog{})
	query = query.Select("blog.*, category.name as categoryName, user.name as userName")
	query = query.Joins("left join category on blog.category_id = category.id")
	query = query.Joins("left join user on blog.user_id = user.id")
	err = query.Where("id = ?", id).First(blog).Error
	return
}

func GetBlogList(blog *Blog) (list []*Blog, err error) {
	query := db.Model(&blog)
	query = query.Select("blog.*, category.name as categoryName, user.name as userName")
	query = query.Joins("left join category on blog.category_id = category.id")
	query = query.Joins("left join user on blog.user_id = user.id")
	if blog.Title != "" {
		query = query.Where("title like ?", "%"+blog.Title+"%")
	}
	if blog.CategoryName != "" {
		query = query.Where("category.name like ?", "%"+blog.CategoryName+"%")
	}
	if blog.UserName != "" {
		query = query.Where("user.name like ?", "%"+blog.UserName+"%")
	}
	err = db.Find(&list).Error
	return
}

func GetBlogByPage(blog *Blog, num int, size int) (list []*Blog, err error) {
	query := db.Model(&blog)
	query = query.Select("blog.*, category.name as categoryName, user.name as userName")
	query = query.Joins("left join category on blog.category_id = category.id")
	query = query.Joins("left join user on blog.user_id = user.id")
	if blog.Title != "" {
		query = query.Where("title like ?", "%"+blog.Title+"%")
	}
	if blog.CategoryName != "" {
		query = query.Where("category.name like ?", "%"+blog.CategoryName+"%")
	}
	if blog.UserName != "" {
		query = query.Where("user.name like ?", "%"+blog.UserName+"%")
	}
	err = query.Offset((num - 1) * size).Limit(size).Find(&list).Error
	return
}

func GetUserBlog(id int64) (list []*Blog, err error) {
	err = db.Where("user_id = ?", id).Find(&list).Error
	return
}

func UpdateBlogReadCount(blogId int64) error {
	return db.Model(&Blog{}).Where("id = ?", blogId).Update("read_count", gorm.Expr("read_count + ?", 1)).Error
}
