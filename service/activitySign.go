package service

import (
	"bloghub/model"
)

func AddActivitySign(a *model.ActivitySign) error {
	return model.InsertActivitySign(a)
}

func DeleteActivitySign(id int64) error {
	return model.DeleteActivitySignById(id)
}

func BatchDeleteActivitySign(ids []int64) error {
	return model.BatchDeleteActivitySign(ids)
}

// DeleteUserActivitySign 删除用户报名的活动
func DeleteUserActivitySign(activityId int64, userId int64) error {
	return model.DeleteUserActivitySign(activityId, userId)
}

func GetActivitySignByPage(a *model.ActivitySign, page int, size int) ([]*model.ActivitySign, error) {
	return model.GetActivitySignByPage(a, page, size)
}
