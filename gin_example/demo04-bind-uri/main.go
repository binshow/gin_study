package main

import "github.com/gin-gonic/gin"

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/7/25 10:36 PM
// @description   :	Bind Uri
// -------------------------------------------


//$ curl -v localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
//$ curl -v localhost:8088/thinkerou/not-uuid

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run(":8088")
}