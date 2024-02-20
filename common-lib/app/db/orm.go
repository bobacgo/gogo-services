package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// sql.Open 方法只是校验参数，而不是连接数据（懒加载方式）
// 需要使用 sql.Ping 保证启动应用时和数据库有效的

func NewDB(dialector gorm.Dialector, conf Config) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, &gorm.Config{
		CreateBatchSize:                          1000, // 批量插入每次拆成 1k 条
		QueryFields:                              true, // 会根据当前model的所有字段名称进行 select
		PrepareStmt:                              true, // 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表命名
		},
		DryRun: conf.DryRun,
		Logger: Logger(conf),
	})
	if err != nil {
		return nil, fmt.Errorf("gorm open db err: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get DB err: %w", err)
	}
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(conf.MaxLifeTime))
	sqlDB.SetMaxOpenConns(conf.MaxOpenConn)
	sqlDB.SetMaxIdleConns(conf.MaxIdleConn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err = sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping db err: %w", err)
	}
	return db, nil
}

func Logger(conf Config) logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Duration(conf.SlowThreshold) * time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Silent,                                   // 日志级别
			IgnoreRecordNotFoundError: false,                                           // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                                           // 禁用彩色打印
		},
	)
}
