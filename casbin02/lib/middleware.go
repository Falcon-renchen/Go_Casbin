package lib

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func CheckLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.Header.Get("token") == "" {
			context.AbortWithStatusJSON(400, gin.H{
				"message": "token required",
			})
		} else {
			//传给RBAC
			context.Set("user_name", context.Request.Header.Get("token"))
			context.Next()
		}
	}
}

func RBAC() gin.HandlerFunc {
	e := casbin.NewEnforcer("resources/model.conf", "resources/p.csv")
	return func(context *gin.Context) {
		//获得login中的名字
		user, _ := context.Get("user_name")
		if !e.Enforce(user, context.Request.RequestURI, context.Request.Method) {
			context.AbortWithStatusJSON(403, gin.H{
				"message": "forbidden",
			})
		} else {
			context.Next()
		}
	}
}

//假设都要登录才能访问
func Middleware() (fs []gin.HandlerFunc) {
	fs = append(fs, CheckLogin(), RBAC())
	return
}
