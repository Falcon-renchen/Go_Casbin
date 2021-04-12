package main

import (
	"github.com/casbin/casbin/v2"
	"log"
)

func main() {
	sub := "shenyi"
	obj := "/depts"
	act := "GET"

	e, _ := casbin.NewEnforcer("resources/model_t.conf", "resources/p_t.csv")

	ok, err := e.Enforce(sub, "domain1", obj, act)
	if err == nil && ok {
		log.Println("运行通过")
	}
}
