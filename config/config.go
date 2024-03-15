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
	Port string
)

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
	logging.Debug(nil, "开始加载程序配置")
	// 环境变量方式加载程序配置
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("PORT", ":8090")
	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	initLogConfig(logLevel)
}
