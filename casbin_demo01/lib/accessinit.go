package lib

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

var E *casbin.Enforcer

func init() {
	initDB()
	adapter, err := gormadapter.NewAdapterByDB(Gorm)
	if err != nil {
		log.Fatal(err)
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

func initPolicy() {
	//E.AddPolicy("member", "/depts", "GET")
	//E.AddPolicy("admin", "/depts", "POST")
	//E.AddRoleForUser("zhangsan", "member")
	return

	m := make([]*RoleRel, 0)
	GetRoles(0, &m, "") //获得角色对应
	for _, r := range m {
		_, err := E.AddRoleForUser(r.PRole, r.Role)
		if err != nil {
			log.Fatal(err)
		}
	}
	////测试获得用户角色权限数据,权限继承数据
	//GetRoles(0, &m, "")
	//fmt.Println(m)

	//初始化用户角色,显示用户权限
	userRoles := GetUserRoles()
	for _, ur := range userRoles {
		_, err := E.AddRoleForUser(ur.UserName, ur.RoleName)
		if err != nil {
			log.Fatal(err)
		}
	}
	//fmt.Println(userRoles)

	//初始化路由角色
	routerRoles := GetRouterRoles()
	for _, rr := range routerRoles {
		_, err := E.AddPolicy(rr.RoleName, rr.RouterUri, rr.RouterMethod)
		if err != nil {
			log.Fatal(err)
		}
	}
}
