package system

type User struct {
	Model
	Username *string `json:"username"`
	Password *string `json:"password"`
	Nickname *string `json:"nickname"`
	RoleId   *uint   `json:"roleId"`
	Status   *int    `json:"status"`
}

func (User) TableName() string {
	return "sys_user"
}
