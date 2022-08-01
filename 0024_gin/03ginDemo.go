package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func StateCost() gin.HandlerFunc {
	return func(context *gin.Context) {
		now := time.Now()
		context.Set("name", "tom")
		context.Next()
		cost := time.Since(now)
		log.Println("login cost:", cost)
	}
}

/*
*
gin中间件,类似于java中的切面
Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。
这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。
Gin中的中间件必须是一个gin.HandlerFunc类型。
*/
func main() {
	engine := gin.Default()
	engine.Use(StateCost())
	engine.GET("/login", func(context *gin.Context) {
		value, exists := context.Get("name")
		if exists {
			fmt.Println("name value is :", value)
		}
		time.Sleep(time.Second * 2)
		context.JSON(http.StatusOK, gin.H{
			"msg": "login success!",
		})
	})
	engine.Run()
}
