package middleware 


import (
	"testing"
)


func TestJWT(t *testing.T) {
	myClaim := &MyClaims{
		Username: "admin",
	}
	jwt := NewJWT()

	// 创建jwt
	tokenString, err := jwt.CreateToken(*myClaim)
	if err != nil {
		t.Errorf("create token failed, err:%s", err.Error())
	}

	t.Log(tokenString)

	err = jwt.ParseToken(tokenString)
	if err != nil {
		t.Errorf("parse token failed, %s", err.Error())
	}
}