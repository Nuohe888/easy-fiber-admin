package service

type ApiSrv struct {
}

var apiSrv *ApiSrv

func InitApiSrv() {
	apiSrv = &ApiSrv{}
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
