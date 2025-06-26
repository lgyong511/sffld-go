package config

import "github.com/sirupsen/logrus"

// 配置信息结构体

// Config 配置信息
type Config struct {
	App *App `yaml:"app"`
	Log *Log `yaml:"log"`
}

// App 主程序配置
type App struct {
	// gin监听端口号
	Port string `yaml:"port"`
	// 鉴权超时时间，单位小时
	AuthTimeout int `yaml:"authTimeout"`
}

// Log 日志
type Log struct {
	// 日志级别
	Level string `yaml:"level"`
	//启用json日志输出格式
	JSONFormat bool `yaml:"JSONFormat"`
	// 日志输出
	// 空值将日志输出到控制台
	// 支持日期格式："data/log-%Y%m%d%H%M.log"
	Output string `yaml:"output"`
	// 是否同时输出到控制台和文件
	// Output有值时有效
	MultiOut bool `yaml:"multiout"`
	// 是否记录调用者
	Caller bool `yaml:"caller"`
}

// ToConfigFields 把配置信息转成logrus.Fields
func (c *Config) ToConfigFields() logrus.Fields {
	return logrus.Fields{
		"app": map[string]interface{}{
			"port":        c.App.Port,
			"authTimeout": c.App.AuthTimeout,
		},
		"log": map[string]interface{}{
			"level":      c.Log.Level,
			"jsonFormat": c.Log.JSONFormat,
			"output":     c.Log.Output,
			"multiout":   c.Log.MultiOut,
			"caller":     c.Log.Caller,
		},
	}
}
