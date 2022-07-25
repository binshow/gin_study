package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// 使用标准的http库来实现
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP Server start failed ,err: %v\n", err)
		return
	}
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	//2. 解析模板
	dir, err := os.Getwd()
	fmt.Println("dir = " , dir)
	t, err := template.ParseFiles("gin02/hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err : %v", err)
		return
	}

	//3. 渲染模板
	err = t.Execute(writer, "binshow")
	if err != nil {
		fmt.Printf("render template failed, err : %v", err)
		return
	}
}
