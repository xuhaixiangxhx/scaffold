package jwtutil

import (
	"errors"
	"scaffold-demo/config"
	"scaffold-demo/utils/logging"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSignKey = []byte(config.JwtSignKey)

// 1.自定义申明类型
type MyCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

// 2.封装生成token的函数
func GenToken(name string) (string, error) {
	claims := MyCustomClaims{
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.JwtExpireTime))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "JWT",
			Subject:   "xuhaixiang",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSignKey)
	return ss, err
}

// 3.封装解析token的函数
func ParseToken(ss string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})
	if err != nil {
		logging.Error(nil, "解析Token失败")
		return nil, err
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		// Token合法
		return claims, nil
	} else {
		// Token不合法
		logging.Error(nil, "Token不合法")
		return nil, errors.New("token不合法：invalid token")
	}

}
