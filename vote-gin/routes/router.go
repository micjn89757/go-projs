package routes

import (
	"net/http"
	"vote-gin/utils"
	"vote-gin/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger
var err error


func InitRouter() {
	sugar = utils.Logger.Sugar()
	defer sugar.Sync()
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		sugar.Infof("username:%s password:%s", username, password)
		code := model.QueryUserInfo(username, password)

		if code == 1000 {
			ctx.JSON(
				http.StatusBadGateway,
				gin.H{
					"msg": "login failed",
				},
			)
		}else {
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"msg": "login success",
				},
			)
		}
		
	})
	if err = r.Run(utils.HttpPort); err != nil {
		sugar.Errorf("%d端口启动失败: %w", utils.HttpPort, err)
		panic("http server port is used")
	}
}