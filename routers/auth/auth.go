package auth

import (
	"scaffold-demo/controllers/auth"

	"github.com/gin-gonic/gin"
)

// 登录接口
func Login(authgroup *gin.RouterGroup) {
	authgroup.POST("/login", auth.Login)
}

// 注销接口
func Logout(authgroup *gin.RouterGroup) {
	authgroup.GET("/logout", auth.Logout)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置auth功能的路由策略
	authGroup := g.Group("/auth")
	Login(authGroup)
	Logout(authGroup)
}
