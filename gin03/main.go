package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Gender string
	Age int
}


// 使用标准的http库来实现
func main() {
	http.HandleFunc("/" , sayHello)
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Println("HTTP Server start failed ,err: %v\n" , err)
		return
	}
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	//2. 解析模板
	t, err := template.ParseFiles("gin03/hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err : %v" , err)
		return
	}

	//3. 渲染模板
	u1 := Person{
		Name:   "binshow",
		Gender: "男",
		Age:    18,
	}

	m1 := map[string]interface{}{
		"Name" : "zkd",
		"Gender" : "女",
		"Age" : 20,
	}

	err = t.Execute(writer, map[string]interface{}{
		"u1" : u1,
		"m1" : m1,
	})
	if err != nil {
		fmt.Printf("render template failed, err : %v" , err)
		return
	}
}
