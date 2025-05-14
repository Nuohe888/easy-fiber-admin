package system

type DictType struct {
	Model
	Name *string `json:"name"`
	Type *string `json:"type"`
	Desc *string `json:"desc"`
}

func (DictType) TableName() string {
	return "sys_dict_type"
}
