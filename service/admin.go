package service

import (
	"bloghub/domain"
	"bloghub/model"
	"bloghub/utils/ginx"
	"fmt"
)

func LoginAdmin(a *domain.Account) (*domain.Admin, error) {
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
	countAdmin := &domain.Admin{
		Id:       admin.Id,
		Username: admin.Username,
		Password: admin.Password,
		Name:     admin.Name,
		Avatar:   admin.Avatar,
		Role:     admin.Role,
		Phone:    admin.Phone,
		Email:    admin.Email,
	}
	//生成token
	countAdmin.Token, err = ginx.CreateToken(admin.Id, admin.Role)
	if err != nil {
		return nil, err
	}
	return countAdmin, nil
}
