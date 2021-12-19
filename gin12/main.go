package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Person 的复数是 people ！！！
type Person struct {
	ID   int64
	Name string `gorm:"default:'binshow'"` //通过tag定义字段的默认值，在创建记录时候生成的 SQL 语句会排除没有值或值为 零值 的字段。 在将记录插入到数据库后，Gorm会从数据库加载那些字段的默认值。
	Age  int64
}

// gorm 连接数据库
func main() {

	db, err := gorm.Open("mysql", "root:bin123456@(127.0.0.1:3306)/binshow?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Person{})

	//1. 安全创建记录
	p1 := Person{Name: "zkd", Age: 18}
	db.NewRecord(&p1) // 主键为空返回`true`
	db.Create(&p1)
	db.NewRecord(&p1) // 创建`person`后返回`false`

	p2 := Person{Name: "", Age: 24}
	db.Create(&p2) //实际执行的SQL语句是INSERT INTO users("age") values('99');，排除了零值字段Name，而在数据库中这一条数据会使用设置的默认值binshow作为Name字段的值。

	/**
	所有字段的零值, 比如0, "",false或者其它零值，都不会保存到数据库内，但会使用他们的默认值.
	如果你想避免这种情况，可以考虑使用指针或实现 Scanner/Valuer接口：

	// 使用指针
	type User struct {
	  ID   int64
	  Name *string `gorm:"default:'小王子'"`
	  Age  int64
	}
	user := User{Name: new(string), Age: 18))}
	db.Create(&user)  // 此时数据库中该条记录name字段的值就是''

	// 使用 Scanner/Valuer
	type User struct {
		ID int64
		Name sql.NullString `gorm:"default:'小王子'"` // sql.NullString 实现了Scanner/Valuer接口
		Age  int64
	}
	user := User{Name: sql.NullString{"", true}, Age:18}
	db.Create(&user)  // 此时数据库中该条记录name字段的值就是''

	*/

}
