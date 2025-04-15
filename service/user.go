package service

import (
	"bloghub/domain"
	"bloghub/model"
	"errors"
)

func LoginUser(d *domain.Account) error {
	return nil
}

func AddUser(user *domain.User) error {
	//先判断用户是否存在，避免插表
	if user.Name == "" {
		return errors.New("用户名不能为空")
	}
	us, err := model.GetUserByName(user.Username)
	if err != nil {
		return err
	}
	if us != nil {
		return errors.New("用户名已存在")
	}
	//用户昵称没输入，设置默认昵称为用户名
	if user.Name == "" {
		user.Name = user.Username
	}
	u := &model.User{
		Username: user.Username,
		Name:     user.Name,
		Phone:    user.Phone,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Sex:      user.Sex,
		Info:     user.Info,
		Birth:    user.Birth,
	}
	//设置为默认用户角色
	u.Role = domain.USER
	if err := model.InsertUser(u); err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int64) error {
	return model.DeleteUserById(id)
}

func BatchDeleteUser(ids []int64) error {
	return model.BatchDeleteUserByIds(ids)
}

func UpdateUser(user *domain.User) error {
	u := &model.User{
		Username: user.Username,
		Name:     user.Name,
		Phone:    user.Phone,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Sex:      user.Sex,
		Info:     user.Info,
		Birth:    user.Birth,
	}
	return model.UpdateUserById(u)
}

func GetUser(id int64) (*domain.User, error) {
	u, err := model.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		Username: u.Username,
		Name:     u.Name,
		Phone:    u.Phone,
		Email:    u.Email,
		Avatar:   u.Avatar,
		Sex:      u.Sex,
		Info:     u.Info,
		Birth:    u.Birth,
	}, nil
}

func GetUserList() ([]*domain.User, error) {
	us, err := model.GetUserList()
	if err != nil {
		return nil, err
	}
	res := make([]*domain.User, 0)
	for _, u := range us {
		res = append(res, &domain.User{
			Username: u.Username,
			Name:     u.Name,
			Phone:    u.Phone,
			Email:    u.Email,
			Avatar:   u.Avatar,
			Sex:      u.Sex,
			Info:     u.Info,
			Birth:    u.Birth,
		})
	}
	return res, nil
}
