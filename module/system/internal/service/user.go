package service

import (
	"errors"
	"go-server/model/system"
	"go-server/module/system/internal/vo"
	"go-server/pkg/jwt"
	"go-server/pkg/logger"
	"go-server/pkg/sql"
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

	if user.Username == "" {
		return nil, errors.New("账号或密码错误")
	}

	if req.Password != user.Password {
		return nil, errors.New("密码错误")
	}

	var roles []string
	roles = append(roles, "admin")

	// 生成token
	now := time.Now()
	expTime, _ := jwt.GetAccessExpTime(now)

	claims := &vo.UserInfoJwtClaims{
		Id:             user.Id,
		Username:       user.Username,
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
		Username:    user.Username,
		AccessToken: accessToken,
	}, nil
}
