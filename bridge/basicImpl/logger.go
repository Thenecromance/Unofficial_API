package basicImpl

import (
	"Unofficial_API/Interface"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true, // 建议开启完整时间戳，方便查看日志
	})
	//logrus.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)
	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		// 使用 io.MultiWriter 将输出同时重定向到 标准输出(控制台) 和 文件
		mw := io.MultiWriter(os.Stdout, f)
		log.SetOutput(mw)
		// defer f.Close() // 在实际程序中在合适位置关闭文件
	} else {
		// 若打开文件失败，继续输出到默认 stderr
		log.Warn("failed to open log file, continue to stderr:", err)
		log.SetOutput(os.Stdout)
	}

}

// NewLogger 返回 logrus 的标准 logger 实例。
// 由于 logrus.Logger 的方法签名与 Interface.Logger 一致，因此可以直接返回，无需额外的 Wrapper 结构体。
func NewLogger() Interface.Logger {
	return log.StandardLogger()
}
