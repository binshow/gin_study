package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint // 默认就是主键
	Name   string
	Gender string
	Hobby  string
}

// gorm 连接数据库
func main() {

	db, err := gorm.Open("mysql", "root:bin123456@(127.0.0.1:3306)/binshow?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//1. 自动创建表，自动迁移，把数据库和数据表进行对应
	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{1, "binshow", "男", "英雄联盟"}
	u2 := UserInfo{2, "zkd", "女", "王者荣耀"}

	//2. 创建记录
	db.Create(&u1)
	db.Create(&u2)

	//3. 查询记录
	u := new(UserInfo)
	db.First(u) // 第一条数据
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "王者荣耀") // 根据条件查询
	fmt.Printf("%#v\n", uu)

	//4. 更新记录
	db.Model(&u).Update("hobby", "云顶之奕")

	//5. 删除记录
	db.Delete(&u)

}
