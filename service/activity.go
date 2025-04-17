package service

import (
	"bloghub/model"
	"bloghub/utils/ginx"
	"errors"
	"gorm.io/gorm"
	"sort"
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

// GetActivity 活动详情
func GetActivity(id int64, account *model.Account) (*model.Activity, error) {
	act, err := model.GetActivityById(id)
	if err != nil {
		return nil, err
	}
	//设置活动额外信息
	err = SetAct(act, account)
	if err != nil {
		return nil, err
	}
	//设置当前博客的点赞、收藏数
	likeCount, err := model.CountLikesByFidAndModule(act.Id, model.ACTIVITY)
	if err != nil {
		return nil, err
	}
	collectCount, err := model.CountCollectByFidAndModule(act.Id, model.ACTIVITY)
	if err != nil {
		return nil, err
	}
	act.LikesCount = likeCount
	act.CollectCount = collectCount
	//当前用户是否点赞收藏
	flag, err := IsUserLike(id, model.ACTIVITY, account.Id)
	if err != nil {
		return nil, err
	}
	act.IsLike = flag
	flag, err = IsUserCollect(id, model.ACTIVITY, account.Id)
	if err != nil {
		return nil, err
	}
	act.IsCollect = flag
	return act, nil
}

func GetActivities(query *model.Activity) ([]*model.Activity, error) {
	return model.GetActivityList(query)
}

func GetActivityByPage(query *model.Activity, pageNum, pageSize int, account *model.Account) ([]*model.Activity, error) {
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
func SetAct(act *model.Activity, account *model.Account) error {
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

// GetUserSignActivity 分页查询用户报名的活动
func GetUserSignActivity(a *model.Activity, page int, size int) ([]*model.Activity, error) {
	return model.GetUserSignActivity(a, page, size)
}

// GetUserLikeActivity 分页查询用户点赞的活动
func GetUserLikeActivity(a *model.Activity, page int, size int) ([]*model.Activity, error) {
	return model.GetUserLikeActivity(a, page, size)
}

// GetUserCollectActivity 分页查询用户收藏的活动
func GetUserCollectActivity(a *model.Activity, page int, size int) ([]*model.Activity, error) {
	return model.GetUserCollectActivity(a, page, size)
}

func GetUserCommentActivity(a *model.Activity, page int, size int) ([]*model.Activity, error) {
	return model.GetUserCommentActivity(a, page, size)
}

// GetTopActivities todo 热榜算法
func GetTopActivities() ([]*model.Activity, error) {
	//先查询所有活动
	list, err := GetActivities(&model.Activity{})
	if err != nil {
		return nil, err
	}
	//使用sort 按 readCount 进行排序
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].ReadCount > list[j].ReadCount
	})
	return list, err
}
