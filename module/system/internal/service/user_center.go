package service

import (
	"easy-fiber-admin/model/system"
	"easy-fiber-admin/pkg/common/utils"
	"easy-fiber-admin/pkg/common/vo"
	"easy-fiber-admin/pkg/logger"
	"easy-fiber-admin/pkg/sql"
	"errors"
	"gorm.io/gorm"
)

type UserCenterSrv struct {
	db  *gorm.DB
	log logger.ILog
}

var userCenterSrv *UserCenterSrv

func InitUserCenterSrv() {
	userCenterSrv = &UserCenterSrv{
		db:  sql.Get(),
		log: logger.Get(),
	}
}

func GetUserCenterSrv() *UserCenterSrv {
	if userCenterSrv == nil {
		panic("service user_center init failed")
	}
	return userCenterSrv
}

func (i *UserCenterSrv) Add(userCenter *system.UserCenter) error {
	return i.db.Create(&userCenter).Error
}

func (i *UserCenterSrv) Del(id any) error {
	return i.db.Where("id = ?", id).Delete(&system.UserCenter{}).Error
}

func (i *UserCenterSrv) Put(id any, userCenter *system.UserCenter) error {
	var _userCenter system.UserCenter
	i.db.Where("id=?", id).Find(&_userCenter)
	if _userCenter.Id == nil || *_userCenter.Id == 0 {
		return errors.New("不存在该Id")
	}
	utils.MergeStructs(&_userCenter, userCenter)
	return i.db.Save(&_userCenter).Error
}

func (i *UserCenterSrv) Get(id any) system.UserCenter {
	var userCenter system.UserCenter
	i.db.Where("id = ?", id).Find(&userCenter)
	return userCenter
}

func (i *UserCenterSrv) List(page, limit int) *vo.List {
	var items []system.UserCenter
	var total int64
	if limit == 0 {
		limit = 20
	}
	db := i.db
	i.db.Limit(limit).Offset((page - 1) * limit).Find(&items)
	db.Model(&system.UserCenter{}).Count(&total)
	return &vo.List{
		Items: items,
		Total: total,
	}
}

func (i *UserCenterSrv) ListAll() []system.UserCenter {
	var items []system.UserCenter
	i.db.Find(&items)
	return items
}

func (i *UserCenterSrv) GetStatus() map[string]interface{} {
	return map[string]interface{}{
		"0": "禁用",
		"1": "启用",
	}
}
