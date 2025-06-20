package vp

import (
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/go-viper/mapstructure/v2"
	"github.com/lgyong511/sffld-go/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// ReloadCallback 配置变更回调函数
type ReloadCallback func(*config.Config)

// MgrViper viper管理
type MgrViper struct {
	// 配置文件路径
	file *string
	// 配置信息
	conf *config.Config
	vp   *viper.Viper
	// 配置变更回调函数列表
	callbacks []ReloadCallback
}

// New 创建一个MgrViper实例
func New() *MgrViper {
	// defaultFile := "." + string(os.PathSeparator) + "data" + string(os.PathSeparator) + "config.yml"
	defaultFile := filepath.Join(".", "data", "config.yml")
	m := &MgrViper{
		file: &defaultFile,
		conf: new(config.Config),
		vp:   viper.New(),
	}
	return m.init().reload()
}

// AddReloadCallback 添加配置变更回调函数
func (m *MgrViper) AddReloadCallback(callbacks ...ReloadCallback) {
	m.callbacks = append(m.callbacks, callbacks...)
}

// init 初始化MgrViper实例，设置配置信息，命令行、环境变量、配置文件
func (m *MgrViper) init() *MgrViper {

	logrus.Info("开始加载配置信息。。。")

	// 处理命令行参数
	if !pflag.Parsed() { //确保只执行一次
		m.file = pflag.StringP("file", "c", *m.file, "配置文件")
		pflag.IntP("app.port", "p", 2580, "app port")
		pflag.Parse()
	}

	// 处理环境变量
	m.vp.AutomaticEnv()
	m.vp.BindEnv("app.port", "DDNS_PORT")
	m.vp.BindEnv("app.authTimeout", "DDNS_AUTHTIMEOUT")

	// 创建 flag 集的副本，排除不需要绑定的 flag
	filteredFlags := pflag.NewFlagSet("filtered", pflag.ContinueOnError)
	pflag.VisitAll(func(flag *pflag.Flag) {
		if flag.Name != "file" { // 排除 file flag
			filteredFlags.AddFlag(flag)
		}
	})
	// 只绑定过滤后的 flag
	if err := m.vp.BindPFlags(filteredFlags); err != nil {
		logrus.WithError(err).Error("viper绑定环境变量失败！")
	}

	// 处理配置文件
	m.vp.SetConfigFile(*m.file)
	if err := m.vp.ReadInConfig(); err != nil {
		logrus.WithError(err).Error("读取配置文件失败！")
	}

	// 设置默认值
	m.vp.SetDefault("app.port", 2580)
	m.vp.SetDefault("app.authTimeout", 2)
	// m.vp.SetDefault("log.level", "info")
	// m.vp.SetDefault("log.format", "json")

	// 反序列化
	if err := m.vp.Unmarshal(m.conf); err != nil {
		logrus.WithError(err).Error("反序列化到结构体失败！")
	}

	logrus.SetLevel(logrus.DebugLevel)

	logrus.Info("加载配置信息完成。。。")

	return m
}

// Get 获取配置信息
func (m *MgrViper) Get() *config.Config {
	return m.conf
}

// reload 启用配置文件修改监控
func (m *MgrViper) reload() *MgrViper {
	// 监控配置文件变化
	m.vp.WatchConfig()
	// 配置文件变更回调函数
	m.vp.OnConfigChange(func(in fsnotify.Event) {
		logrus.Info("检测到配置文件变更，重新加载配置")
		// 反序列化
		if err := m.vp.Unmarshal(m.conf); err != nil {
			logrus.WithError(err).Error("反序列化到结构体失败！")
		}
		// 调用回调函数
		for _, callback := range m.callbacks {
			callback(m.conf)
		}
		logrus.WithFields(m.conf.ToLogFields()).Info("重新加载配置完成。。。")
		logrus.Info("重新加载配置完成。。。")

	})

	return m
}

// Save 保存配置到文件
func (m *MgrViper) Save() error {

	// 创建配置文件目录，已存在和创建成功返回nil
	if err := os.MkdirAll(filepath.Dir(*m.file), 0755); err != nil {
		logrus.WithError(err).Error("创建目录失败！")
		return err
	}
	if err := m.mergeConfigMap(); err != nil {
		logrus.WithError(err).Error("map绑定到viper失败！")
		return err
	}

	// 写入配置文件
	err := m.vp.WriteConfig() // 如果文件不存在会报错
	if err != nil {
		logrus.WithError(err).Error("配置文件不存储，尝试创建中！")
		// 如果文件不存在，使用 SafeWriteConfig
		err = m.vp.SafeWriteConfig()
		if err != nil {
			logrus.WithError(err).Error("创建配置文件失败！")
			return err
		}
	}
	logrus.Info("配置文件保存成功！")
	return nil
}

// mergeConfigMap 把配置信息结构体重新合并到vp中
func (m *MgrViper) mergeConfigMap() error {
	// 把结构体转成map
	configMap := make(map[string]interface{})
	if err := mapstructure.Decode(m.conf, &configMap); err != nil {
		logrus.WithError(err).Error("结构体转成map失败！")
		return err
	}
	// 保存前先把结构体重新设置到viper
	if err := m.vp.MergeConfigMap(configMap); err != nil {
		logrus.WithError(err).Error("map绑定到viper失败！")
		return err
	}
	return nil
}
