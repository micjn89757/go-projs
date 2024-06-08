package middleware

import (
	"fmt"
	"time"
	"vote-gin/utils"

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
func (j *JWT) ParseToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token)(any ,error) {
		return j.JwtKey, nil
	}, jwt.WithLeeway(5*time.Second))

	if err != nil {
		NewErr := fmt.Errorf("parse jwt string failed: %w", err)
		sugar.Error(NewErr.Error())
		return NewErr
	} else if _, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return nil
	} else {
		sugar.Error(jwt.ErrInvalidKeyType.Error())
		return jwt.ErrInvalidType
	}
}
