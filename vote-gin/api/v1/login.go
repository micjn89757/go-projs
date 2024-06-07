package v1

import (
	"net/http"
	"vote-gin/model"
	"vote-gin/utils"
	"vote-gin/utils/msgcode"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func Login(ctx *gin.Context) {
	var user model.User
	var err error
	var code int 

	err = ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		sugar.Errorf("get params failed, %s", err.Error())
		ctx.JSON(http.StatusBadGateway, gin.H{
			"msg": "error",
		})
		ctx.Abort()
		return 
	}

	user, code = model.CheckLogin(user.Username, user.Password)
	if code == msgcode.SUCCESS {
		// TODO 设置token
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "login success",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": code,
			"data": user.Username,
			"id": user.ID,
			"msg": msgcode.GetErrMsg(code),
			"token": "token",
		})
	}


}



func init() {
	sugar = utils.Logger.Sugar()
}