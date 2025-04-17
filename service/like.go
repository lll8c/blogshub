package service

import (
	"bloghub/model"
	"errors"
	"gorm.io/gorm"
)

// SetLikes 投票或取消投票
func SetLikes(like *model.Likes) error {
	//先查询数据是否存在
	_, err := model.GetUserLikes(like)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	//数据存在
	if err == nil {
		err := model.DeleteLikesById(like.Id)
		if err != nil {
			return err
		}
	} else {
		err := model.InsertLikes(like)
		if err != nil {
			return err
		}
	}
	return nil
}

// IsUserLike 查询当前用户是否点赞过
func IsUserLike(fid int64, module string, userId int64) (bool, error) {
	like := &model.Likes{
		Fid:    fid,
		UserId: userId,
		Module: module,
	}
	_, err := model.GetUserLikes(like)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	if err == nil {
		return true, nil
	}
	return false, nil
}
