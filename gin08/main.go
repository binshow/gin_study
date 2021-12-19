package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		//1. 使用map
		data := map[string]interface{}{
			"name":    "binshow",
			"message": "hello,world",
			"age":     20,
		}

		//2. gin.H
		data = gin.H{
			"name":    "zkd",
			"message": "test",
			"age":     18,
		}

		//3. 定义结构体
		type msg struct {
			Name    string
			Age     int
			Message string
		}

		m := msg{
			Name:    "测试",
			Age:     21,
			Message: "I am struct",
		}

		fmt.Println(data)
		c.JSON(http.StatusOK, m) //默认使用的是json的序列化， 如果结构体中字段首字母是小写的，就不能导出了
	})

	r.Run(":9090")
}
