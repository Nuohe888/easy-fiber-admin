package controller

import (
	"easy-fiber-admin/model/system"
	"easy-fiber-admin/module/system/internal/service"
	"easy-fiber-admin/pkg/common/vo"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type userCenterCtl struct {
	srv *service.UserCenterSrv
}

var UserCenterCtl *userCenterCtl

func InitUserCenterCtl() {
	UserCenterCtl = &userCenterCtl{
		srv: service.GetUserCenterSrv(),
	}
}

func (i *userCenterCtl) Add(c *fiber.Ctx) error {
	var req system.UserCenter
	if err := vo.BodyParser(&req, c); err != nil {
		return err
	}
	if err := i.srv.Add(&req); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *userCenterCtl) Del(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := i.srv.Del(id); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *userCenterCtl) Put(c *fiber.Ctx) error {
	id := c.Params("id")
	var req system.UserCenter
	if err := vo.BodyParser(&req, c); err != nil {
		return err
	}
	if err := i.srv.Put(id, &req); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *userCenterCtl) Get(c *fiber.Ctx) error {
	id := c.Query("id")
	return vo.ResultOK(i.srv.Get(id), c)
}

func (i *userCenterCtl) List(c *fiber.Ctx) error {
	page := c.Query("page")
	limit := c.Query("limit")
	username := c.Query("username")
	nickname := c.Query("nickname")
	phone := c.Query("phone")
	email := c.Query("email")
	status := c.Query("status")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	var statusInt *int
	if status != "" {
		if s, err := strconv.Atoi(status); err == nil {
			statusInt = &s
		}
	}

	return vo.ResultOK(i.srv.List(pageInt, limitInt, username, nickname, phone, email, statusInt), c)
}

func (i *userCenterCtl) GetStatus(c *fiber.Ctx) error {
	return vo.ResultOK(i.srv.GetStatus(), c)
}
