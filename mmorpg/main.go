package main

import "mmorpg/world"

func main() {
	world.MM = world.NewMgrMgr()
	world.MM.Pm.Run()
}