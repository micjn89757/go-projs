package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() {
	gin.Mode()
	router := gin.New()

	router.GET("/register", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
			"status": 200,
		})
	})



	err := router.Run(":3000")

	if err != nil {
		panic(err)
	}

	fmt.Println("auth 服务启动成功")
}