package main

import (
	"vote-gin/model"
	"vote-gin/routes"
)

func main() {
	model.InitDB()
	routes.InitRouter()
}