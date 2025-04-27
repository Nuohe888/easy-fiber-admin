package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-server/module/system/internal/service"
	"go-server/module/system/internal/utils"
	"go-server/module/system/internal/vo"
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
	var req vo.LoginReq
	if err := vo.BodyParser(&req, c); err != nil {
		return err
	}
	res, err := i.srv.Login(&req)
	if err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(res, c)
}

func (i *userCtl) Info(c *fiber.Ctx) error {
	info := utils.GetUserInfo(c)
	var res vo.InfoRes
	var roles []string
	roles = append(roles, "admin")
	res.RealName = "管理员"
	res.Id = info.Id
	res.Username = info.Username
	res.Roles = roles
	return vo.ResultOK(info, c)
}

func (i *userCtl) Codes(c *fiber.Ctx) error {
	var res []string
	res = append(res, "admin0")
	res = append(res, "admin1")
	res = append(res, "admin2")
	return vo.ResultOK(res, c)
}

func (i *userCtl) Logout(c *fiber.Ctx) error {
	return vo.ResultOK(nil, c)
}
