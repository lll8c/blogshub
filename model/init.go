package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB 初始化MySQL连接
func InitDB() (err error) {
	//fmt.Println(cfg.Dsn)
	db, err = gorm.Open(mysql.Open("root:123@tcp(localhost:3306)/bloghub?parseTime=True&loc=Local"))
	if err != nil {
		return err
	}
	return nil
}
