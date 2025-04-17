package service

import (
	"bloghub/model"
	"bloghub/utils/ginx"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func LoginUser(a *model.Account) (*model.User, error) {
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
	user.Token, err = ginx.CreateToken(user.Id, user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func RegisterUser(a *model.Account) error {
	user := &model.User{}
	copier.Copy(user, a)
	return AddUser(user)
}

func AddUser(user *model.User) error {
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
	//设置为默认用户角色
	user.Role = model.USER
	if err := model.InsertUser(user); err != nil {
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

func UpdateUser(user *model.User) error {
	return model.UpdateUserById(user)
}

func GetUser(id int64) (*model.User, error) {
	u, err := model.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserList(user *model.User) ([]*model.User, error) {
	list, err := model.GetUserList(user)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetUserByPage(user *model.User, page int, size int) ([]*model.User, error) {
	list, err := model.GetUserByPage(user, page, size)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func UpdateUserPassword(account *model.Account) error {
	//先查用户是否存在
	user, err := model.GetUserById(account.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ginx.UserNotExistErr
		}
		return err
	}
	if account.Password != user.Password {
		return ginx.ParamPasswordErr
	}
	user.Password = account.NewPassword
	err = model.UpdateUserById(user)
	return err
}
