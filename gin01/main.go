package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	//1. 返回默认的路由引擎
	r := gin.Default()

	//2. 设置路由对应的处理函数
	r.GET("/hello" , sayHello)

	//3. 启动服务，设置监听的端口号
	r.Run(":8090")

	//4. 浏览器打开 http://localhost:8090/hello
}

func sayHello(c *gin.Context)  {
	c.JSON(http.StatusOK , gin.H{
		"message" : "hello,goland!",
	})
}
