package main

import (
	"mmorpg/network"
)

func main() {
	server := network.NewServer(":8023")
	server.Run()
	select{}
}