// 自定义日志打印模块，封装logrus
package logging

import (
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

func addCallerInfo(fields map[string]interface{}) map[string]interface{} {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return fields
	}
	if fields == nil {
		fields = make(map[string]interface{})
	}
	funcName := runtime.FuncForPC(pc).Name()
	fields["file"] = file + ":" + strconv.Itoa(line)
	fields["func"] = path.Base(funcName)
	return fields
}

// 打印debug类型日志
func Debug(fields map[string]interface{}, msg string) {
	fields = addCallerInfo(fields)
	logrus.WithFields(fields).Debug(msg)
}

// 打印info类型日志
func Info(fields map[string]interface{}, msg string) {
	fields = addCallerInfo(fields)
	logrus.WithFields(fields).Info(msg)
}

// 打印warning类型日志
func Warning(fields map[string]interface{}, msg string) {
	fields = addCallerInfo(fields)
	logrus.WithFields(fields).Warning(msg)
}

// 打印error类型日志
func Error(fields map[string]interface{}, msg string) {
	fields = addCallerInfo(fields)
	logrus.WithFields(fields).Error(msg)
}
