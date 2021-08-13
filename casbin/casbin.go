package casbin

import (
	"fmt"

	"github.com/authfun/gauthfun/database"
	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

var Enforcer *casbin.Enforcer

func init() {
	drop()

	adapter, err := gormadapter.NewAdapterByDB(database.AuthDatabase)
	if err != nil {
		panic(fmt.Errorf("fatal error init casbin adapter: %s", err))
	}

	enforcer, err := casbin.NewEnforcer("model.conf", adapter)
	if err != nil {
		panic(fmt.Errorf("fatal error init casbin enforcer: %s", err))
	}

	enforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddDomainMatchingFunc("keyMatch", util.KeyMatch)

	initRule(enforcer)

	Enforcer = enforcer
}

func drop() {
	db, err := database.AuthDatabase.DB()
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("DROP TABLE casbin_rule")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}

func initRule(enforcer *casbin.Enforcer) {
	ok, err := enforcer.AddNamedPolicy("p", "admin", "domain1", "data1", "read")
	fmt.Println(ok)
	if err != nil {
		panic(err)
	}
}
