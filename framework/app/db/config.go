package db

type Config struct {
	Source        string `mapstructure:"source"` // root:root@tcp(127.0.0.1:3306)/test
	DryRun        bool   `mapstructure:"dryRun"`
	SlowThreshold int    `mapstructure:"slowThreshold"`
	MaxLifeTime   int    `mapstructure:"maxLifeTime"`
	MaxOpenConn   int    `mapstructure:"maxOpenConn"`
	MaxIdleConn   int    `mapstructure:"maxIdleConn" `
}
