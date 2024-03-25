// 程序的入口
package main

import (
	"fmt"
	"scaffold-demo/config"
	"scaffold-demo/utils/jwtutil"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1 加载程序配置
	// 2 配置启动gin
	r := gin.Default()
	// 验证生成token的方法
	ss, _ := jwtutil.GenToken("apple")
	fmt.Println(ss)
	// 验证解析token的方法
	claims, err := jwtutil.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYXBwbGUiLCJpc3MiOiJKV1QiLCJzdWIiOiJ4dWhhaXhpYW5nIiwiZXhwIjoxNzExMzQwMjkyLCJuYmYiOjE3MTEzMzMwOTIsImlhdCI6MTcxMTMzMzA5Mn0.F6pQruZc7efS_qyfcu-u4Cx5QZ6oUBDs8NlL00VD87U")
	if err != nil {
		// 解析token失败
		fmt.Println("解析token失败:", err.Error())
	} else {
		// 解析token成功
		fmt.Println(claims)
	}
	r.Run(config.Port)
}
