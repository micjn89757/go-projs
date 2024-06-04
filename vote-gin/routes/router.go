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
	r := gin.Default()
	
	r.POST("/login", v1.Login)
	if err = r.Run(utils.HttpPort); err != nil {
		sugar.Errorf("%d端口启动失败: %w", utils.HttpPort, err)
		panic("http server port is used")
	}
}