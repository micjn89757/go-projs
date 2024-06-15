package middleware

import "github.com/gin-gonic/gin"

type cookie struct {

}

func NewCookie() *cookie {
	return &cookie{}
}

func CookieAuthMiddleware(ctx *gin.Context) {
	cookie, err := ctx.Cookie("username") // 获取cookie

	if err != nil || cookie == ""{
		cookie = "NotSet"
		// 设置Cookie
		ctx.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		ctx.Abort()
	}

	ctx.Next()
}