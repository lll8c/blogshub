package main

import (
	"bloghub/model"
	"bloghub/web/router"
	"fmt"
)

func main() {
	r := router.SetupRouter()
	if err := model.InitDB(); err != nil {
		fmt.Println("初始化数据库失败")
		return
	}
	r.Run(":9091")
}
