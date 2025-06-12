package system

import "gorm.io/gorm"

type UserCenter struct {
	Model
	Username *string `json:"username"`
	Password *string `json:"password"`
	Salt     *string `json:"salt"`
	Nickname *string `json:"nickname"`
	Avatar   *string `json:"avatar"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email"`
	RealName *string `json:"realName"`
	IdCard   *string `json:"idCard"`
	Status   *int    `json:"status"`
}

func (*UserCenter) TableName() string {
	return "sys_user_center"
}

func (i *UserCenter) BeforeCreate(tx *gorm.DB) (err error) {
	setDefault(&i.Avatar, "")
	setDefault(&i.Phone, "")
	setDefault(&i.Email, "")
	setDefault(&i.RealName, "")
	setDefault(&i.IdCard, "")

	return nil
}
