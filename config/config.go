package config

// 配置信息结构体

// Config 配置信息
type Config struct {
	App *App `yaml:"app"`
	Log *Log `yaml:"log"`
}

// App 应用程序
type App struct {
	// gin监听端口号
	Port int `yaml:"port"`
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
	Output string `yaml:"output"`
	// 是否同时输出到控制台和文件
	MultiOut bool
	// 是否记录调用者
	Caller bool `yaml:"caller"`
}
