package orm

import (
	"github.com/gogoclouds/gogo-services/common-lib/internal/dns/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlServer struct{}

func (mysqlServer) Open(conf *config.Configuration) (*gorm.DB, error) {
	source := conf.Database().Source
	return gorm.Open(mysql.Open(source), &gorm.Config{
		CreateBatchSize: 1000, // 批量插入每次拆成 1k 条
		QueryFields:     true, // 会根据当前model的所有字段名称进行 select
		PrepareStmt:     true, // 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
		Logger:          Server.Logger(),
	})
}