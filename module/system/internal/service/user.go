package service

import (
	"easy-fiber-admin/model/system"
	"easy-fiber-admin/pkg/common/utils"
	"easy-fiber-admin/pkg/common/vo"
	"easy-fiber-admin/pkg/jwt"
	"easy-fiber-admin/pkg/logger"
	"easy-fiber-admin/pkg/sql"
	"easy-fiber-admin/plugin"
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserSrv struct {
	db     *gorm.DB
	log    logger.ILog
	crypto plugin.ICrypto
}

var userSrv *UserSrv

func InitUserSrv() {
	userSrv = &UserSrv{
		db:     sql.Get(),
		log:    logger.Get(),
		crypto: plugin.GetCrypto(),
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

	if user.Username == nil || *user.Username == "" {
		return nil, errors.New("账号或密码错误")
	}

	if user.Password == nil {
		return nil, errors.New("系统出错,请检查后台管理员账户")
	}

	if !i.crypto.VerifyPassword(req.Password, *user.Password) {
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
		//RealName:    "管理员",
		//Roles:       roles,
		//Username:    *user.Username,
		AccessToken: accessToken,
	}, nil
}

func (i *UserSrv) Info(id uint) (*vo.InfoRes, error) {
	var user system.User
	i.db.Where("id=?", id).Find(&user)
	if user.Username == nil || *user.Username == "" {
		return nil, errors.New("该用户不存在")
	}
	var role system.Role
	i.db.Where("id=?", user.RoleId).Find(&role)
	return &vo.InfoRes{
		Id:       id,
		Avatar:   *user.Avatar,
		Username: *user.Username,
		Nickname: *user.Nickname,
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
	if _user.Id == nil || *_user.Id == 0 {
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

func (i *UserSrv) GetStatus() map[string]interface{} {
	return map[string]interface{}{
		"0": "禁用",
		"1": "启用",
	}
}

func (i *UserSrv) EditPassword(req vo.EditPasswordReq, userId any) error {
	var user system.User
	i.db.Where("id=?", userId).Find(&user)
	if user.Password == nil || *user.Password == "" {
		return errors.New("用户信息错误")
	}
	if !i.crypto.VerifyPassword(req.OldPassword, *user.Password) {
		return errors.New("旧密码错误")
	}

	password, err := i.crypto.EncryptPassword(req.NewPassword)
	if err != nil {
		return err
	}

	*user.Password = password

	return i.db.Save(&user).Error
}
