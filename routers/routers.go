// 路由层 管理程序的路由信息
package routers

import (
	"scaffold-demo/routers/auth"

	"github.com/gin-gonic/gin"
)

// 注册路由的方法
func RegisterRouters(r *gin.Engine) {
	// 路由分组
	apiGroup := r.Group("/api")
	/*
		登录的路由配置
		/api/auth
		/api/auth/login
		/api/auth/logout
	*/
	auth.RegisterSubRouter(apiGroup)
}
