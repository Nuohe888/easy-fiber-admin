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

type DictDataSrv struct {
	db  *gorm.DB
	log logger.ILog
}

var dictDataSrv *DictDataSrv

func InitDictDataSrv() {
	dictDataSrv = &DictDataSrv{
		db:  sql.Get(),
		log: logger.Get(),
	}
}

func GetDictDataSrv() *DictDataSrv {
	if dictDataSrv == nil {
		panic("service dictData init failed")
	}
	return dictDataSrv
}

func (i *DictDataSrv) Add(dictData *system.DictData) error {
	return i.db.Create(&dictData).Error
}

func (i *DictDataSrv) Del(id any) error {
	return i.db.Where("id = ?", id).Delete(&system.DictData{}).Error
}

func (i *DictDataSrv) Put(id any, dictData *system.DictData) error {
	var _data system.DictData
	i.db.Where("id = ?", id).Find(&_data)
	if *_data.Id == 0 {
		return errors.New("不存在该Id")
	}
	utils.MergeStructs(&_data, dictData)
	return i.db.Save(&_data).Error
}

func (i *DictDataSrv) Get(id any) system.DictData {
	var dictData system.DictData
	i.db.Where("id = ?", id).Find(&dictData)
	return dictData
}

func (i *DictDataSrv) List(page, limit int, pid string) *vo.List {
	var items []system.DictData
	var total int64
	if limit == 0 {
		limit = 20
	}
	if pid == "" || pid == "0" {
		return &vo.List{}
	}
	db := i.db
	i.db.Where("pid=?", pid).Limit(limit).Offset((page - 1) * limit).Find(&items)
	db.Model(&system.DictData{}).Count(&total)
	return &vo.List{
		Items: items,
		Total: total,
	}
}
