package vo

// LoginReq 登录接口入参
type LoginReq struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

// LoginRes 登录接口返回
type LoginRes struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

// InfoRes 用户信息接口返回
type InfoRes struct {
	UserId   string   `json:"userId"`
	UserName string   `json:"userName"`
	Roles    []string `json:"roles"`
	Buttons  []string `json:"buttons"`
}

type EditPasswordReq struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
