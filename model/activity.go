package model

import "gorm.io/gorm"

type Activity struct {
	Id      int64  `gorm:"column:id" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Descr   string `gorm:"column:descr" json:"descr"`
	Content string `gorm:"column:content" json:"content"`
	Start   string `gorm:"column:start" json:"start"`
	End     string `gorm:"column:end" json:"end"`
	//活动形式
	Form      string `gorm:"column:form" json:"form"`
	Address   string `gorm:"column:address" json:"address"`
	Host      string `gorm:"column:host" json:"host"`
	ReadCount int64  `gorm:"column:read_count" json:"read_count"`
	Cover     string `gorm:"column:cover" json:"cover"`

	IsEnd        bool  `gorm:"-" json:"is_end"`
	IsSign       bool  `gorm:"-" json:"is_sign"`
	LikesCount   int   `gorm:"-" json:"likes_count"`
	CollectCount int   `gorm:"-" json:"collect_count"`
	IsLike       bool  `gorm:"-" json:"is_like"`
	IsCollect    bool  `gorm:"-" json:"is_collect"`
	UserId       int64 `gorm:"-" json:"user_id"`
}

func (*Activity) TableName() string {
	return "activity"
}

func InsertActivity(a *Activity) error {
	err := db.Create(&a).Error
	return err
}

func DeleteActivityById(id int64) error {
	return db.Where("id = ?", id).Delete(&Activity{}).Error
}

func BatchDeleteActivityByIds(ids []int64) error {
	return db.Where("id in ?", ids).Delete(&Activity{}).Error
}

func UpdateActivity(a *Activity) error {
	return db.Save(&a).Error
}

func GetActivityById(id int64) (*Activity, error) {
	var category Activity
	err := db.Where("id = ?", id).First(&category).Error
	return &category, err
}

func GetActivityList(a *Activity) (list []*Activity, err error) {
	query := db.Model(&a)
	if a.Id != 0 {
		query = query.Where("id = ?", a.Id)
	}
	if a.Name != "" {
		query = query.Where("name like", "%"+a.Name+"%")
	}
	if a.Descr != "" {
		query = query.Where("name like", "%"+a.Descr+"%")
	}
	if a.Start != "" {
		query = query.Where("name like", "%"+a.Start+"%")
	}
	if a.End != "" {
		query = query.Where("name like", "%"+a.End+"%")
	}
	if a.Form != "" {
		query = query.Where("name like", "%"+a.Form+"%")
	}
	if a.Address != "" {
		query = query.Where("name like", "%"+a.Address+"%")
	}
	if a.Host != "" {
		query = query.Where("name like", "%"+a.Host+"%")
	}
	err = query.Find(&list).Error
	return
}

func GetActivityByPage(a *Activity, page int, pageSize int) (list []*Activity, err error) {
	query := db.Model(&a)
	if a.Id != 0 {
		query = query.Where("id = ?", a.Id)
	}
	if a.Name != "" {
		query = query.Where("name like", "%"+a.Name+"%")
	}
	if a.Descr != "" {
		query = query.Where("name like", "%"+a.Descr+"%")
	}
	if a.Start != "" {
		query = query.Where("name like", "%"+a.Start+"%")
	}
	if a.End != "" {
		query = query.Where("name like", "%"+a.End+"%")
	}
	if a.Form != "" {
		query = query.Where("name like", "%"+a.Form+"%")
	}
	if a.Address != "" {
		query = query.Where("name like", "%"+a.Address+"%")
	}
	if a.Host != "" {
		query = query.Where("name like", "%"+a.Host+"%")
	}
	err = query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&list).Error
	return
}

func UpdateActivityReadCount(activityId int64) error {
	return db.Model(&Activity{}).Where("id = ?", activityId).Update("read_count", gorm.Expr("read_count + ?", 1)).Error
}
