package system

type Role struct {
	Model
	Name   *string `json:"name"`
	Code   *string `json:"code"`
	Desc   *string `json:"desc"`
	Status *int    `json:"status"`
}

func (Role) TableName() string {
	return "sys_role"
}
