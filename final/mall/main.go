package main

import (
	"mall/conf"
	"mall/model"
	"mall/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	model.ListenOrder()
	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
