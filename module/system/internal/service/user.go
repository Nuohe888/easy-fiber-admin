package service

import (
	"easy-fiber-admin/model/system"
	"easy-fiber-admin/module/system/internal/utils"
	"easy-fiber-admin/module/system/internal/vo"
	"easy-fiber-admin/pkg/jwt"
	"easy-fiber-admin/pkg/logger"
	"easy-fiber-admin/pkg/sql"
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserSrv struct {
	db  *gorm.DB
	log logger.ILog
}

var userSrv *UserSrv

func InitUserSrv() {
	userSrv = &UserSrv{
		db:  sql.Get(),
		log: logger.Get(),
	}
}

func GetUserSrv() *UserSrv {
	if userSrv == nil {
		panic("service user init failed")
	}
	return userSrv
}

func (i *UserSrv) Ping() error {
	return nil
}

func (i *UserSrv) Login(req *vo.LoginReq) (*vo.LoginRes, error) {
	var user system.User
	if err := i.db.Where("username =?", req.Username).Find(&user).Error; err != nil {
		return nil, errors.New("账号或密码错误")
	}

	if *user.Username == "" {
		return nil, errors.New("账号或密码错误")
	}

	if req.Password != *user.Password {
		return nil, errors.New("密码错误")
	}

	var role system.Role
	i.db.Where("id = ?", user.RoleId).Find(&role)

	var roles []string
	roles = append(roles, *role.Code)

	// 生成token
	now := time.Now()
	expTime, _ := jwt.GetAccessExpTime(now)

	claims := &vo.UserInfoJwtClaims{
		Id:             *user.Id,
		Username:       *user.Username,
		RoleCode:       *role.Code,
		IssuedAt:       now,
		ExpirationTime: expTime,
	}

	accessToken, err := jwt.GenToken(claims)
	if err != nil {
		return nil, errors.New("系统错误")
	}

	return &vo.LoginRes{
		RealName:    "管理员",
		Roles:       roles,
		Username:    *user.Username,
		AccessToken: accessToken,
	}, nil
}

func (i *UserSrv) Info(id uint) (*vo.InfoRes, error) {
	var user system.User
	i.db.Where("id=?", id).Find(&user)
	if *user.Username == "" {
		return nil, errors.New("该用户不存在")
	}
	var role system.Role
	i.db.Where("id=?", user.RoleId).Find(&role)
	return &vo.InfoRes{
		Id:       id,
		RealName: *user.Username,
		Roles:    []string{*role.Code},
		Username: *user.Username,
	}, nil
}

func (i *UserSrv) Add(user *system.User) error {
	return i.db.Create(&user).Error
}

func (i *UserSrv) Del(id any) error {
	return i.db.Where("id = ?", id).Delete(&system.User{}).Error
}

func (i *UserSrv) Put(id any, user *system.User) error {
	var _user system.User
	i.db.Where("id=?", id).Find(&_user).Find(&_user)
	if *_user.Id == 0 {
		return errors.New("不存在该Id")
	}
	utils.MergeStructs(&_user, user)
	return i.db.Save(&_user).Error
}

func (i *UserSrv) Get(id any) system.User {
	var user system.User
	i.db.Where("id = ?", id).Find(&user)
	return user
}

func (i *UserSrv) List(page, limit int) *vo.List {
	var items []system.User
	var total int64
	if limit == 0 {
		limit = 20
	}
	db := i.db
	i.db.Limit(limit).Offset((page - 1) * limit).Find(&items)
	db.Model(&system.User{}).Count(&total)
	return &vo.List{
		Items: items,
		Total: total,
	}
}
