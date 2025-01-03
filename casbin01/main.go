package main

import (
	"github.com/casbin/casbin"
	"log"
)

func main() {
	sub := "wyp" //想要访问资源的用户
	obj := "/depts" //将被访问的资源
	act := "POST" //用户对资源执行的操作
	e := casbin.NewEnforcer("resources/model.conf","resources/p.csv")
	ok := e.Enforce(sub,obj,act)

	if ok {
		log.Println("运行通过")
	}
}
