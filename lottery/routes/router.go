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
	r.Static("/static", "./view/static")	
	r.LoadHTMLFiles("/front", "./view/**/*.html") // 应该是index.html


	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "front", nil)
	})


	/*
		后台管理接口
	*/
	// auth := r.Group("api/v1")


	/*
		前端展示页面接口
	*/

	front := r.Group("api/v1")
	front.POST("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errmsg.SUCCESS,
			"msg": "ok",
		})
	})


	err := r.Run(utils.Conf.Server.HttpPort)
	if err != nil {
		panic("server start error")
	}
}