package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
	"vote-gin/utils"
	"vote-gin/utils/msgcode"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func init() {
	sugar = utils.Logger.Sugar()
}

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}


// CreateToken 创建token
func (j *JWT) CreateToken(claim MyClaims)(string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	return token.SignedString(j.JwtKey)
}

// ParseToken 解析token 字符串
func (j *JWT) ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token)(any ,error) {
		return j.JwtKey, nil
	}, jwt.WithLeeway(5*time.Second))

	if err != nil {
		NewErr := fmt.Errorf("parse jwt string failed: %w", err)
		sugar.Error(NewErr.Error())
		return nil, NewErr
	} else if claim, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claim, nil
	} else {
		sugar.Error(jwt.ErrInvalidKeyType.Error())
		return nil, jwt.ErrInvalidType
	}
}

// JWTAuthMiddleware
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		var code int
		// token可能放在请求头，请求体或URI中
		// 主要考虑放在请求头中
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == ""{
			code = msgcode.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg": msgcode.GetErrMsg(code),
			})

			c.Abort()  // 防止调用该请求的其他处理程序
			return 
		}


		// 获取tokenString
		// 如果放在Authorization，格式为"Bear [tokenString]"
		checkToken := strings.Split(authHeader, " ")

		if !(len(checkToken) == 2 && checkToken[0] == "Bearer") {
			c.JSON(
				http.StatusOK,
				gin.H{
					"status": code,
					"msg": msgcode.GetErrMsg(code),
				},
			)
			c.Abort()
			return
		}
		
		j := NewJWT()
		// 解析token
		// checkToken[1]是获取到的Tokenstr
		token, err := j.ParseToken(checkToken[1])
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.JSON(http.StatusOK, gin.H{
					"status": msgcode.ERROR,
					"msg": "token expired, please login again",
					"data": nil,
				})
			}

			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status": msgcode.ERROR,
				"msg": err.Error(),
				"data": nil,
			})

			c.Abort()
			return 
		}

		// 将当前请求的username信息保存到请求上下文
		c.Set("username", token.Username)
		c.Next() // 进行后续处理
	}
}
