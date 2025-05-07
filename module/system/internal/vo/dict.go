package vo

import "easy-fiber-admin/model/system"

type Dict struct {
	Type string            `json:"type"`
	List []system.DictData `json:"list"`
}
