package main

import "gamebackend/network"

func main() {
	client := network.NewClient(":8023")
	client.Run()
	select {}
}
