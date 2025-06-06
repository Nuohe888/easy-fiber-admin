package vo

import (
	"time"
)

type UserInfoJwtClaims struct {
	Id             uint      `json:"id"`
	RoleCode       string    `json:"roleCode"`
	Username       string    `json:"username"`
	IssuedAt       time.Time `json:"issuedAt"`
	ExpirationTime time.Time `json:"expirationTime"`
}

func (a *UserInfoJwtClaims) GetSubject() string           { return a.Username }
func (a *UserInfoJwtClaims) GetIssuer() string            { return "AdminUser" }
func (a *UserInfoJwtClaims) GetIssuedAt() time.Time       { return a.IssuedAt }
func (a *UserInfoJwtClaims) GetExpirationTime() time.Time { return a.ExpirationTime }

type UserCenterInfoJwtClaims struct {
	Id             uint      `json:"id"`
	RoleCode       string    `json:"roleCode"`
	Username       string    `json:"username"`
	IssuedAt       time.Time `json:"issuedAt"`
	ExpirationTime time.Time `json:"expirationTime"`
}

func (a *UserCenterInfoJwtClaims) GetSubject() string           { return a.Username }
func (a *UserCenterInfoJwtClaims) GetIssuer() string            { return "WebUserCenter" }
func (a *UserCenterInfoJwtClaims) GetIssuedAt() time.Time       { return a.IssuedAt }
func (a *UserCenterInfoJwtClaims) GetExpirationTime() time.Time { return a.ExpirationTime }
