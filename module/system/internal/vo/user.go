package vo

// LoginReq 登录接口入参
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Captcha  bool   `json:"captcha"`
}

// LoginRes 登录接口返回
type LoginRes struct {
	RealName    string   `json:"realName"`
	Roles       []string `json:"roles"`
	Username    string   `json:"username"`
	AccessToken string   `json:"accessToken"`
}

// InfoRes 用户信息接口返回
type InfoRes struct {
	Id       uint     `json:"id"`
	RealName string   `json:"realName"`
	Roles    []string `json:"roles"`
	Username string   `json:"username"`
}
