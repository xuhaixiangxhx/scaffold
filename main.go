// 程序的入口
package main

import (
	// 下划线方式导入含义：只执行包中的init函数不引用包中的其他内容
	_ "scaffold-demo/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1 加载程序配置
	// 2 配置启动gin
	r := gin.Default()
	r.Run()
}
