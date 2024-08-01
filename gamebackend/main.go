package main

import (
	"gamebackend/aop/logger"
	"gamebackend/common"
	"gamebackend/world"
)

func main() {
	logger.InitLogger()
	defer logger.Logger.Sync()

	world.MM = world.NewMgrMgr()
	go world.MM.Run()
	common.WaitSignal(world.MM.OnSystemSignal)
}
