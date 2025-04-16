package model

type ActivitySign struct {
	Id           int64  `gorm:"column:id" json:"id,omitempty"`
	ActivityId   int64  `gorm:"column:activity_id" json:"activity_id,omitempty"`
	UserId       int64  `gorm:"column:user_id" json:"user_id,omitempty"`
	Time         string `gorm:"column:time" json:"time,omitempty"`
	ActivityName string `gorm:"-" json:"activity_name,omitempty"`
	UserName     string `gorm:"-" json:"user_name,omitempty"`
}

func (*ActivitySign) TableName() string {
	return "activity_sign"
}

func InsertActivitySign(a *ActivitySign) error {
	return db.Create(&a).Error
}

func GetByActivityIdAndUserId(activityId int64, userId int64) (*ActivitySign, error) {
	var activitySign ActivitySign
	err := db.Where("activity_id = ? and user_id = ?", activityId, userId).First(&activitySign).Error
	return &activitySign, err
}

func GetActivitySignList(a *ActivitySign) (list []*ActivitySign, err error) {
	query := db.Model(&a).Select("ActivitySign.*, activity.name as ActivityName, user.name as UserName")
	query = db.Joins("left join activity on activity_sign.activity_id = activity.id")
	query = db.Joins("left join user on activity_sign.UserId = user.Id")
	err = query.Find(&list).Error
	return
}

func DeleteActivitySignById(id int64) error {
	return db.Where("id = ?", id).Delete(&ActivitySign{}).Error
}

func DeteleByActivityIdAndUserId(activityId int64, userId int64) error {
	return db.Where("activity_id = ? and user_id = ?", activityId, userId).Delete(&ActivitySign{}).Error
}
