package controller

import (
	"easy-fiber-admin/model/system"
	"easy-fiber-admin/module/system/internal/service"
	"easy-fiber-admin/pkg/common/utils"
	vo2 "easy-fiber-admin/pkg/common/vo"
	"errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type userCtl struct {
	srv *service.UserSrv
}

var UserCtl *userCtl

func InitUserCtl() {
	UserCtl = &userCtl{
		srv: service.GetUserSrv(),
	}
}

func (i *userCtl) Login(c *fiber.Ctx) error {
	var req vo2.LoginReq
	if err := vo2.BodyParser(&req, c); err != nil {
		return err
	}
	res, err := i.srv.Login(&req)
	if err != nil {
		return vo2.ResultErr(err, c)
	}
	return vo2.ResultOK(res, c)
}

func (i *userCtl) Info(c *fiber.Ctx) error {
	info := utils.GetUserInfo(c)
	res, err := i.srv.Info(info.Id)
	if err != nil {
		return vo2.ResultErr(err, c)
	}
	return vo2.ResultOK(res, c)
}

func (i *userCtl) Refresh(c *fiber.Ctx) error {
	return vo2.ResultErr(errors.New("token已过期,请重新登录"), c)
}

func (i *userCtl) Codes(c *fiber.Ctx) error {
	var res []string
	res = append(res, utils.GetUserInfo(c).RoleCode)
	return vo2.ResultOK(res, c)
}

func (i *userCtl) Logout(c *fiber.Ctx) error {
	return vo2.ResultOK(nil, c)
}

func (i *userCtl) Add(c *fiber.Ctx) error {
	var req system.User
	if err := vo2.BodyParser(&req, c); err != nil {
		return err
	}
	if err := i.srv.Add(&req); err != nil {
		return vo2.ResultErr(err, c)
	}
	return vo2.ResultOK(nil, c)
}

func (i *userCtl) Del(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := i.srv.Del(id); err != nil {
		return vo2.ResultErr(err, c)
	}
	return vo2.ResultOK(nil, c)
}

func (i *userCtl) Put(c *fiber.Ctx) error {
	id := c.Params("id")
	var req system.User
	if err := vo2.BodyParser(&req, c); err != nil {
		return err
	}
	if err := i.srv.Put(id, &req); err != nil {
		return vo2.ResultErr(err, c)
	}
	return vo2.ResultOK(nil, c)
}

func (i *userCtl) Get(c *fiber.Ctx) error {
	id := c.Query("id")
	return vo2.ResultOK(i.srv.Get(id), c)
}

func (i *userCtl) List(c *fiber.Ctx) error {
	page := c.Query("page")
	limit := c.Query("limit")
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	return vo2.ResultOK(i.srv.List(pageInt, limitInt), c)
}

func (i *userCtl) GetStatus(c *fiber.Ctx) error {
	return vo2.ResultOK(i.srv.GetStatus(), c)
}
