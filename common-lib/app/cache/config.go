package cache

type RedisConf struct {
	Addrs        []string `yaml:"addrs"` // [127.0.0.1:6379, 127.0.0.1:7000]
	Username     string   `yaml:"username"`
	Password     string   `yaml:"password"`
	DB           uint8    `yaml:"db"`
	PoolSize     int      `yaml:"poolSize"`
	ReadTimeout  string   `yaml:"readTimeout"`  // 0.2s
	WriteTimeout string   `yaml:"writeTimeout"` // 0.2s
}
