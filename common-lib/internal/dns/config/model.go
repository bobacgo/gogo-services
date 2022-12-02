package config

type FileMetadata struct {
	Namespace string
	Group     string
	Filenames []string
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

func (a App) IsEnabledRPC() bool {
	return a.Server.Rpc.Addr != ""
}

func (a App) IsEnabledHttp() bool {
	return a.Server.Http.Addr != ""
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
	Addr         []string // [127.0.0.1:6379, 127.0.0.1:7000]
	Password     string
	ReadTimeout  string `yaml:"readTimeout"`  // 0.2s
	WriteTimeout string `yaml:"writeTimeout"` // 0.2s
}
