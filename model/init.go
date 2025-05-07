package model

import (
	"easy-fiber-admin/model/system"
	"easy-fiber-admin/pkg/sql"
)

func Init() {
	if err := sql.Get().AutoMigrate(
		system.User{},
		system.Role{},
		system.DictType{},
		system.DictData{},
	); err != nil {
		panic("初始化数据库结构失败: " + err.Error())
	}
}
