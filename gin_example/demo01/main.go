package main

import "github.com/gin-gonic/gin"

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/7/25 10:27 PM
// @description   : Bind form-data request with custom struct
// -------------------------------------------

// 接收表单数据并转换成 struct 属性
//$ curl "http://localhost:8080/getb?field_a=hello&field_b=world"
//{"a":{"FieldA":"hello"},"b":"world"}
//$ curl "http://localhost:8080/getc?field_a=hello&field_c=world"
//{"a":{"FieldA":"hello"},"c":"world"}
//$ curl "http://localhost:8080/getd?field_x=hello&field_d=world"
//{"d":"world","x":{"FieldX":"hello"}}

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

func main() {
	r := gin.Default()
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)

	r.Run()
}