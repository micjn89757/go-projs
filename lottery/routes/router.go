package routes

import (
	"lottery/utils"
	"lottery/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() {
	gin.SetMode(utils.Conf.Server.AppMode)
	r := gin.New()

	// 上线要设置
	// err := r.SetTrustedProxies()

	// 设置静态文件目录
	r.Static("/img", "./view/img")	
	r.Static("/js", "./view/js")
	r.LoadHTMLFiles("view/lottery.html") // 应该是index.html


	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "lottery.html", nil)
	})


	front := r.Group("api/v1")
	{
		front.GET("/gifts", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errmsg.SUCCESS,
			"msg": "gifts",
		})

		front.GET("/lucky", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": errmsg.SUCCESS,
				"msg": "lucky",
			})
		})
	})
	}
	


	err := r.Run(utils.Conf.Server.HttpPort)
	if err != nil {
		panic("server start error")
	}
}