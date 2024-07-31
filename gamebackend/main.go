package main

import (
	"gamebackend/common"
	"gamebackend/world"
)

func main() {
	common.InitLogger()
	defer common.Logger.Sync()

	world.MM = world.NewMgrMgr()
	go world.MM.Run()
	common.WaitSignal(world.MM.OnSystemSignal)
}
