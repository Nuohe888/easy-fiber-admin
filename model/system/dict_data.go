package system

type DictData struct {
	Model
	Value *int    `json:"value"`
	Key   *string `json:"key"`
	Sort  *int    `json:"sort"`
	Desc  *string `json:"desc"`
	Color *string `json:"color"`
	Pid   *uint   `json:"pid"`
}

func (DictData) TableName() string {
	return "sys_dict_data"
}
