package system

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "sys_user"
}
