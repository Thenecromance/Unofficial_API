package basicImpl

import (
	"Unofficial_API/Interface"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	//logrus.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		log.SetOutput(f)
		// defer f.Close() // 在实际程序中在合适位置关闭文件
	} else {
		// 若打开文件失败，继续输出到默认 stderr
		log.Warn("failed to open log file, continue to stderr:", err)
	}

}

type logger struct {
}

func (l logger) Debug(args ...interface{}) {
	log.Debug(args...)
}

func (l logger) Info(args ...interface{}) {
	log.Info(args...)
}

func (l logger) Warn(args ...interface{}) {
	log.Warn(args...)
}

func (l logger) Error(args ...interface{}) {
	log.Error(args...)
}

func (l logger) Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func (l logger) Panic(args ...interface{}) {
	log.Panic(args...)
}

func NewLogger() Interface.Logger {
	return &logger{}
}
