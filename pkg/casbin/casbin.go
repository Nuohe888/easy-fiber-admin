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
	adminEnforcer *casbin.Enforcer
	userEnforcer  *casbin.Enforcer
	adminOnce     = sync.Once{}
	userOnce      = sync.Once{}
)

func initAdmin(db *gorm.DB) {
	adminOnce.Do(func() {
		var err error
		policyPath, err := gormadapter.NewAdapterByDBUseTableName(db, "casbin_rule", "admin")
		if err != nil {
			panic(fmt.Sprintf("Failed to create Admin Casbin adapter: %v", err))
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
		adminEnforcer, err = casbin.NewEnforcer(m, policyPath)
		if err != nil {
			panic(fmt.Sprintf("Failed to create Admin Casbin enforcer: %v", err))
		}
	})
}

func initUser(db *gorm.DB) {
	userOnce.Do(func() {
		var err error
		policyPath, err := gormadapter.NewAdapterByDBUseTableName(db, "casbin_rule", "user")
		if err != nil {
			panic(fmt.Sprintf("Failed to create User Casbin adapter: %v", err))
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
		userEnforcer, err = casbin.NewEnforcer(m, policyPath)
		if err != nil {
			panic(fmt.Sprintf("Failed to create User Casbin enforcer: %v", err))
		}
	})
}

func GetAdmin() *casbin.Enforcer {
	if adminEnforcer == nil {
		panic("Admin Casbin enforcer not initialized")
	}
	return adminEnforcer
}

func GetUser() *casbin.Enforcer {
	if userEnforcer == nil {
		panic("User Casbin enforcer not initialized")
	}
	return userEnforcer
}

func Init(db *gorm.DB) {
	initAdmin(db)
}
