package service

import (
	"bloghub/model"
	"bloghub/utils/ginx"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func LoginAdmin(a *model.Account) (*model.Admin, error) {
	admin, err := model.GetAdminByName(a.Username)
	if err != nil {
		//用户不存在
		if admin == nil {
			fmt.Println("用户不存在")
			return nil, ginx.UserNotExistErr
		}
		return nil, err
	}
	//比较密码
	if admin.Password != a.Password {
		return nil, ginx.UserAccountErr
	}
	//生成token
	admin.Token, err = ginx.CreateToken(admin.Id, admin.Role)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func AddAdmin(admin *model.Admin) error {
	//先判断用户是否存在，避免插表
	_, err := model.GetAdminByName(admin.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	//用户已存在
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ginx.UserExistErr
	}
	if admin.Password == "" {
		admin.Password = "123"
	}
	//us == nil
	//用户昵称没输入，设置默认昵称为用户名
	if admin.Name == "" {
		admin.Name = admin.Username
	}
	//设置为默认用户角色
	admin.Role = model.USER
	if err := model.InsertAdmin(admin); err != nil {
		return err
	}
	return nil
}

func DeleteAdmin(id int64) error {
	return model.DeleteAdminById(id)
}

func BatchDeleteAdmin(ids []int64) error {
	return model.BatchDeleteAdminByIds(ids)
}

func UpdateAdmin(admin *model.Admin) error {
	return model.UpdateAdminById(admin)
}

func GetAdmin(id int64) (*model.Admin, error) {
	admin, err := model.GetAdminByID(id)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func GetAdminList(a *model.Admin) ([]*model.Admin, error) {
	list, err := model.GetAdminList(a)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetAdminByPage(a *model.Admin, page int, size int) ([]*model.Admin, error) {
	list, err := model.GetAdminByPage(a, page, size)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func UpdateAdminPassword(account *model.Account) error {
	//先查用户是否存在
	admin, err := model.GetAdminByID(account.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ginx.UserNotExistErr
		}
		return err
	}
	if account.Password != admin.Password {
		return ginx.ParamPasswordErr
	}
	admin.Password = account.NewPassword
	err = model.UpdateAdminById(admin)
	return err
}
