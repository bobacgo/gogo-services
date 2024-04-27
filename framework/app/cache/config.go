package cache

type RedisConf struct {
	Addrs        []string `mapstructure:"addrs"` // [127.0.0.1:6379, 127.0.0.1:7000]
	Username     string   `mapstructure:"username"`
	Password     string   `mapstructure:"password"`
	DB           uint8    `mapstructure:"db"`
	PoolSize     int      `mapstructure:"poolSize"`
	ReadTimeout  string   `mapstructure:"readTimeout"`  // 0.2s
	WriteTimeout string   `mapstructure:"writeTimeout"` // 0.2s
}
