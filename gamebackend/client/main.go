package main

import "gamebackend/common"

func main() {
	c := NewClient()
	c.InputHandlerRegister()
	c.Run()

	common.WaitSignal()
}
