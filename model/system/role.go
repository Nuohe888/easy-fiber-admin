package system

import "gorm.io/gorm"

type Role struct {
	Model
	Name   *string `json:"name"`
	Code   *string `json:"code"`
	Desc   *string `json:"desc"`
	Status *int    `json:"status"`
}

func (*Role) TableName() string {
	return "sys_role"
}

func (i *Role) BeforeCreate(tx *gorm.DB) (err error) {
	setDefault(&i.Desc, "")

	return nil
}
