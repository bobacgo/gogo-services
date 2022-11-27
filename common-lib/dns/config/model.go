package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"sync"
)

type Configuration struct {
	mu sync.RWMutex
	c  *config
}

func New() Configuration {
	return Configuration{mu: sync.RWMutex{}, c: &config{}}
}

// Sync 局部更新
func (c *Configuration) Sync(data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if err := yaml.Unmarshal(data, c.c); err != nil {
		log.Printf("sync config: %v", err)
	}
}

func (c *Configuration) App() App {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.App
}

func (c *Configuration) AppServiceKV() map[string]any {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.App.ServiceKV
}

func (c *Configuration) Log() App {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.App
}

func (c *Configuration) Database() Database {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.Data.Database
}

func (c *Configuration) Redis() Redis {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.Data.Redis
}

func (c Configuration) String() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return fmt.Sprintf("%+v\n", c.c)
}

// --------------------------------------

type config struct {
	App  App
	Log  Log
	Data struct {
		Database Database
		Redis    Redis
	}
}

// Transport 传输协议
type Transport struct {
	Addr    string // 0.0.0.0:8000
	Timeout string // 1s
}

// --------------- app -----------------

// App 应用服务相关配置信息
type App struct {
	Name    string // 服务名
	Version string // 版本号
	Server  struct {
		Http Transport
		Rpc  Transport
	}
	ServiceKV map[string]any `yaml:"serviceKV"` // 业务自定义kv
}

// --------------- log ----------------

// Log config model
//
// 日志文件名 xxx/logs/${App.Service}-2006-01-01-150405.log
type Log struct {
	Level       string // 日志级别 默认值是 info
	FileSizeMax uint16 `yaml:"fileSizeMax"`             // 单位是MB 默认值是 10MB
	FileAgeMax  uint16 `yaml:"fileAgeMax"`              // 留存天数
	DirPath     string `validator:"dir" yaml:"dirPath"` // 日志文件夹路径 默认 ./logs
}

// ------------- data ----------------

type Database struct {
	Driver string // mysql
	Source string // root:root@tcp(127.0.0.1:3306)/test
}

type Redis struct {
	Addr         string // 127.0.0.1:6379
	ReadTimeout  string `yaml:"readTimeout"`  // 0.2s
	WriteTimeout string `yaml:"writeTimeout"` // 0.2s
}
