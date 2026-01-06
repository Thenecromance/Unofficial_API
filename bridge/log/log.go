package log

import (
	"Unofficial_API/Interface"
	"Unofficial_API/bridge/basicImpl"
)

var _log Interface.Logger

func init() {
	_log = basicImpl.NewLogger()
}

func Debug(args ...interface{}) {
	_log.Debug(args...)
}
func Info(args ...interface{}) {
	_log.Info(args...)
}
func Warn(args ...interface{}) {
	_log.Warn(args...)
}
func Error(args ...interface{}) {
	_log.Error(args...)
}
func Fatal(args ...interface{}) {
	_log.Fatal(args...)
}
func Panic(args ...interface{}) {
	_log.Panic(args...)
}
