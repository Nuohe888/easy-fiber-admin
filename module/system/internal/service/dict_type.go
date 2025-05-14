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

type DictTypeSrv struct {
	db  *gorm.DB
	log logger.ILog
}

var dictTypeSrv *DictTypeSrv

func InitDictTypeSrv() {
	dictTypeSrv = &DictTypeSrv{
		db:  sql.Get(),
		log: logger.Get(),
	}
}

func GetDictTypeSrv() *DictTypeSrv {
	if dictTypeSrv == nil {
		panic("service dictType init failed")
	}
	return dictTypeSrv
}

func (i *DictTypeSrv) Add(dictType *system.DictType) error {
	return i.db.Create(&dictType).Error
}

func (i *DictTypeSrv) Del(id any) error {
	return i.db.Where("id = ?", id).Delete(&system.DictType{}).Error
}

func (i *DictTypeSrv) Put(id any, dictType *system.DictType) error {
	var _dictType system.DictType
	i.db.Where("id = ?", id).Find(&_dictType)
	if *_dictType.Id == 0 {
		return errors.New("不存在该Id")
	}

	utils.MergeStructs(&_dictType, dictType)

	return i.db.Save(&_dictType).Error
}

func (i *DictTypeSrv) Get(id any) system.DictType {
	var dictType system.DictType
	i.db.Where("id = ?", id).Find(&dictType)
	return dictType
}

func (i *DictTypeSrv) List(page, limit int) *vo.List {
	var items []system.DictType
	var total int64
	if limit == 0 {
		limit = 20
	}
	db := i.db
	i.db.Limit(limit).Offset((page - 1) * limit).Find(&items)
	db.Model(&system.DictType{}).Count(&total)
	return &vo.List{
		Items: items,
		Total: total,
	}
}

func (i *DictTypeSrv) Dict() []vo.Dict {
	var types []system.DictType
	var datas []system.DictData

	i.db.Find(&types)
	i.db.Find(&datas)

	var list []vo.Dict

	for _, v := range types {
		var item vo.Dict
		item.Type = *v.Type
		for _, v2 := range datas {
			item.List = append(item.List, v2)
		}
		list = append(list, item)
	}
	return list
}
