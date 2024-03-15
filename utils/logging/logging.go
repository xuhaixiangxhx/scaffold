// 自定义日志打印模块，封装logrus
package logging

import "github.com/sirupsen/logrus"

// 打印debug类型日志
func Debug(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Debug(msg)
}

// 打印info类型日志
func Info(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Info(msg)
}

// 打印warning类型日志
func Warning(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Warning(msg)
}

// 打印error类型日志
func Error(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Error(msg)
}
