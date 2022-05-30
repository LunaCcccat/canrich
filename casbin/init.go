package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

var Casbin CasbinClient

type CasbinClient struct {
	enforcer *casbin.Enforcer
}

func InitCasbin() error {
	e, _ := casbin.NewEnforcer("./casbin/model.conf", "./casbin/policy.csv")
	Casbin.enforcer = e
	return nil
}

func (casbin *CasbinClient) Check(sub, obj, act string) bool {
	ok, _ := casbin.enforcer.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
		return true
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
		return false
	}
}
