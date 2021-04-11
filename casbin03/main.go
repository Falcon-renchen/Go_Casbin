package main

import (
	"Go_Casbin/casbin03/lib"
	"github.com/gin-gonic/gin"
)

func main() {

	//登录才能访问
	r := gin.New()
	//gin 的 func里面的函数写在lib/middleware
	r.Use(lib.Middleware()...)

	r.GET("/depts", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "部门列表",
		})
	})

	r.POST("/depts", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "批量修改部门列表",
		})
	})
	r.Run(":8080")
}
