package lg

import (
	"io"
	"os"

	"github.com/lgyong511/logcuting"
	"github.com/lgyong511/sffld-go/config"
	"github.com/sirupsen/logrus"
)

// 日志配置

// setLogurs 配置日志
func SetLogurs(logConf *config.Log) {
	logrus.SetReportCaller(logConf.Caller)
	logrus.SetLevel(getLevel(logConf.Level))
	logrus.SetOutput(getWriter(logConf.Output, logConf.MultiOut))
	// 设置日志输出格式
	if logConf.JSONFormat {
		logrus.SetFormatter(&logrus.JSONFormatter{
			// 日志时间格式
			// TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			// 日志时间格式
			// TimestampFormat: "2006-01-02 15:04:05",
		})
	}
}

// getWriter 获取日志输出
// "data/log-%Y%m%d%H%M.log"
func getWriter(path string, multiOut bool) io.Writer {

	if len(path) == 0 {
		return os.Stdout
	}
	// 创建logcuting实例
	logcut := logcuting.NewLogcuting(&logcuting.Config{
		Name: path,
		// Time: time.Minute,
		// Size: 1,
	})
	if multiOut {
		return io.MultiWriter(logcut, os.Stdout)
	}

	return logcut

}

// getLevel 获取日志级别
func getLevel(level string) logrus.Level {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return logrus.InfoLevel
	}
	return l
}
