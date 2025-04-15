package service

import (
	"bloghub/domain"
	"bloghub/model"
	"bloghub/utils/ginx"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func LoginUser(a *domain.Account) (*domain.User, error) {
	user, err := model.GetUserByName(a.Username)
	if err != nil {
		//用户不存在
		if user == nil {
			fmt.Println("用户不存在")
			return nil, ginx.UserNotExistErr
		}
		return nil, err
	}
	//比较密码
	if user.Password != a.Password {
		return nil, ginx.UserAccountErr
	}
	accountUser := &domain.User{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Name:     user.Name,
		Avatar:   user.Avatar,
		Role:     user.Role,
		Sex:      user.Sex,
		Phone:    user.Phone,
		Email:    user.Email,
		Info:     user.Info,
		Birth:    user.Birth,
	}
	accountUser.Token, err = ginx.CreateToken(user.Id, user.Role)
	if err != nil {
		return nil, err
	}
	return accountUser, nil
}

func RegisterUser(a *domain.Account) error {
	user := &domain.User{}
	copier.Copy(user, a)
	return AddUser(user)
}

func AddUser(user *domain.User) error {
	//先判断用户是否存在，避免插表
	if user.Username == "" {
		return errors.New("用户名不能为空")
	}
	_, err := model.GetUserByName(user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	//用户已存在
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ginx.UserExistErr
	}
	//us == nil
	//用户昵称没输入，设置默认昵称为用户名
	if user.Name == "" {
		user.Name = user.Username
	}
	u := &model.User{}
	copier.Copy(u, user)
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

			Name:   u.Name,
			Phone:  u.Phone,
			Email:  u.Email,
			Avatar: u.Avatar,
			Sex:    u.Sex,
			Info:   u.Info,
			Birth:  u.Birth,
		})
	}
	return res, nil
}
