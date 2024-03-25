// 中间件层
package middlewares

import (
	"scaffold-demo/utils/jwtutil"
	"scaffold-demo/utils/logging"

	"github.com/gin-gonic/gin"
)

func JWTAuth(c *gin.Context) {
	// 1 login和logout接口无需验证token
	requestUrl := c.FullPath()
	logging.Debug(map[string]interface{}{"请求路径": requestUrl}, "")
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		logging.Debug(map[string]interface{}{"请求路径": requestUrl}, "登录注销无需验证token")
		c.Next()
		return
	}
	// 2 其他接口需要验证token
	tokenString := c.Request.Header.Get("Authorization")
	// 2.1 验证是否携带token
	if tokenString == "" {
		c.JSON(200, gin.H{
			"message": "请求未携带Token，请先登录",
			"status":  401,
		})
		c.Abort()
		return
	}
	// 2.2 验证token是否合法
	claims, err := jwtutil.ParseToken(tokenString)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Token验证未通过",
		})
		c.Abort()
		return
	}
	c.Set("claims", claims)
	c.Next()
}
