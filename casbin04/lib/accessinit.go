package lib

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	"log"
)

var E *casbin.Enforcer

func init() {
	initDB()
	adapter, err := gormadapter.NewAdapterByDB(Gorm)
	if err != nil {
		log.Fatal()
	}

	e, err := casbin.NewEnforcer("resources/model.conf", adapter)
	if err != nil {
		log.Fatal()
	}
	err = e.LoadPolicy()
	if err != nil {
		log.Fatal()
	}
	E = e
	initPolicy()
}

//从我们库里初始化 策略数据
func initPolicy() {
	//用代码写入数据库里面的p的数据
	E.AddPolicy("member", "/depts", "GET")
	E.AddPolicy("admin", "/depts", "POST")
	//用代码写入g的数据中
	E.AddRoleForUser("zhangsan", "member")
}
