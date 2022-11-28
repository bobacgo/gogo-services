package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type Configuration struct {
	mu sync.RWMutex
	c  *config
}

func New() *Configuration {
	return &Configuration{mu: sync.RWMutex{}, c: &config{}}
}

// Sync 局部更新
func (c *Configuration) Sync(data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return yaml.Unmarshal(data, c.c)
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

func (c *Configuration) Log() Log {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c.Log
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

func (c *Configuration) Config() config {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return *c.c
}

var Bootstrap = &bootstrap{}

type bootstrap struct {
	ConfigMD *FileMetadata `yaml:"_config"`
}

// Unmarshal 解析启动文件配置
func (s *bootstrap) Unmarshal(configFilePath string) *FileMetadata {
	buff, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(buff, s); err != nil {
		panic(err)
	}
	return s.ConfigMD
}