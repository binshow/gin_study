package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/7/25 10:32 PM
// @description   : Bind query string or post data
// -------------------------------------------

// 绑定查询参数 or post 请求中的数据
// curl -X GET "localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15"


type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success")
}