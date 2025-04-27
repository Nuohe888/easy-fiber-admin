package model

import (
	"go-server/model/system"
	"go-server/pkg/sql"
)

func Init() {
	if err := sql.Get().AutoMigrate(
		system.User{},
	); err != nil {
		panic("初始化数据库结构失败: " + err.Error())
	}
}
