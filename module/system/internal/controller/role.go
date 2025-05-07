package controller

import (
	"easy-fiber-admin/model/system"
	"easy-fiber-admin/module/system/internal/service"
	"easy-fiber-admin/module/system/internal/vo"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type roleCtl struct {
	srv *service.RoleSrv
}

var RoleCtl *roleCtl

func InitRoleCtl() {
	RoleCtl = &roleCtl{
		srv: service.GetRoleSrv(),
	}
}

func (i *roleCtl) Add(c *fiber.Ctx) error {
	var req system.Role
	if err := vo.BodyParser(&req, c); err != nil {
		return err
	}
	if err := i.srv.Add(&req); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *roleCtl) Del(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := i.srv.Del(id); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *roleCtl) Put(c *fiber.Ctx) error {
	var req system.Role
	id := c.Params("id")
	if err := vo.BodyParser(&req, c); err != nil {
		return err
	}
	if err := i.srv.Put(id, &req); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *roleCtl) Get(c *fiber.Ctx) error {
	id := c.Query("id")
	return vo.ResultOK(i.srv.Get(id), c)
}

func (i *roleCtl) List(c *fiber.Ctx) error {
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	return vo.ResultOK(i.srv.List(pageInt, pageSizeInt), c)
}

func (i *roleCtl) ListAll(c *fiber.Ctx) error {
	return vo.ResultOK(i.srv.ListAll(), c)
}
