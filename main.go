// 程序的入口
package main

import (
	"scaffold-demo/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1 加载程序配置
	// 2 配置启动gin
	r := gin.Default()
	r.Run(config.Port)
}
