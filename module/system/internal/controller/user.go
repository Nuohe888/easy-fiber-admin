package controller

import (
	"easy-fiber-admin/model/system"
	"easy-fiber-admin/module/system/internal/service"
	"easy-fiber-admin/pkg/common/utils"
	"easy-fiber-admin/pkg/common/vo"
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

// @Summary User login
// @Description Authenticates a user and returns a JWT token.
// @Tags User
// @Accept json
// @Produce json
// @Param loginBody body vo.LoginReq true "Login Credentials"
// @Success 200 {object} vo.Response{data=vo.LoginRes} "Successful login"
// @Failure 400 {object} vo.Response "Bad Request"
// @Failure 401 {object} vo.Response "Unauthorized"
// @Router /user/login [post]
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

// @Summary Get user information
// @Description Retrieves information for the currently authenticated user.
// @Tags User
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} vo.Response{data=vo.InfoRes} "User information"
// @Failure 401 {object} vo.Response "Unauthorized"
// @Router /user/info [get]
func (i *userCtl) Info(c *fiber.Ctx) error {
	info := utils.GetUserInfo(c)
	res, err := i.srv.Info(info.Id)
	if err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(res, c)
}

func (i *userCtl) Refresh(c *fiber.Ctx) error {
	return vo.ResultErr(errors.New("token已过期,请重新登录"), c)
}

func (i *userCtl) Codes(c *fiber.Ctx) error {
	var res []string
	res = append(res, utils.GetUserInfo(c).RoleCode)
	return vo.ResultOK(res, c)
}

func (i *userCtl) Logout(c *fiber.Ctx) error {
	return vo.ResultOK(nil, c)
}

func (i *userCtl) Add(c *fiber.Ctx) error {
	var req system.User
	if err := vo.BodyParser(&req, c); err != nil {
		return err
	}
	if err := i.srv.Add(&req); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *userCtl) Del(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := i.srv.Del(id); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *userCtl) Put(c *fiber.Ctx) error {
	id := c.Params("id")
	var req system.User
	if err := vo.BodyParser(&req, c); err != nil {
		return err
	}
	if err := i.srv.Put(id, &req); err != nil {
		return vo.ResultErr(err, c)
	}
	return vo.ResultOK(nil, c)
}

func (i *userCtl) Get(c *fiber.Ctx) error {
	id := c.Query("id")
	return vo.ResultOK(i.srv.Get(id), c)
}

// @Summary List users
// @Description Retrieves a list of users with pagination.
// @Tags User
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} vo.Response{data=vo.List{items=[]system.User}} "List of users"
// @Failure 401 {object} vo.Response "Unauthorized"
// @Router /user/list [get]
func (i *userCtl) List(c *fiber.Ctx) error {
	page := c.Query("page")
	limit := c.Query("limit")
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	return vo.ResultOK(i.srv.List(pageInt, limitInt), c)
}

func (i *userCtl) GetStatus(c *fiber.Ctx) error {
	return vo.ResultOK(i.srv.GetStatus(), c)
}

func (i *userCtl) EditPassword(c *fiber.Ctx) error {
	var req vo.EditPasswordReq
	if err := vo.BodyParser(&req, c); err != nil {
		return vo.ResultErr(err, c)
	}
	info := utils.GetUserInfo(c)
	return vo.ResultOK(i.srv.EditPassword(req, info.Id), c)
}
