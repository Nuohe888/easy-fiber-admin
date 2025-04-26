package service

import (
	"go-server/pkg/sql"
	"gorm.io/gorm"
)

type ApiSrv struct {
	db *gorm.DB
}

var apiSrv *ApiSrv

func InitApiSrv() {
	apiSrv = &ApiSrv{
		db: sql.Get(),
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
