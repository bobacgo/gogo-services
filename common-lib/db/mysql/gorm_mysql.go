package mysql

import (
	"github.com/gogoclouds/gogo-services/common-lib/db/orm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Open() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("orm.db"), &gorm.Config{
		CreateBatchSize: 1000, // 批量插入每次拆成 1k 条
		QueryFields:     true, // 会根据当前model的所有字段名称进行 select
		PrepareStmt:     true, // 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
		Logger:          orm.Gorm.Logger(),
	})
	if err != nil {
		log.Panicln(err)
	}
	return db
}
