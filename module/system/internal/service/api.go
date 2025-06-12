package service

import (
	"easy-fiber-admin/pkg/common/vo"
	"easy-fiber-admin/pkg/config"
	"easy-fiber-admin/pkg/logger"
	"easy-fiber-admin/plugin"
	"errors"
	"mime/multipart"
	"strings"
)

type ApiSrv struct {
	storage plugin.IStorage
	log     logger.ILog
	cfg     *config.Config
}

var apiSrv *ApiSrv

func InitApiSrv() {
	apiSrv = &ApiSrv{
		storage: plugin.GetStorage(),
		log:     logger.Get(),
		cfg:     config.Get(),
	}
}

func GetApiSrv() *ApiSrv {
	if apiSrv == nil {
		panic("service api init failed")
	}
	return apiSrv
}

//func (i *ApiSrv) UpdateFile(file *multipart.FileHeader) (string, error) {
//	//一般来说文件系统上传确定是固定的 在这里初始化是错误的
//	//比如是支付的话就没关系因为你会对接多个支付网关
//	//而且就算是支付的话也需要先从数据库加载一遍配置！！！
//	i.storage.Init("", "", "", "./upload/file", false)
//
//	path, key, err := i.storage.UploadFile(file)
//	if err != nil {
//		i.log.Errorf("上传文件失败: %v", err)
//		return "", errors.New("上传文件失败")
//	}
//
//	i.log.Infof("文件上传成功，key: %s", key)
//	return path, nil
//}

//func (i *ApiSrv) DelFile(key string) error {
//	err := i.storage.DeleteFile(key)
//	if err != nil {
//		i.log.Errorf("删除文件失败: %v", err)
//		return errors.New("删除文件失败")
//	}
//
//	return nil
//}

func (i *ApiSrv) FileUploadImg(file *multipart.FileHeader) (*vo.FileUrlRes, error) {
	path, key, err := i.storage.UploadFile(file)
	if err != nil {
		i.log.Errorf("上传文件失败: %v", err)
		return nil, errors.New("上传文件失败")
	}
	i.log.Infof("文件上传成功，key: %s", key)
	return &vo.FileUrlRes{
		FileUrl: i.formatImageUrl(path),
	}, nil
}

func (i *ApiSrv) formatImageUrl(url string) string {
	if url == "" {
		return ""
	}

	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	} else {
		return i.cfg.Server.Domain + "/" + url
	}
}
