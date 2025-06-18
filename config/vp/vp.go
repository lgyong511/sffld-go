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

type MgrViper struct {
	// 配置文件路径
	file *string
	// 配置信息
	conf *config.Config
	vp   *viper.Viper
}

func New() *MgrViper {
	defaultFile := "." + string(os.PathSeparator) + "data" + string(os.PathSeparator) + "config.yml"
	return &MgrViper{
		pflag.StringP("file", "c", defaultFile, "配置文件"),
		new(config.Config),
		viper.GetViper(),
	}
}

// 设置配置信息，命令行、环境变量、配置文件
func (m *MgrViper) Set() *MgrViper {
	logrus.Info("开始加载配置信息。。。")
	// 处理命令行参数
	pflag.IntP("app.port", "p", 2580, "app port")
	pflag.Parse()

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
	m.vp.SetDefault("log.level", "info")
	m.vp.SetDefault("log.format", "json")

	logrus.WithFields(logrus.Fields(m.vp.AllSettings())).Info("配置信息")
	return m
}

// 获取所有配置信息，命令行、环境变量、配置文件
func (m *MgrViper) Get() *config.Config {
	// 反序列化
	if err := m.vp.Unmarshal(m.conf); err != nil {
		logrus.WithError(err).Error("反序列化到结构体失败！")
	}

	return m.conf
}

// 启用配置文件修改监控
func (m *MgrViper) Relod() *MgrViper {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		m.Get()
		logrus.WithFields(logrus.Fields(m.vp.AllSettings())).Info("配置信息")
	})

	return m
}

// 保存配置到文件
func (m *MgrViper) Save() error {

	// 创建配置文件目录，已存在和创建成功返回nil
	if err := os.MkdirAll(filepath.Dir(*m.file), 0755); err != nil {
		logrus.WithError(err).Error("创建目录失败！")
		return err
	}
	// 保存前先把结构体重新设置到viper
	configMap := make(map[string]interface{})
	if err := mapstructure.Decode(m.conf, &configMap); err != nil {
		logrus.WithError(err).Error("结构体转成map失败！")
		return err
	}
	if err := viper.MergeConfigMap(configMap); err != nil {
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
	logrus.Info("配置保存成功！")
	return nil
}
