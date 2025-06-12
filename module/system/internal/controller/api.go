package controller

import (
	"easy-fiber-admin/module/system/internal/service"
	"easy-fiber-admin/pkg/common/vo"
	"errors"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type apiCtl struct {
	srv *service.ApiSrv
}

var ApiCtl *apiCtl

func InitApiCtl() {
	ApiCtl = &apiCtl{
		srv: service.GetApiSrv(),
	}
}

//func (i *apiCtl) UpdateFile(c *fiber.Ctx) error {
//	file, err := c.FormFile("file")
//	if err != nil {
//		return vo.ResultErr(errors.New("上传文件失败"), c)
//	}
//
//	if file.Size > 1024*1024 {
//		return vo.ResultErr(errors.New("文件大小超过1MB限制"), c)
//	}
//
//	contentType := file.Header.Get("Content-Type")
//	if !strings.HasPrefix(contentType, "image/") {
//		return vo.ResultErr(errors.New("只允许上传图片文件"), c)
//	}
//
//	_, err = i.srv.UpdateFile(file)
//	if err != nil {
//		return vo.ResultErr(err, c)
//	}
//	return vo.ResultOK(nil, c)
//}

//func (i *apiCtl) DelFile(c *fiber.Ctx) error {
//	key := c.Query("key")
//	err := i.srv.DelFile(key)
//	if err != nil {
//		return vo.ResultErr(err, c)
//	}
//	return vo.ResultOK(nil, c)
//}

func (i *apiCtl) FileUploadImg(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return vo.ResultErr(errors.New("上传文件失败"), c)
	}

	if file.Size > 3*1024*1024 {
		return vo.ResultErr(errors.New("文件大小超过3MB限制"), c)
	}

	contentType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		return vo.ResultErr(errors.New("只允许上传图片文件"), c)
	}

	res, err := i.srv.FileUploadImg(file)
	if err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(res, c)
}
