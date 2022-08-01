package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type msg struct {
	Name    string `json:"user"`
	Message string
	Age     int
}

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/ping：请求的路径
	r.GET("/ping", func(context *gin.Context) {
		// context.JSON：返回JSON格式的数据
		context.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/moreJson", func(context *gin.Context) {
		var msg msg
		context.JSON(http.StatusOK, msg)
	})
	r.GET("/moreXml", func(context *gin.Context) {
		var msg = msg{"tom", "hello", 18}
		context.XML(http.StatusOK, msg)
	})
	r.GET("/moreYaml", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})
	//获取获取querystring参数
	r.POST("/queryString", func(context *gin.Context) {
		name := context.DefaultQuery("name", "tom")
		age := context.Query("age")
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
			"name":    name,
			"age":     age,
		})
	})

	//获取form参数
	r.POST("/formData", func(context *gin.Context) {
		name := context.PostForm("name")
		age := context.PostForm("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	//获取json参数
	r.POST("/jsonData", func(c *gin.Context) {
		b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
		// 定义map或结构体
		var m map[string]interface{}
		// 反序列化
		_ = json.Unmarshal(b, &m)
		c.JSON(http.StatusOK, m)
	})
	//获取path参数
	r.POST("/path/:name/:age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
