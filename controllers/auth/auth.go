package auth

import (
	"scaffold-demo/config"
	"scaffold-demo/utils/jwtutil"
	"scaffold-demo/utils/logging"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录的逻辑
func Login(c *gin.Context) {
	// 1 获取前端传过来的用户名和密码
	userInfo := UserInfo{}
	returnData := config.NewReturnData()
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		returnData.Status = 401
		returnData.Message = err.Error()
		c.JSON(200, returnData)
		return
	}
	logging.Debug(map[string]interface{}{"用户名": userInfo.Username, "密码": userInfo.Password}, "开始验证登录信息")
	// 2 验证用户名密码是否正确
	// 数据库 环境变量
	if userInfo.Username == "admin" && userInfo.Password == "123456" {
		// 认证成功，创建JWT Token
		ss, err := jwtutil.GenToken(userInfo.Username)
		// 创建JWT Token失败
		if err != nil {
			logging.Error(map[string]interface{}{"用户名": userInfo.Username, "错误信息": err.Error()}, "用户名密码正确但生成Token失败")
			returnData.Status = 401
			returnData.Message = "生成Token失败"
			c.JSON(200, returnData)
			return
		}
		// 创建JWT Token成功，返回给前端
		logging.Info(map[string]interface{}{"用户名": userInfo.Username}, "登录成功")
		returnData.Message = "登录成功"
		returnData.Data["token"] = ss
		c.JSON(200, returnData)
		return
	} else {
		// 用户名密码错误
		returnData.Status = 401
		returnData.Message = "用户名或密码错误"
		c.JSON(200, returnData)
	}
}

// 注销的逻辑
func Logout(c *gin.Context) {
	returnData := config.NewReturnData()
	returnData.Message = "注销成功"
	c.JSON(200, returnData)
	logging.Debug(nil, "用户已注销")
}
