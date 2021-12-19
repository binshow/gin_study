package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name   string
	Gender string
	Age    int
}

// 使用标准的http库来实现
func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/demo", demo1)
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Println("HTTP Server start failed ,err: %v\n", err)
		return
	}
}

func f1(w http.ResponseWriter, r *http.Request) {

	//1. 定义一个函数
	kua := func(name string) (string, error) {
		return name + "真棒！", nil
	}

	//2. 创建一个模板，注意名称 必须和模板名称一致
	t := template.New("f.tmpl")

	//3. 告诉模板引擎，我们现在多了一个自定义的函数kua
	t.Funcs(template.FuncMap{
		"kua": kua,
	})

	// 解析模板
	_, err := t.ParseFiles("gin04/f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err : %v", err)
		return
	}
	// 渲染模板
	name := "binshow"
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {

	//定义模板
	t, err := template.ParseFiles("gin04/t.tmpl", "gin04/ul.tmpl") // 注意先写t在写uml
	//解析模板
	if err != nil {
		fmt.Printf("parse template failed, err : %v", err)
		return
	}
	// 渲染模板
	name := "binshow"
	t.Execute(w, name)

}
