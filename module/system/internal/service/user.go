package service

import (
	"easy-fiber-admin/model/system"
	"easy-fiber-admin/pkg/common/utils"
	"easy-fiber-admin/pkg/common/vo"
	"easy-fiber-admin/pkg/jwt"
	"easy-fiber-admin/pkg/logger"
	"easy-fiber-admin/pkg/sql"
	"easy-fiber-admin/plugin"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"

	redisClient "easy-fiber-admin/pkg/redis"
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
		sentry.CaptureException(fmt.Errorf("failed to generate token for user %s: %w", req.Username, err))
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
	cacheKey := fmt.Sprintf("user:%v", id)

	// Try to get from Redis
	cachedUser, err := redisClient.Get(cacheKey)
	if err != nil && err != redis.Nil {
		i.log.Errorf("failed to get user from redis: %v", err)
		// Proceed to fetch from DB if Redis error occurs
	}

	if cachedUser != "" {
		err := json.Unmarshal([]byte(cachedUser), &user)
		if err == nil {
			i.log.Infof("user %v found in cache", id)
			return user
		}
		i.log.Errorf("failed to unmarshal cached user data: %v", err)
		// Proceed to fetch from DB if unmarshal error
	}

	// If not in Redis or error, get from DB
	i.log.Infof("user %v not found in cache, fetching from DB", id)
	if errDb := i.db.Where("id = ?", id).Find(&user).Error; errDb != nil {
		i.log.Errorf("failed to get user from db: %v", errDb)
		return system.User{} // Return empty user on DB error
	}

	// Store in Redis
	userJSON, err := json.Marshal(user)
	if err != nil {
		i.log.Errorf("failed to marshal user data for caching: %v", err)
	} else {
		err = redisClient.Set(cacheKey, userJSON, 10*time.Minute)
		if err != nil {
			i.log.Errorf("failed to set user in redis: %v", err)
		}
	}

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
