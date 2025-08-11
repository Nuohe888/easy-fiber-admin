package sqlite

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var db *gorm.DB
var cfg *Config

func Init(c *Config) {
	var err error
	cfg = c
	
	// 构建DSN
	dsn := cfg.Path
	if cfg.EnableWAL {
		dsn += "?_journal_mode=WAL"
	}
	if cfg.BusyTimeout > 0 {
		if cfg.EnableWAL {
			dsn += fmt.Sprintf("&_busy_timeout=%d", cfg.BusyTimeout)
		} else {
			dsn += fmt.Sprintf("?_busy_timeout=%d", cfg.BusyTimeout)
		}
	}

	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		panic(fmt.Sprintf("SQLite数据库连接失败: %s", err))
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("SQLite数据库获取失败: %s", err))
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 启用外键约束
	db.Exec("PRAGMA foreign_keys = ON;")
}

func Get() *gorm.DB {
	return db
}