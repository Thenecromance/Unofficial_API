package log

import (
	"Unofficial_API/Interface"
	"Unofficial_API/bridge/basicImpl"
)

// 利用 Go 的函数变量特性，直接导出对应的日志方法。
// 这样既避免了编写大量重复的 Wrapper 函数，也让实现更加通透。
var (
	Debug  func(args ...interface{})
	Debugf func(format string, args ...interface{})
	Info   func(args ...interface{})
	Infof  func(format string, args ...interface{})
	Warn   func(args ...interface{})
	Warnf  func(format string, args ...interface{})
	Error  func(args ...interface{})
	Errorf func(format string, args ...interface{})
	Fatal  func(args ...interface{})
	Fatalf func(format string, args ...interface{})
	Panic  func(args ...interface{})
	Panicf func(format string, args ...interface{})
)

func init() {
	// 默认使用 basicImpl (logrus) 初始化
	SetLogger(basicImpl.NewLogger())
}

// SetLogger 允许在运行时（例如 main 函数启动时）替换底层的 Logger 实现。
// 通过方法值绑定 (Method Values) 将接口实现的方法直接赋值给包级变量。
func SetLogger(l Interface.Logger) {
	Debug = l.Debug
	Debugf = l.Debugf
	Info = l.Info
	Infof = l.Infof
	Warn = l.Warn
	Warnf = l.Warnf
	Error = l.Error
	Errorf = l.Errorf
	Fatal = l.Fatal
	Fatalf = l.Fatalf
	Panic = l.Panic
	Panicf = l.Panicf
}
