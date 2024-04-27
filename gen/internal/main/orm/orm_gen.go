package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gogoclouds/gogo-services/framework/pkg/word"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	l := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)

	var err error
	var mysqlDSN = "root:root@tcp(localhost:3306)/mall-ums?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{
		CreateBatchSize:                          1000, // 批量插入每次拆成 1k 条
		QueryFields:                              true, // 会根据当前model的所有字段名称进行 select
		PrepareStmt:                              true, // 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表命名
		},
		DryRun: false,
		Logger: l,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	//defer sqlDB.Close()
	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}

	l.Info(context.Background(), "db stats %+v", sqlDB.Stats())
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "../main-service/internal/query",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
	})
	g.WithJSONTagNameStrategy(func(columnName string) (tagContent string) {
		return word.UderscoreToLowerCamelCase(columnName)
	})
	g.UseDB(db)
	g.ApplyBasic(
		g.GenerateModel("admin"),
	)
	g.Execute()
}
