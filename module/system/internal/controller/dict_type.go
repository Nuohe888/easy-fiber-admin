package controller

import (
	"easy-fiber-admin/model/system"
	"easy-fiber-admin/module/system/internal/service"
	"easy-fiber-admin/module/system/internal/vo"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type dictTypeCtl struct {
	srv *service.DictTypeSrv
}

var DictTypeCtl *dictTypeCtl

func InitDictTypeCtl() {
	DictTypeCtl = &dictTypeCtl{
		srv: service.GetDictTypeSrv(),
	}
}

func (i *dictTypeCtl) Add(c *fiber.Ctx) error {
	var req system.DictType
	if err := vo.BodyParser(&req, c); err != nil {
		return err
	}
	if err := i.srv.Add(&req); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *dictTypeCtl) Del(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := i.srv.Del(id); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *dictTypeCtl) Put(c *fiber.Ctx) error {
	id := c.Params("id")
	var req system.DictType
	if err := vo.BodyParser(&req, c); err != nil {
		return err
	}
	if err := i.srv.Put(id, &req); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *dictTypeCtl) Get(c *fiber.Ctx) error {
	id := c.Query("id")
	return vo.ResultOK(i.srv.Get(id), c)
}

func (i *dictTypeCtl) List(c *fiber.Ctx) error {
	page := c.Query("page")
	limit := c.Query("limit")
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	return vo.ResultOK(i.srv.List(pageInt, limitInt), c)
}

func (i *dictTypeCtl) Dict(c *fiber.Ctx) error {
	return vo.ResultOK(i.srv.Dict(), c)
}
