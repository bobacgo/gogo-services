package db

type Config struct {
	Source        string `yaml:"source"` // root:root@tcp(127.0.0.1:3306)/test
	DryRun        bool   `yaml:"dryRun"`
	SlowThreshold int    `yaml:"slowThreshold"`
	MaxLifeTime   int    `yaml:"maxLifeTime"`
	MaxOpenConn   int    `yaml:"maxOpenConn"`
	MaxIdleConn   int    `yaml:"maxIdleConn"`
}
