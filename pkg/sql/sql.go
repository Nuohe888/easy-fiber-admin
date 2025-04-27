package sql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var cfg *Config

func Init(c *Config) {
	var err error
	cfg = c
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DbName)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败: %s", err))
	}
	s, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("数据库获取失败: %s", err))
	}
	s.SetMaxIdleConns(cfg.MaxIdleConns)
	s.SetMaxOpenConns(cfg.MaxOpenConns)
}

func Get() *gorm.DB {
	return db
}
