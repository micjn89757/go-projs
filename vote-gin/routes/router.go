package routes

import (
	"vote-gin/api/v1"
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
	r := gin.New()
	// 设置信任网络
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)
	
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	
	// 后台管理路由接口
	// auth := r.Group("api/v1")
	// auth.Use(middleware.JWTAuthMiddleware())
	// {
	// 	auth.POST("/login", v1.Login)
	// }

	// 前端接口
	router := r.Group("api/v1")
	{
		// 登录控制
		router.POST("login", v1.Login)
	}
	if err = r.Run(utils.HttpPort); err != nil {
		sugar.Errorf("%d端口启动失败: %w", utils.HttpPort, err)
		panic("http server port is used")
	}
}