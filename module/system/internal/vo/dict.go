package vo

import "go-server/model/system"

type Dict struct {
	Type string            `json:"type"`
	List []system.DictData `json:"list"`
}
