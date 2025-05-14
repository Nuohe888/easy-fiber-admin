package service

import (
	"easy-fiber-admin/model/system"
	"easy-fiber-admin/module/system/internal/utils"
	"easy-fiber-admin/module/system/internal/vo"
	"easy-fiber-admin/pkg/logger"
	"easy-fiber-admin/pkg/sql"
	"errors"
	"gorm.io/gorm"
)

type RoleSrv struct {
	db  *gorm.DB
	log logger.ILog
}

var roleSrv *RoleSrv

func InitRoleSrv() {
	roleSrv = &RoleSrv{
		db:  sql.Get(),
		log: logger.Get(),
	}
}

func GetRoleSrv() *RoleSrv {
	if roleSrv == nil {
		panic("service role init failed")
	}
	return roleSrv
}

func (i *RoleSrv) Add(role *system.Role) error {
	return i.db.Create(&role).Error
}

func (i *RoleSrv) Del(id any) error {
	return i.db.Where("id = ?", id).Delete(&system.Role{}).Error
}

func (i *RoleSrv) Put(idStr any, role *system.Role) error {
	var _role system.Role
	i.db.Where("id = ?", idStr).Find(&_role)
	if *_role.Id == 0 {
		return errors.New("不存在该Id")
	}

	utils.MergeStructs(&_role, role)

	return i.db.Save(&_role).Error
}

func (i *RoleSrv) Get(id any) system.Role {
	var role system.Role
	i.db.Where("id = ?", id).Find(&role)
	return role
}

func (i *RoleSrv) List(page, pageSize int) *vo.List {
	var items []system.Role
	var total int64
	if pageSize == 0 {
		pageSize = 20
	}
	db := i.db
	i.db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&items)
	db.Model(&system.Role{}).Count(&total)
	return &vo.List{
		Items: items,
		Total: total,
	}
}

func (i *RoleSrv) ListAll() []system.Role {
	var items []system.Role
	i.db.Find(&items)
	return items
}
