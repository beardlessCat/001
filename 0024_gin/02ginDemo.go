package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login /*
type Login struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

/*
*
参数绑定
*/
func main() {
	engine := gin.Default()
	//json绑定
	engine.POST("/json/bind", func(context *gin.Context) {
		var login Login
		if err := context.ShouldBind(&login); err == nil {
			context.JSON(http.StatusOK, gin.H{
				"name":     login.Name,
				"password": login.Password,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	//重定向
	engine.GET("/index", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.sd.sgcc.com.cn")
	})
	//路由组
	userGroup := engine.Group("/user")
	{
		userGroup.GET("/get", func(context *gin.Context) {

		})
		userGroup.POST("/update", func(context *gin.Context) {

		})
		userGroup.POST("/delete", func(context *gin.Context) {

		})
	}

	engine.Run()
}
