package controller

import (
	"easy-fiber-admin/model/system"
	"easy-fiber-admin/module/system/internal/service"
	"easy-fiber-admin/module/system/internal/vo"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type dictDataCtl struct {
	srv *service.DictDataSrv
}

var DictDataCtl *dictDataCtl

func InitDictDataCtl() {
	DictDataCtl = &dictDataCtl{
		srv: service.GetDictDataSrv(),
	}
}

func (i *dictDataCtl) Add(c *fiber.Ctx) error {
	var req system.DictData
	if err := vo.BodyParser(&req, c); err != nil {
		return err
	}
	if err := i.srv.Add(&req); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *dictDataCtl) Del(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := i.srv.Del(id); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *dictDataCtl) Put(c *fiber.Ctx) error {
	id := c.Params("id")
	var req system.DictData
	if err := vo.BodyParser(&req, c); err != nil {
		return err
	}
	if err := i.srv.Put(id, &req); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *dictDataCtl) Get(c *fiber.Ctx) error {
	id := c.Query("id")
	return vo.ResultOK(i.srv.Get(id), c)
}

func (i *dictDataCtl) List(c *fiber.Ctx) error {
	page := c.Query("page")
	limit := c.Query("limit")
	pid := c.Query("pid")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	return vo.ResultOK(i.srv.List(pageInt, limitInt, pid), c)
}
