package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"sync"
)

var (
	enforcer *casbin.Enforcer
	once     = sync.Once{}
)

func Init(db *gorm.DB) {
	once.Do(func() {
		var err error
		policyPath, err := gormadapter.NewAdapterByDBUseTableName(db, "casbin_rule", "admin")
		if err != nil {
			panic(fmt.Sprintf("Failed to create Casbin gorm adapter: %v", err))
		}
		text := `
			[request_definition]
			r = sub, obj, act

			[policy_definition]
			p = sub, obj, act

			[role_definition]
			g = _, _

			[policy_effect]
			e = some(where (p.eft == allow))

			[matchers]
			m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && (r.act == p.act || p.act == "*")
			`
		m, _ := model.NewModelFromString(text)
		enforcer, err = casbin.NewEnforcer(m, policyPath)
		if err != nil {
			panic(fmt.Sprintf("Failed to create Casbin enforcer: %v", err))
		}
	})
}

func Get() *casbin.Enforcer {
	if enforcer == nil {
		panic(fmt.Sprintf("Casbin enforcer not initialized"))
	}
	return enforcer
}
