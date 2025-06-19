package vp

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	mgr := New()
	assert.NotNil(t, mgr)
	assert.NotNil(t, mgr.vp)
	assert.NotNil(t, mgr.conf)
	assert.Equal(t, filepath.Join(".", "data", "config.yml"), *mgr.file)
}

func TestSetAndGet(t *testing.T) {
	// 创建临时测试目录和配置文件
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yml")
	err := os.WriteFile(configFile, []byte(`
app:
  port: 3000
  authTimeout: 5
log:
  level: "debug"
`), 0644)
	require.NoError(t, err)

	// 测试设置和获取配置
	mgr := New()
	mgr.file = &configFile // 使用测试配置文件

	conf := mgr.Get()

	assert.Equal(t, 3000, conf.App.Port)
	assert.Equal(t, 5, conf.App.AuthTimeout)
	assert.Equal(t, "debug", conf.Log.Level)
}

func TestSaveAndMergeConfigMap(t *testing.T) {
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "save_config.yml")

	// 第一次初始化
	mgr := New()
	mgr.file = &configFile // 直接设置，不通过flag

	// 修改配置
	conf := mgr.Get()
	conf.App.Port = 6000
	conf.App.AuthTimeout = 10

	// 保存配置
	err := mgr.Save()
	require.NoError(t, err)

	// 第二次初始化时不调用Set()，直接读取
	mgr2 := New()
	mgr2.file = &configFile
	// 不调用Set()，直接手动读取配置
	mgr2.vp.SetConfigFile(configFile)
	require.NoError(t, mgr2.vp.ReadInConfig())
	require.NoError(t, mgr2.vp.Unmarshal(mgr2.conf))

	conf2 := mgr2.Get()
	assert.Equal(t, 6000, conf2.App.Port)
	assert.Equal(t, 10, conf2.App.AuthTimeout)
}

func TestDefaultValues(t *testing.T) {
	mgr := New()

	conf := mgr.Get()
	assert.Equal(t, 2580, conf.App.Port)     // 默认值
	assert.Equal(t, 2, conf.App.AuthTimeout) // 默认值
	assert.Equal(t, "", conf.Log.Level)      // 未设置默认值
}

func TestEnvironmentVariables(t *testing.T) {
	t.Setenv("DDNS_PORT", "7000")
	t.Setenv("DDNS_AUTHTIMEOUT", "15")

	mgr := New()

	conf := mgr.Get()
	assert.Equal(t, 7000, conf.App.Port)
	assert.Equal(t, 15, conf.App.AuthTimeout)
}
