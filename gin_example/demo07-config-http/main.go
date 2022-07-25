package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/7/25 10:42 PM
// @description   : Custom HTTP configuration
// -------------------------------------------

func main() {
	router := gin.Default()

	// Handler 用的是 gin 的Engine，其余的可以用配置文件来进行配置
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}