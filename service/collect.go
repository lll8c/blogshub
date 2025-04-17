package service

import (
	"bloghub/model"
	"errors"
	"gorm.io/gorm"
)

// SetCollect 收藏或取消收藏
func SetCollect(collect *model.Collect) error {
	//先查询数据是否存在
	_, err := model.GetUserCollect(collect)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	//数据存在
	if err == nil {
		err := model.DeleteCollectById(collect.Id)
		if err != nil {
			return err
		}
	} else {
		err := model.InsertCollect(collect)
		if err != nil {
			return err
		}
	}
	return nil
}

// IsUserCollect 查询当前用户是否点赞过
func IsUserCollect(fid int64, module string, userId int64) (bool, error) {
	like := &model.Collect{
		Fid:    fid,
		UserId: userId,
		Module: module,
	}
	_, err := model.GetUserCollect(like)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	if err == nil {
		return true, nil
	}
	return false, nil
}
