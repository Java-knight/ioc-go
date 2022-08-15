package logger

import (
	"github.com/fatih/color"
)

var disableLogs = false // 禁用日志

// 日志颜色 （蓝色）
func Blue(format string, a ...interface{}) {
	if disableLogs {
		return
	}
	color.Blue(format, a...)
}

// 红色
func Red(format string, a ...interface{}) {
	if disableLogs {
		return
	}
	color.Red(format, a...)
}

// 青色
func Cyan(format string, a ...interface{}) {
	if disableLogs {
		return
	}
	color.Cyan(format, a...)
}

// 开启日志
func Disable() {
	disableLogs = true
}
