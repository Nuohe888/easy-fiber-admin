package service

import (
	"easy-fiber-admin/pkg/logger"
	"easy-fiber-admin/pkg/sql"
	"gorm.io/gorm"
)

type ApiSrv struct {
	db  *gorm.DB
	log logger.ILog
}

var apiSrv *ApiSrv

func InitApiSrv() {
	apiSrv = &ApiSrv{
		db:  sql.Get(),
		log: logger.Get(),
	}
}

func GetApiSrv() *ApiSrv {
	if apiSrv == nil {
		panic("service api init failed")
	}
	return apiSrv
}

func (i *ApiSrv) Ping() error {
	return nil
}
