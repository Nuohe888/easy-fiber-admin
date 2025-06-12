package system

import "gorm.io/gorm"

type User struct {
	Model
	Username *string `json:"username"`
	Password *string `json:"password"`
	Nickname *string `json:"nickname"`
	Avatar   *string `json:"avatar"`
	RoleId   *uint   `json:"roleId"`
	Status   *int    `json:"status"`
}

func (i *User) TableName() string {
	return "sys_user"
}

func (i *User) BeforeCreate(tx *gorm.DB) (err error) {
	setDefault(&i.Nickname, "")
	setDefault(&i.Avatar, "")

	return nil
}
