package main

import (
	"gamebackend/network"
)

func main() {
	server := network.NewServer(":8023")
	server.Run()
	select {}
}
