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
	r.GET("/login", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusAccepted,
			gin.H{
				"message": "hello world!",
			},
		)
	})

	r.POST("/login", func(ctx *gin.Context) {
		var user model.User
		err = ctx.ShouldBind(&user)
		if err != nil {
			
		}
	})
	if err := r.Run(utils.HttpPort); err != nil {
		sugar.Errorf("%d端口启动失败: %w", utils.HttpPort, err)
	}
}