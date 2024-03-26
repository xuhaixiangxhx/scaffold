// 程序的入口
package main

import (
	"scaffold-demo/config"
	"scaffold-demo/middlewares"
	"scaffold-demo/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1 加载程序配置
	// 2 配置启动gin
	r := gin.Default()
	// 全局中间件拦截请求验证合法性
	r.Use(middlewares.JWTAuth)

	routers.RegisterRouters(r)

	r.Run(config.Port)
}
