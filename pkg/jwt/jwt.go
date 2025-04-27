package jwt

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type _jwt struct {
	SigningKey []byte
}

type IClaims interface {
	GetSubject() string           // GetSubject 主题，具有本地唯一性或者全局唯一性
	GetIssuer() string            // GetIssuer 颁发人
	GetIssuedAt() time.Time       // GetIssuedAt 颁发时间
	GetExpirationTime() time.Time // GetExpirationTime 过期时间
}

var (
	ErrTokenExpired = errors.New("令牌过期")
	ErrTokenInvalid = errors.New("无法处理这个令牌")

	runtimeJWT *_jwt
	one        = sync.Once{}
)

const (
	AccessTokenExpiredDuration = 7 * 24 * time.Hour
	SigningKey                 = "CloudSale20250101"
)

func init() {
	one.Do(func() {
		runtimeJWT = &_jwt{
			SigningKey: []byte(SigningKey),
		}
	})
}

// 适配 jwt.KeyFunc 方法
func (_this *_jwt) keyFunc(_ *jwt.Token) (any, error) {
	return _this.SigningKey, nil
}

// GenToken 生成Token
func GenToken(claims IClaims) (token string, err error) {
	s, err := newClaims(claims)
	if err != nil {
		return "", err
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, s).SignedString(runtimeJWT.SigningKey)
	if err != nil {
		return "", err
	}
	return
}

// VerifyToken 验证Token
func VerifyToken[T IClaims](tokenString string) (T, error) {
	var sc = new(selfClaims)
	var t T
	token, err := jwt.ParseWithClaims(tokenString, sc, runtimeJWT.keyFunc)
	if err != nil {
		return t, err
	}
	// 验证有效性
	if !token.Valid {
		err = ErrTokenInvalid
		return t, err
	}
	err = json.Unmarshal(sc.Payload, &t)
	if err != nil {
		return t, err
	}
	if time.Now().After(t.GetExpirationTime()) {
		return t, ErrTokenExpired
	}
	return t, err
}
