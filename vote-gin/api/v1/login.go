package v1

import (
	"net/http"
	"time"
	"vote-gin/middleware"
	"vote-gin/model"
	"vote-gin/utils"
	"vote-gin/utils/msgcode"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

// 后台登录
func Login(ctx *gin.Context) {
	var user model.User
	var err error
	var code int 
	var tokenString string

	err = ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		sugar.Errorf("get params failed, %s", err.Error())
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": msgcode.ERROR,
			"msg": err.Error(),
		})
		ctx.Abort()
		return 
	}

	user, code = model.CheckLogin(user.Username, user.Password)
	if code == msgcode.SUCCESS {
		setToken(ctx, user)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": code,
			"data": user.Username,
			"id": user.ID,
			"msg": msgcode.GetErrMsg(code),
			"token": tokenString,
		})
	}
}

// 前台登录
func LoginFront(ctx *gin.Context) {
	var err error 
	var user model.User
	var code int 

	err = ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		sugar.Errorf("get params failed, %s", err.Error())
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": msgcode.ERROR,
			"msg": err.Error(),
		})
		ctx.Abort()
		return 
	}

	user, code = model.CheckLoginFront(user.Username, user.Password)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":	user.Username,
		"id": user.ID,
		"msg": msgcode.GetErrMsg(code),
	})


}

// setToken 生成Token
func setToken(ctx *gin.Context, user model.User) {
	j := middleware.NewJWT()
	myClaim := middleware.MyClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			Issuer: "vote",
		},
	}
	tokenString, err := j.CreateToken(myClaim)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": msgcode.ERROR,
			"msg": msgcode.GetErrMsg(msgcode.ERROR),
			"token": tokenString,
		})
	}
		
	ctx.JSON(http.StatusOK, gin.H{
		"status": msgcode.SUCCESS,
		"msg": msgcode.GetErrMsg(msgcode.SUCCESS),
		"data": user.Username,
		"id": user.ID,
		"token": tokenString,
	})
	return 
}

func init() {
	sugar = utils.Logger.Sugar()
}