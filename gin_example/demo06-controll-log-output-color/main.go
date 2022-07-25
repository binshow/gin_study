package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/7/25 10:40 PM
// @description   : Controlling Log output coloring
// -------------------------------------------


func main() {
	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		log.Println("test")
		c.String(200, "pong")
	})

	router.Run(":8080")
}