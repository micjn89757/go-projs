package main

import (
	"lottery/model"
	"lottery/routes"
)

func main() {
	// 初始化数据库
	model.InitDB()
	routes.InitRoute()
}