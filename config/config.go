// 配置层 存放程序的配置信息
package config

import (
	"scaffold-demo/utils/logging"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	// 时间格式化配置
	TimeFormat string = "2006-01-02 15:04:05"
)

var (
	Port          string
	JwtSignKey    string
	JwtExpireTime int64 //jwt token过期时间，单位：分钟
	Username      string
	Password      string
)

// 规范化定义返回格式
type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 使用构造函数初始化返回数据
func NewReturnData() ReturnData {
	returnData := ReturnData{}
	returnData.Status = 200
	data := make(map[string]interface{})
	returnData.Data = data
	return returnData
}

func initLogConfig(logLevel string) {
	// 日志输出级别
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	// 日志输出格式，添加文件名和行号字段
	logrus.SetReportCaller(true)
	// 日志格式调整为json
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: TimeFormat})
}

func init() {
	// 环境变量方式加载程序配置
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("PORT", ":8090")
	viper.SetDefault("JWT_SIGN_KEY", "likely")
	viper.SetDefault("JWT_EXPIRE_TIME", 120)
	viper.SetDefault("USERNAME", "admin")
	viper.SetDefault("PASSWORD", "123456")
	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL")
	// 加载日志输出格式
	initLogConfig(logLevel)
	logging.Debug(nil, "开始加载程序配置")
	Port = viper.GetString("PORT")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpireTime = viper.GetInt64("JWT_EXPIRE_TIME")
	Username = viper.GetString("USERNAME")
	Password = viper.GetString("PASSWORD")
}
