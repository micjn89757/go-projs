package routes

import (
	"net/http"
	"vote-gin/api/v1"
	"vote-gin/middleware"
	"vote-gin/model"
	"vote-gin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger
var err error


func InitRouter() {
	sugar = utils.Logger.Sugar()
	defer sugar.Sync()
	defer model.Close()  // 关闭数据库连接
	gin.SetMode(utils.AppMode)
	// 强制日志颜色化
	gin.ForceConsoleColor()

	r := gin.New()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
	// 设置信任网络
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)
	
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	// 后台管理路由接口
	// auth := r.Group("api/v1")
	// auth.Use(middleware.JWTAuthMiddleware())
	// {
	// 	auth.POST("/login", v1.Login)
	// }

	// 展示接口
	router := r.Group("api/v1")
	{
		// 显示
		router.GET("", func(ctx *gin.Context){
			ctx.HTML(http.StatusOK, "login.html"),
		})
		// 登录控制
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)
		// router.POST("login", func(ctx *gin.Context) {
		// 	username,_ := ctx.Get("username")
		// 	username, _ = username.(string)
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"msg": "login success",
		// 		"data": username,
		// 	})
		// })
	}
	if err = r.Run(utils.HttpPort); err != nil {
		sugar.Errorf("%d端口启动失败: %w", utils.HttpPort, err)
		panic("http server port is used")
	}
}