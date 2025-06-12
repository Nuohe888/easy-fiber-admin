package vo

// LoginReq 登录接口入参
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginRes 登录接口返回
type LoginRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// InfoRes 用户信息接口返回
type InfoRes struct {
	Id       uint   `json:"userId"`
	Avatar   string `json:"avatar"`
	Username string `json:"userName"`
	Nickname string `json:"nickName"`
}

type EditPasswordReq struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
