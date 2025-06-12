package model

import (
	"easy-fiber-admin/model/system"
	"easy-fiber-admin/pkg/sql"
	"easy-fiber-admin/plugin"
)

func Init() {
	if err := sql.Get().AutoMigrate(
		system.User{},
		system.Role{},
		system.UserCenter{},
	); err != nil {
		panic("初始化数据库结构失败: " + err.Error())
	}

	//下面代码是初始化数据
	var count int64

	password, err := plugin.GetCrypto().EncryptPassword("123456")
	if err != nil {
		panic("密码生成失败: " + err.Error())
	}

	// 检查用户表是否有数据
	sql.Get().Model(&system.User{}).Count(&count)
	if count == 0 {
		var (
			username = "admin"
			nickname = ""
			status   = 1
			roleName = "管理员"
			roleCode = "admin"
		)

		// 先创建角色
		role := system.Role{
			Name:   &roleName,
			Code:   &roleCode,
			Status: &status,
		}
		if err := sql.Get().Create(&role).Error; err != nil {
			panic("创建角色失败: " + err.Error())
		}

		// 再创建用户，使用角色的实际ID
		user := system.User{
			Username: &username,
			Password: &password,
			Nickname: &nickname,
			RoleId:   role.Id, // 使用创建后的角色ID
			Status:   &status,
		}
		if err := sql.Get().Create(&user).Error; err != nil {
			panic("创建用户失败: " + err.Error())
		}
	}
}
