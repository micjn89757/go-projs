package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/demo", func(ctx *gin.Context) {
		ctx.JSON(
			200,
			gin.H{
				"message": "hello world!",
			},
		)
	})

	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}