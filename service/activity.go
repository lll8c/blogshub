package service

import (
	"bloghub/domain"
	"bloghub/model"
	"bloghub/utils/ginx"
	"errors"
	"gorm.io/gorm"
	"time"
)

func AddActivity(a *model.Activity) error {
	return model.InsertActivity(a)
}

func DeleteActivity(id int64) error {
	return model.DeleteActivityById(id)
}

func BatchDeleteActivity(ids []int64) error {
	return model.BatchDeleteActivityByIds(ids)
}

func UpdateActivity(a *model.Activity) error {
	return model.UpdateActivity(a)
}

// GetActivity todo 查询活动详情
func GetActivity(id int64) (*model.Activity, error) {
	return model.GetActivityById(id)
}

func GetActivities(query *model.Activity) ([]*model.Activity, error) {
	return model.GetActivityList(query)
}

func GetActivityByPage(query *model.Activity, pageNum, pageSize int, account *domain.Account) ([]*model.Activity, error) {
	list, err := model.GetActivityByPage(query, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	//设置额外信息
	for _, activity := range list {
		err := SetAct(activity, account)
		if err != nil {
			return nil, err
		}
	}
	return list, nil
}

// SetAct 获取的活动后为其设置额外信息
func SetAct(act *model.Activity, account *domain.Account) error {
	//是否已经结束
	layout := "2006-01-02"
	endTime, err := time.Parse(layout, act.End)
	if err != nil {
		return ginx.SystemErr
	}
	if endTime.Before(time.Now()) {
		act.IsEnd = true
	}
	//当前用户是否报名
	_, err = model.GetByActivityIdAndUserId(act.Id, account.Id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	//找到了
	if err == nil {
		act.IsSign = true
	}
	return nil
}

func ReadActivityCount(id int64) error {
	return model.UpdateActivityReadCount(id)
}
